package deviceio

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/jpoirier/gousb/usb"
)

func GetPangaeaDevice(usbContext *usb.Context) (*usb.Device, func(), error) {
	devs, err := usbContext.ListDevices(func(desc *usb.Descriptor) bool {
		return desc.Vendor == GetPangaeaVendor() && desc.Product == GetPangaeaProduct()
	})

	// All Devices returned from ListDevices must be closed.
	closeFunc := func() {
		fmt.Println("#", "close devices")
		for i, dev := range devs {
			if err := dev.Close(); err != nil {
				fmt.Printf("# device[%d] close error: %s\n", i, err)
			} else {
				fmt.Printf("# device[%d] closed\n", i)
			}
		}
	}

	if err != nil {
		return nil, closeFunc, err
	}

	if len(devs) == 0 {
		return nil, closeFunc, errors.New("pangaea device not found")
	}

	return devs[0], closeFunc, nil
}

// GetPangaeaVendor returns the vendor ID of Pangaea CP-100 USB FS Mode
func GetPangaeaVendor() usb.ID {
	value, err := strconv.ParseUint("0483", 16, 16)
	if err != nil {
		log.Fatalln(err)
	}
	return usb.ID(value)
}

// GetPangaeaProduct returns the product ID of Pangaea CP-100 USB FS Mode
func GetPangaeaProduct() usb.ID {
	value, err := strconv.ParseUint("5740", 16, 16)
	if err != nil {
		log.Fatalln(err)
	}
	return usb.ID(value)
}
