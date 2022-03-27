package change

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	get_bank "github.com/atkhx/gopangaea/internal/pkg/commands/get-bank"
	"github.com/atkhx/gopangaea/internal/pkg/library"
	"github.com/atkhx/gopangaea/internal/pkg/library/impulse"
	"github.com/pkg/errors"
)

type Device interface {
	SetMode(value int) (bool, error)
	SetMasterVolume(value int) (bool, error)
	SetCabinetFromDevice(value int) (bool, error)

	SetReverbState(value bool) (bool, error)
	SetReverbType(value int) (bool, error)
	SetReverbVolume(value int) (bool, error)

	SetPreampState(value bool) (bool, error)
	SetPreampVolume(volume int) (bool, error)
	SetPreampHigh(high int) (bool, error)
	SetPreampMid(mid int) (bool, error)
	SetPreampLow(low int) (bool, error)

	SetPowerAmpState(value bool) (bool, error)
	SetPowerAmpType(value int) (bool, error)
	SetPowerAmpSlave(slave int) (bool, error)
	SetPowerAmpVolume(volume int) (bool, error)

	SetPresenceState(value bool) (bool, error)
	SetPresenceVolume(value int) (bool, error)

	SetImpulseState(value bool) (bool, error)

	SetNoiseGateState(value bool) (bool, error)
	SetNoiseGateThresh(value int) (bool, error)
	SetNoiseGateDecay(value int) (bool, error)

	SetCompressorState(value bool) (bool, error)
	SetCompressorSustain(value int) (bool, error)
	SetCompressorVolume(value int) (bool, error)

	SetLPFilterState(value bool) (bool, error)
	SetLPFilterValue(value int) (bool, error)
	SetHPFilterState(value bool) (bool, error)
	SetHPFilterValue(value int) (bool, error)

	SetEqualizerState(value bool) (bool, error)
	SetEqualizerPosition(value bool) (bool, error)
	SetEqualizerQuantityFactor(idx, value int) (bool, error)
	SetEqualizerMixer(idx, value int) (bool, error)
	SetEqualizerFrequencies(idx, value int) (bool, error)

	SetImpulse(name string, persist bool, impulse *impulse.Impulse) (bool, error)

	GetBank() (get_bank.Response, error)
	ChangePreset(bank, preset int) (bool, error)
}

type Library interface {
	GetPreset(index int) (library.Preset, error)
}

type Renderer interface {
	RenderLayoutWithView(w io.Writer, layout, view string, layoutData, viewData map[string]interface{}) error
}

type handler struct {
	device   Device
	library  Library
	renderer Renderer
}

func New(device Device, library Library, renderer Renderer) *handler {
	return &handler{device: device, library: library, renderer: renderer}
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
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.writeError(w, errors.Wrap(err, "read body failed"))
		return
	}

	var req request
	if err := json.Unmarshal(data, &req); err != nil {
		h.writeError(w, errors.Wrap(err, "unmarshal params failed"))
		return
	}

	result, err := h.changeParams(req.Params)

	w.Write([]byte("\n\n"))
	w.Write([]byte(fmt.Sprintf("result: %s\n", result)))
	w.Write([]byte(fmt.Sprintf("error: %s\n", err)))
	w.WriteHeader(http.StatusOK)
}

var re = regexp.MustCompile(`(.*?)(\d+)$`)

func (h *handler) changeParams(params []Param) (interface{}, error) {
	mustBeInt := func(v string) int {
		value, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		return value
	}

	mustBeBool := func(v string) bool {
		return v == "true"
	}

	for _, param := range params {
		var name = param.Name
		var idx = 0

		m := re.FindStringSubmatch(param.Name)
		if len(m) > 1 {
			name = m[1]
			i, err := strconv.Atoi(m[2])
			if err != nil {
				return nil, err
			}

			idx = i
		}

		switch name {

		case "reverbType":
			return h.device.SetReverbType(mustBeInt(param.Value))
		case "reverbVolume":
			return h.device.SetReverbVolume(mustBeInt(param.Value))
		case "reverbEnabled":
			return h.device.SetReverbState(mustBeBool(param.Value))

		case "presenceEnabled":
			return h.device.SetPresenceState(mustBeBool(param.Value))
		case "presenceValue":
			return h.device.SetPresenceVolume(mustBeInt(param.Value))

		case "powerAmpType":
			return h.device.SetPowerAmpType(mustBeInt(param.Value))
		case "powerAmpEnabled":
			return h.device.SetPowerAmpState(mustBeBool(param.Value))
		case "powerAmpVolume":
			return h.device.SetPowerAmpVolume(mustBeInt(param.Value))
		case "powerAmpSlave":
			return h.device.SetPowerAmpSlave(mustBeInt(param.Value))

		case "preampEnabled":
			return h.device.SetPreampState(mustBeBool(param.Value))
		case "preampVolume":
			return h.device.SetPreampVolume(mustBeInt(param.Value))
		case "preampHigh":
			return h.device.SetPreampHigh(mustBeInt(param.Value))
		case "preampMid":
			return h.device.SetPreampMid(mustBeInt(param.Value))
		case "preampLow":
			return h.device.SetPreampLow(mustBeInt(param.Value))

		case "cabinetEnabled":
			return h.device.SetImpulseState(mustBeBool(param.Value))

		case "noiseGateEnabled":
			return h.device.SetNoiseGateState(mustBeBool(param.Value))
		case "noiseGateThresh":
			return h.device.SetNoiseGateThresh(mustBeInt(param.Value))
		case "noiseGateDecay":
			return h.device.SetNoiseGateDecay(mustBeInt(param.Value))

		case "compressorEnabled":
			return h.device.SetCompressorState(mustBeBool(param.Value))
		case "compressorSustain":
			return h.device.SetCompressorSustain(mustBeInt(param.Value))
		case "compressorVolume":
			return h.device.SetCompressorVolume(mustBeInt(param.Value))

		case "highPassFilterActive":
			return h.device.SetHPFilterState(mustBeBool(param.Value))
		case "highPassFilterValue":
			return h.device.SetHPFilterValue(mustBeInt(param.Value))

		case "lowPassFilterActive":
			return h.device.SetLPFilterState(mustBeBool(param.Value))
		case "lowPassFilterValue":
			return h.device.SetLPFilterValue(mustBeInt(param.Value))

		case "equalizerEnabled":
			return h.device.SetEqualizerState(mustBeBool(param.Value))
		case "equalizerPosition":
			return h.device.SetEqualizerPosition(mustBeBool(param.Value))

		case "qualityFactor":
			return h.device.SetEqualizerQuantityFactor(idx, mustBeInt(param.Value))
		case "mixer":
			return h.device.SetEqualizerMixer(idx, mustBeInt(param.Value))
		case "frequencies":
			return h.device.SetEqualizerFrequencies(idx, mustBeInt(param.Value))

		case "mode":
			return h.device.SetMode(mustBeInt(param.Value))
		case "masterVolume":
			return h.device.SetMasterVolume(mustBeInt(param.Value))
		case "bank":
			b, err := h.device.GetBank()
			if err != nil {
				return nil, err
			}

			return h.device.ChangePreset(mustBeInt(param.Value), b.Preset)
		case "preset":
			b, err := h.device.GetBank()
			if err != nil {
				return nil, err
			}

			return h.device.ChangePreset(b.Bank, mustBeInt(param.Value))
		case "cabinetType":
			preset, err := h.library.GetPreset(mustBeInt(param.Value))
			if err != nil {
				return nil, err
			}
			return h.device.SetImpulse(preset.Name, false, preset.Impulse)
		}
	}
	return "unknown param", nil
}
