package index

import (
	"io"
	"log"
	"net/http"

	get_bank "github.com/atkhx/gopangaea/internal/pkg/commands/get-bank"
	get_device "github.com/atkhx/gopangaea/internal/pkg/commands/get-device"
	get_impulse_name "github.com/atkhx/gopangaea/internal/pkg/commands/get-impulse-name"
	get_impulse_names "github.com/atkhx/gopangaea/internal/pkg/commands/get-impulse-names"
	get_mode "github.com/atkhx/gopangaea/internal/pkg/commands/get-mode"
	get_settings "github.com/atkhx/gopangaea/internal/pkg/commands/get-settings"
	get_version "github.com/atkhx/gopangaea/internal/pkg/commands/get-version"
	"github.com/atkhx/gopangaea/internal/pkg/library/settings"
	"github.com/pkg/errors"
)

type Device interface {
	IsConnected() bool
	GetSettings() (get_settings.Response, error)
	GetMode() (get_mode.Response, error)
	GetDevice() (get_device.Response, error)
	GetVersion() (get_version.Response, error)
	GetBank() (get_bank.Response, error)
	GetImpulseNames() (get_impulse_names.Response, error)
	GetImpulseName() (get_impulse_name.Response, error)
}

type Renderer interface {
	RenderLayoutWithView(w io.Writer, layout, view string, layoutData, viewData map[string]interface{}) error
}

type handler struct {
	device   Device
	renderer Renderer
}

func New(device Device, renderer Renderer) *handler {
	return &handler{device: device, renderer: renderer}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mode, err := h.device.GetMode()
	if err != nil {
		log.Println("error on get mode:", err)
	}

	device, err := h.device.GetDevice()
	if err != nil {
		log.Println("error on get device:", err)
	}

	version, err := h.device.GetVersion()
	if err != nil {
		log.Println("error on get version:", err)
	}

	bank, err := h.device.GetBank()
	if err != nil {
		log.Println("error on get bank:", err)
	}

	impulseName, err := h.device.GetImpulseName()
	if err != nil {
		log.Println("error on get impulse name:", err)
	}

	connected := h.device.IsConnected()

	viewData := map[string]interface{}{
		"Connected":       connected,
		"Mode":            mode.Mode,
		"ModeName":        mode.String(),
		"Device":          device.String(),
		"Bank":            bank.Bank,
		"Preset":          bank.Preset,
		"PresetSelection": GetPresetSelection(bank.Bank, bank.Preset),
		"ImpulseName":     impulseName.String(),
		"Version":         version.String(),
		"AmpTypes":        AmpTypes{},
		"Settings":        settings.Settings{},
		"DeviceImpulses":  DeviceImpulses{},
	}

	if connected {
		dto, err := h.device.GetSettings()
		if err != nil {
			log.Println("error on get settings:", err)
		} else {
			settingsObj := dto.Settings()
			viewData["Settings"] = settingsObj
			viewData["AmpTypes"] = GetAmpTypes(version.String(), settingsObj.PowerAmp.Index)
		}

		impulses, err := h.device.GetImpulseNames()
		if err != nil {
			log.Println("error on get impulses:", err)
		} else {
			viewData["DeviceImpulses"] = DeviceImpulsesFromDto(impulses)
		}
	}

	layoutData := map[string]interface{}{
		"title": "PANGAEA CP-100 GUI",
	}

	if err := h.renderer.RenderLayoutWithView(w, "main", "index", layoutData, viewData); err != nil {
		log.Println(errors.Wrap(err, "render layout with view failed"))
	}
}
