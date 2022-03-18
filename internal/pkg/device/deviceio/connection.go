package deviceio

import (
	"github.com/jpoirier/gousb/usb"
	"github.com/pkg/errors"
)

func New(usbContext *usb.Context) *connection {
	return &connection{usbContext: usbContext}
}

type connection struct {
	usbContext *usb.Context

	connected bool
	device    *usb.Device
	closeFn   func()

	epBulkWrite usb.Endpoint
	epBulkRead  usb.Endpoint

	commandWriter  *commandWriter
	responseReader *responseReader
}

func (c *connection) IsConnected() bool {
	return c.connected
}

func (c *connection) Device() *usb.Device {
	return c.device
}

func (c *connection) WriteCommand(command string) error {
	if !c.connected {
		return errors.New("device not connected")
	}
	return c.commandWriter.Write(command)
}

func (c *connection) ReadResponse(command string, length int) ([]byte, error) {
	if !c.connected {
		return nil, errors.New("device not connected")
	}
	return c.responseReader.ReadWithSkipTails(command, length)
}

func (c *connection) Connect() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.Errorf("recovered panic: %v", e)
		}
		if err != nil {
			c.Disconnect()
		}
	}()

	c.device, c.closeFn, err = GetPangaeaDevice(c.usbContext)
	if err != nil {
		return errors.Wrap(err, "get devices list failed")
	}

	c.epBulkWrite, err = c.device.OpenEndpoint(
		c.device.Configs[0].Config,
		c.device.Configs[0].Interfaces[1].Number,
		0,
		c.device.Configs[0].Interfaces[1].Setups[0].Endpoints[0].Address|uint8(usb.ENDPOINT_DIR_OUT),
	)
	if err != nil {
		return errors.Errorf("OpenEndpoint Write error for %v: %v", c.device.Address, err)
	}

	c.epBulkRead, err = c.device.OpenEndpoint(
		c.device.Configs[0].Config,
		c.device.Configs[0].Interfaces[1].Number,
		0,
		c.device.Configs[0].Interfaces[1].Setups[0].Endpoints[1].Address,
	)
	if err != nil {
		return errors.Errorf("OpenEndpoint Read error for %v: %v", c.device.Address, err)
	}

	c.commandWriter = NewCommandWriter(c.epBulkWrite)
	c.responseReader = NewResponseReader(c.epBulkRead)

	c.connected = true
	return nil
}

func (c *connection) Disconnect() error {
	if c.closeFn != nil {
		c.closeFn()
		c.closeFn = nil
		c.device = nil
		c.epBulkRead = nil
		c.epBulkWrite = nil
	}
	return nil
}
