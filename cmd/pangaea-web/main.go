package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/atkhx/gopangaea/internal/pkg/device"
	"github.com/atkhx/gopangaea/internal/pkg/device/deviceio"
	"github.com/atkhx/gopangaea/internal/pkg/web/templates"
	"github.com/atkhx/gopangaea/internal/web/handler/change"
	"github.com/atkhx/gopangaea/internal/web/handler/index"
	"github.com/jpoirier/gousb/usb"
	"github.com/pkg/errors"
)

var root = "./"

var templatePaths = []string{
	"templates/layout/*.html",
	"templates/views/*.html",
}

var httpHost = "localhost"
var httpPort = "8181"

func main() {
	ctx := context.Background()
	//ctx, done := context.WithCancel(context.Background())
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

	tpls, err := templates.New(root, templatePaths...)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "can't create templates"))
	}

	http.Handle("/", index.New(pangaea, tpls))
	http.Handle("/change", change.New(pangaea, tpls))
	http.Handle("/static/", http.FileServer(http.Dir(root)))

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf("%s:%s", httpHost, httpPort), nil); err != nil {
			log.Fatalln(errors.Wrap(err, "start server failed"))
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
