package index

import (
	"io"
	"log"
	"net/http"

	get_mode "github.com/atkhx/gopangaea/internal/pkg/commands/get-mode"
	get_settings "github.com/atkhx/gopangaea/internal/pkg/commands/get-settings"
	"github.com/atkhx/gopangaea/internal/web/form"
	"github.com/pkg/errors"
)

type Device interface {
	IsConnected() bool
	GetSettings() (get_settings.Response, error)
	GetMode() (get_mode.Response, error)
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
	connected := h.device.IsConnected()

	mode, err := h.device.GetMode()
	if err != nil {
		log.Println("error on get mode:", err)
	}

	viewData := map[string]interface{}{
		"Connected": connected,
		"Mode":      mode.Mode,
		"Settings":  form.Form{},
	}

	if connected {
		settings, err := h.device.GetSettings()
		if err != nil {
			log.Println("error on get settings:", err)
		}
		viewData["Settings"] = form.FromDto(settings.Settings())
	}

	layoutData := map[string]interface{}{
		"title": "PANGAEA CP-100 GUI",
	}

	if err := h.renderer.RenderLayoutWithView(w, "main", "index", layoutData, viewData); err != nil {
		log.Println(errors.Wrap(err, "render layout with view failed"))
	}
}
