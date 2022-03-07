package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/atkhx/gopangaea/internal/cli/command"
	change_preset "github.com/atkhx/gopangaea/internal/cli/command/change-preset"
	"github.com/atkhx/gopangaea/internal/cli/command/exit"
	get_bank "github.com/atkhx/gopangaea/internal/cli/command/get-bank"
	get_device "github.com/atkhx/gopangaea/internal/cli/command/get-device"
	get_impulse_name "github.com/atkhx/gopangaea/internal/cli/command/get-impulse-name"
	get_impulse_names "github.com/atkhx/gopangaea/internal/cli/command/get-impulse-names"
	get_mode "github.com/atkhx/gopangaea/internal/cli/command/get-mode"
	get_settings "github.com/atkhx/gopangaea/internal/cli/command/get-settings"
	get_version "github.com/atkhx/gopangaea/internal/cli/command/get-version"
	"github.com/atkhx/gopangaea/internal/cli/command/info"
	set_hp_filter_state "github.com/atkhx/gopangaea/internal/cli/command/set-hp-filter-state"
	set_hp_filter_value "github.com/atkhx/gopangaea/internal/cli/command/set-hp-filter-value"
	set_impulse_state "github.com/atkhx/gopangaea/internal/cli/command/set-impulse-state"
	set_lp_filter_state "github.com/atkhx/gopangaea/internal/cli/command/set-lp-filter-state"
	set_lp_filter_value "github.com/atkhx/gopangaea/internal/cli/command/set-lp-filter-value"
	set_master_volume "github.com/atkhx/gopangaea/internal/cli/command/set-master-volume"
	set_presence_state "github.com/atkhx/gopangaea/internal/cli/command/set-presence-state"
	set_presence_value "github.com/atkhx/gopangaea/internal/cli/command/set-presence-value"
	set_reverb_state "github.com/atkhx/gopangaea/internal/cli/command/set-reverb-state"
	set_reverb_type "github.com/atkhx/gopangaea/internal/cli/command/set-reverb-type"
	set_reverb_volume "github.com/atkhx/gopangaea/internal/cli/command/set-reverb-volume"
	"github.com/atkhx/gopangaea/internal/cli/command/usage"
	"github.com/atkhx/gopangaea/internal/cli/parser"
	"github.com/atkhx/gopangaea/internal/pkg/device"
	"github.com/atkhx/gopangaea/internal/pkg/device/deviceio"
	"github.com/jpoirier/gousb/usb"
)

func main() {
	ctx, done := context.WithCancel(context.Background())
	fmt.Println("#", "open connection")

	usbContext := usb.NewContext()

	dev, closeFn, err := deviceio.GetPangaeaDevice(usbContext)
	defer closeFn()
	if err != nil {
		log.Println("#", "get devices list failed:", err)
		return
	}

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

	go func() {
		defer done()

		scannerContext, stopScan := context.WithCancel(ctx)

		knownCommands := map[string]string{
			get_bank.CliCommand:          get_bank.CliDescription,
			get_device.CliCommand:        get_device.CliDescription,
			get_version.CliCommand:       get_version.CliDescription,
			get_impulse_name.CliCommand:  get_impulse_name.CliDescription,
			get_impulse_names.CliCommand: get_impulse_names.CliDescription,
			get_mode.CliCommand:          get_mode.CliDescription,
			get_settings.CliCommand:      get_settings.CliDescription,

			change_preset.CliCommand:     change_preset.CliDescription,
			set_master_volume.CliCommand: set_master_volume.CliDescription,

			set_reverb_state.CliCommand:  set_reverb_state.CliDescription,
			set_reverb_type.CliCommand:   set_reverb_type.CliDescription,
			set_reverb_volume.CliCommand: set_reverb_volume.CliDescription,

			set_presence_state.CliCommand: set_presence_state.CliDescription,
			set_presence_value.CliCommand: set_presence_value.CliDescription,

			set_lp_filter_state.CliCommand: set_lp_filter_state.CliDescription,
			set_lp_filter_value.CliCommand: set_lp_filter_value.CliDescription,
			set_hp_filter_state.CliCommand: set_hp_filter_state.CliDescription,
			set_hp_filter_value.CliCommand: set_hp_filter_value.CliDescription,

			set_impulse_state.CliCommand: set_impulse_state.CliDescription,

			exit.CliCommand: exit.CliDescription,
			info.CliCommand: info.CliDescription,
		}

		cmdParser := parser.New(map[string]command.CliCommand{
			get_bank.CliCommand:          get_bank.New(pangaea),
			get_device.CliCommand:        get_device.New(pangaea),
			get_version.CliCommand:       get_version.New(pangaea),
			get_impulse_name.CliCommand:  get_impulse_name.New(pangaea),
			get_impulse_names.CliCommand: get_impulse_names.New(pangaea),
			get_mode.CliCommand:          get_mode.New(pangaea),
			get_settings.CliCommand:      get_settings.New(pangaea),

			change_preset.CliCommand:     change_preset.New(pangaea),
			set_master_volume.CliCommand: set_master_volume.New(pangaea),

			set_reverb_state.CliCommand:  set_reverb_state.New(pangaea),
			set_reverb_type.CliCommand:   set_reverb_type.New(pangaea),
			set_reverb_volume.CliCommand: set_reverb_volume.New(pangaea),

			set_presence_state.CliCommand: set_presence_state.New(pangaea),
			set_presence_value.CliCommand: set_presence_value.New(pangaea),

			set_lp_filter_state.CliCommand: set_lp_filter_state.New(pangaea),
			set_lp_filter_value.CliCommand: set_lp_filter_value.New(pangaea),
			set_hp_filter_state.CliCommand: set_hp_filter_state.New(pangaea),
			set_hp_filter_value.CliCommand: set_hp_filter_value.New(pangaea),

			set_impulse_state.CliCommand: set_impulse_state.New(pangaea),

			exit.CliCommand:  exit.New(stopScan),
			info.CliCommand:  info.New(dev),
			usage.CliCommand: usage.New(knownCommands),
		})

		scanner := bufio.NewScanner(os.Stdin)

		readCommand := func() bool {
			fmt.Print("# ")
			return scanner.Scan()
		}

		for {
			select {
			case <-scannerContext.Done():
				return
			default:
				if !readCommand() {
					return
				}
			}

			in, err := cmdParser.Parse(scanner.Text())
			if err != nil {
				if err == parser.ErrCommandIsUnknown {
					fmt.Println("unknown command:", scanner.Text())
					fmt.Println("use 'help' for usage", scanner.Text())
				} else if err != parser.ErrCommandIsEmpty {
					fmt.Println("parse command error:", err)
				}
				continue
			}

			out, err := in.Execute()
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
