package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/atkhx/gopangaea/internal/cli/command"
	"github.com/atkhx/gopangaea/internal/cli/parser"
	"github.com/atkhx/gopangaea/internal/cli/processor"
	change_preset "github.com/atkhx/gopangaea/internal/pkg/commands/change-preset"
	get_bank "github.com/atkhx/gopangaea/internal/pkg/commands/get-bank"
	get_device "github.com/atkhx/gopangaea/internal/pkg/commands/get-device"
	get_impulse_name "github.com/atkhx/gopangaea/internal/pkg/commands/get-impulse-name"
	get_impulse_names "github.com/atkhx/gopangaea/internal/pkg/commands/get-impulse-names"
	get_mode "github.com/atkhx/gopangaea/internal/pkg/commands/get-mode"
	get_settings "github.com/atkhx/gopangaea/internal/pkg/commands/get-settings"
	get_version "github.com/atkhx/gopangaea/internal/pkg/commands/get-version"
	"github.com/atkhx/gopangaea/internal/pkg/device"
	"github.com/atkhx/gopangaea/internal/pkg/device/deviceio"
	"github.com/jpoirier/gousb/usb"
)

func main() {
	ctx, done := context.WithCancel(context.Background())
	fmt.Println("#", "open connection")
	usbContext := usb.NewContext()

	defer func() {
		fmt.Println("#", "close usb context")
		if err := usbContext.Close(); err != nil {
			fmt.Println("#", "close connection error:", err)
		} else {
			fmt.Println("#", "close connection success")
		}
	}()

	devs, err := usbContext.ListDevices(func(desc *usb.Descriptor) bool {
		//devs, err := usbContext.OpenDevices(func(desc *usb.DeviceDesc) bool {
		return desc.Vendor == GetPangaeaVendor() && desc.Product == GetPangaeaProduct()
	})
	if err != nil {
		log.Println("#", "get devices list failed:", err)
		return
	}

	// All Devices returned from ListDevices must be closed.
	defer func() {
		fmt.Println("#", "close devices")
		for i, dev := range devs {
			if err := dev.Close(); err != nil {
				fmt.Printf("# device[%d] close error: %s\n", i, err)
			} else {
				fmt.Printf("# device[%d] closed\n", i)
			}
		}
	}()

	showInfo := func(dev *usb.Device) {
		desc := dev.Descriptor
		s, _ := dev.GetStringDescriptor(2)
		fmt.Printf("device descriptor: %s\n", s)
		fmt.Printf("device dev.desc.Vendor: %s\n", desc.Vendor)
		fmt.Printf("device dev.desc.Product: %s\n", desc.Product)
	}

	if len(devs) == 0 {
		fmt.Println("# pangaea not found")
		return
	}

	dev := devs[0]

	epBulkWrite, err := dev.OpenEndpoint(
		dev.Configs[0].Config,
		dev.Configs[0].Interfaces[1].Number,
		0,
		dev.Configs[0].Interfaces[1].Setups[0].Endpoints[0].Address|uint8(usb.ENDPOINT_DIR_OUT),
	)
	if err != nil {
		log.Fatalf("OpenEndpoint Write error for %v: %v", dev.Address, err)
	}

	epBulkRead, err := dev.OpenEndpoint(
		dev.Configs[0].Config,
		dev.Configs[0].Interfaces[1].Number,
		0,
		dev.Configs[0].Interfaces[1].Setups[0].Endpoints[1].Address,
	)
	if err != nil {
		log.Fatalf("OpenEndpoint Read error for %v: %v", dev.Address, err)
	}

	pangaea := device.New(
		deviceio.NewCommandWriter(epBulkWrite),
		deviceio.NewResponseReader(epBulkRead),
	)

	if s, err := pangaea.GetDevice(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("device:", s)
	}

	knownCommands := []string{
		get_bank.CliCommand,
		get_device.CliCommand,
		get_version.CliCommand,
		get_impulse_name.CliCommand,
		get_impulse_names.CliCommand,
		get_mode.CliCommand,
		get_settings.CliCommand,
		change_preset.CliCommand,
		"exit",
		"info",
		"help",
	}

	cmdParser := parser.New(map[string]command.Command{
		get_bank.CliCommand:          get_bank.New(),
		get_device.CliCommand:        get_device.New(),
		get_version.CliCommand:       get_version.New(),
		get_impulse_name.CliCommand:  get_impulse_name.New(),
		get_impulse_names.CliCommand: get_impulse_names.New(),
		get_mode.CliCommand:          get_mode.New(),
		get_settings.CliCommand:      get_settings.New(),
		change_preset.CliCommand:     change_preset.New(),
	})

	usage := func() {
		fmt.Println("use commands:")
		for _, cmd := range knownCommands {
			fmt.Println("\t", cmd)
		}
	}

	cmdProcessor := processor.New(pangaea)

	go func() {
		defer done()

		scanner := bufio.NewScanner(os.Stdin)

		readCommand := func() bool {
			fmt.Print("# ")
			return scanner.Scan()
		}

		for readCommand() {
			if scanner.Text() == "exit" {
				break
			}

			if scanner.Text() == "help" {
				usage()
				continue
			}

			if scanner.Text() == "info" {
				showInfo(dev)
				continue
			}

			in, err := cmdParser.Parse(scanner.Text())
			if err != nil {
				if err == parser.ErrCommandIsUnknown {
					usage()
					continue
				}

				fmt.Println("parse command error:", err)
				continue
			}

			out, err := cmdProcessor.Exec(in)
			if err != nil {
				fmt.Println("execute command error:", err)
				continue
			}

			if out != nil {
				fmt.Println(out)
			}
		}
	}()

	waitSignal(ctx)
}

func waitSignal(ctx context.Context) {
	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	select {
	case s := <-sigChan:
		fmt.Println("signal", s)

	case <-ctx.Done():
		fmt.Println("context closed")
	}
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
