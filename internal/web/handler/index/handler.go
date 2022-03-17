package index

import (
	"fmt"
	"io"
	"log"
	"net/http"

	get_mode "github.com/atkhx/gopangaea/internal/pkg/commands/get-mode"
	get_settings "github.com/atkhx/gopangaea/internal/pkg/commands/get-settings"
	"github.com/atkhx/gopangaea/internal/web/form"
	"github.com/pkg/errors"
)

type Device interface {
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
	settings, err := h.device.GetSettings()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("get settings failed: %s", err)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mode, err := h.device.GetMode()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("get mode failed: %s", err)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	layoutData := map[string]interface{}{
		"title": "PANGAEA CP-100 GUI",
	}

	viewData := map[string]interface{}{
		"Settings": form.FromDto(settings.Settings()),
		"Mode":     mode.Mode,
	}

	if err := h.renderer.RenderLayoutWithView(w, "main", "index", layoutData, viewData); err != nil {
		log.Println(errors.Wrap(err, "render layout with view failed"))
	}
}
