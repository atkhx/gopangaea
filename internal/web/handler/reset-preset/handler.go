package reset_preset

import (
	"fmt"
	"net/http"

	change_preset "github.com/atkhx/gopangaea/internal/pkg/commands/change-preset"
	get_bank "github.com/atkhx/gopangaea/internal/pkg/commands/get-bank"
)

type Device interface {
	GetBank() (get_bank.Response, error)
	ChangePreset(bank, preset int) (change_preset.Response, error)
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
	b, err := h.device.GetBank()
	if err != nil {
		h.writeError(w, err)
		return
	}

	_, err = h.device.ChangePreset(b.Bank, b.Preset)
	if err != nil {
		h.writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
