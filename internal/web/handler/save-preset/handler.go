package save_preset

import (
	"fmt"
	"net/http"

	save_preset "github.com/atkhx/gopangaea/internal/pkg/commands/save-preset"
)

type Device interface {
	SavePreset() (save_preset.Response, error)
}

type handler struct {
	device Device
}

func New(device Device) *handler {
	return &handler{device: device}
}

func (h *handler) writeError(w http.ResponseWriter, err error) {
	w.Write([]byte(fmt.Sprintf("get settings failed: %s", err)))
	w.WriteHeader(http.StatusInternalServerError)
}

type Param struct {
	Name  string
	Value string
}

type request struct {
	Params []Param
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := h.device.SavePreset()
	if err != nil {
		h.writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
