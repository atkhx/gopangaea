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
	"github.com/atkhx/gopangaea/internal/pkg/library"
	"github.com/atkhx/gopangaea/internal/web/handler/change"
	"github.com/atkhx/gopangaea/internal/web/handler/index"
	reset_preset "github.com/atkhx/gopangaea/internal/web/handler/reset-preset"
	save_preset "github.com/atkhx/gopangaea/internal/web/handler/save-preset"
	"github.com/atkhx/gopangaea/internal/web/templates"
	"github.com/jpoirier/gousb/usb"
	"github.com/pkg/errors"
)

var root = "./"

var templatePaths = []string{
	"templates/layout/*.html",
	"templates/views/*.html",
	"templates/views/*/*.html",
}

var httpHost = "localhost"
var httpPort = "8181"

func main() {
	ctx := context.Background()
	usbContext := usb.NewContext()

	conn := deviceio.New(usbContext)
	defer conn.Disconnect()

	conn.Connect()

	pangaea := device.New(conn)

	tpls, err := templates.New(root, templatePaths...)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "can't create templates"))
	}

	lib := library.New()

	//if err := lib.LoadFromDevice(pangaea); err != nil {
	//	log.Println("load from device failed", err)
	//}

	http.Handle("/", index.New(pangaea, tpls))
	http.Handle("/change", change.New(pangaea, lib, tpls))
	http.Handle("/save-preset", save_preset.New(pangaea))
	http.Handle("/reset-preset", reset_preset.New(pangaea))
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
