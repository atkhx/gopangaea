package change

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	set_compressor_state "github.com/atkhx/gopangaea/internal/pkg/commands/set-compressor-state"
	set_compressor_sustain "github.com/atkhx/gopangaea/internal/pkg/commands/set-compressor-sustain"
	set_compressor_volume "github.com/atkhx/gopangaea/internal/pkg/commands/set-compressor-volume"
	set_equalizer_frequencies "github.com/atkhx/gopangaea/internal/pkg/commands/set-equalizer-frequencies"
	set_equalizer_mixer "github.com/atkhx/gopangaea/internal/pkg/commands/set-equalizer-mixer"
	set_equalizer_position "github.com/atkhx/gopangaea/internal/pkg/commands/set-equalizer-position"
	set_equalizer_quantity_factor "github.com/atkhx/gopangaea/internal/pkg/commands/set-equalizer-quantity-factor"
	set_equalizer_state "github.com/atkhx/gopangaea/internal/pkg/commands/set-equalizer-state"
	set_hp_filter_state "github.com/atkhx/gopangaea/internal/pkg/commands/set-hp-filter-state"
	set_hp_filter_value "github.com/atkhx/gopangaea/internal/pkg/commands/set-hp-filter-value"
	set_impulse_state "github.com/atkhx/gopangaea/internal/pkg/commands/set-impulse-state"
	set_lp_filter_state "github.com/atkhx/gopangaea/internal/pkg/commands/set-lp-filter-state"
	set_lp_filter_value "github.com/atkhx/gopangaea/internal/pkg/commands/set-lp-filter-value"
	set_master_volume "github.com/atkhx/gopangaea/internal/pkg/commands/set-master-volume"
	set_mode "github.com/atkhx/gopangaea/internal/pkg/commands/set-mode"
	set_noisegate_decay "github.com/atkhx/gopangaea/internal/pkg/commands/set-noisegate-decay"
	set_noisegate_state "github.com/atkhx/gopangaea/internal/pkg/commands/set-noisegate-state"
	set_noisegate_thresh "github.com/atkhx/gopangaea/internal/pkg/commands/set-noisegate-thresh"
	set_poweramp_slave "github.com/atkhx/gopangaea/internal/pkg/commands/set-poweramp-slave"
	set_poweramp_state "github.com/atkhx/gopangaea/internal/pkg/commands/set-poweramp-state"
	set_poweramp_type "github.com/atkhx/gopangaea/internal/pkg/commands/set-poweramp-type"
	set_poweramp_volume "github.com/atkhx/gopangaea/internal/pkg/commands/set-poweramp-volume"
	set_preamp_high "github.com/atkhx/gopangaea/internal/pkg/commands/set-preamp-high"
	set_preamp_low "github.com/atkhx/gopangaea/internal/pkg/commands/set-preamp-low"
	set_preamp_mid "github.com/atkhx/gopangaea/internal/pkg/commands/set-preamp-mid"
	set_preamp_state "github.com/atkhx/gopangaea/internal/pkg/commands/set-preamp-state"
	set_preamp_volume "github.com/atkhx/gopangaea/internal/pkg/commands/set-preamp-volume"
	set_presence_state "github.com/atkhx/gopangaea/internal/pkg/commands/set-presence-state"
	set_presence_value "github.com/atkhx/gopangaea/internal/pkg/commands/set-presence-value"
	set_reverb_state "github.com/atkhx/gopangaea/internal/pkg/commands/set-reverb-state"
	set_reverb_type "github.com/atkhx/gopangaea/internal/pkg/commands/set-reverb-type"
	set_reverb_volume "github.com/atkhx/gopangaea/internal/pkg/commands/set-reverb-volume"
	"github.com/pkg/errors"
)

type Device interface {
	SetMode(value int) (set_mode.Response, error)
	SetMasterVolume(value int) (set_master_volume.Response, error)
	SetCabinetFromDevice(value int) (bool, error)

	SetReverbState(value bool) (set_reverb_state.Response, error)
	SetReverbType(value int) (set_reverb_type.Response, error)
	SetReverbVolume(value int) (set_reverb_volume.Response, error)

	SetPreampState(value bool) (set_preamp_state.Response, error)
	SetPreampVolume(volume int) (set_preamp_volume.Response, error)
	SetPreampHigh(high int) (set_preamp_high.Response, error)
	SetPreampMid(mid int) (set_preamp_mid.Response, error)
	SetPreampLow(low int) (set_preamp_low.Response, error)

	SetPowerAmpState(value bool) (set_poweramp_state.Response, error)
	SetPowerAmpType(value int) (set_poweramp_type.Response, error)
	SetPowerAmpSlave(slave int) (set_poweramp_slave.Response, error)
	SetPowerAmpVolume(volume int) (set_poweramp_volume.Response, error)

	SetPresenceState(value bool) (set_presence_state.Response, error)
	SetPresenceVolume(value int) (set_presence_value.Response, error)

	SetImpulseState(value bool) (set_impulse_state.Response, error)

	SetNoiseGateState(value bool) (set_noisegate_state.Response, error)
	SetNoiseGateThresh(value int) (set_noisegate_thresh.Response, error)
	SetNoiseGateDecay(value int) (set_noisegate_decay.Response, error)

	SetCompressorState(value bool) (set_compressor_state.Response, error)
	SetCompressorSustain(value int) (set_compressor_sustain.Response, error)
	SetCompressorVolume(value int) (set_compressor_volume.Response, error)

	SetLPFilterState(value bool) (set_lp_filter_state.Response, error)
	SetLPFilterValue(value int) (set_lp_filter_value.Response, error)
	SetHPFilterState(value bool) (set_hp_filter_state.Response, error)
	SetHPFilterValue(value int) (set_hp_filter_value.Response, error)

	SetEqualizerState(value bool) (set_equalizer_state.Response, error)
	SetEqualizerPosition(value bool) (set_equalizer_position.Response, error)
	SetEqualizerQuantityFactor(idx, value int) (set_equalizer_quantity_factor.Response, error)
	SetEqualizerMixer(idx, value int) (set_equalizer_mixer.Response, error)
	SetEqualizerFrequencies(idx, value int) (set_equalizer_frequencies.Response, error)
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
		case "cabinetType":
			return h.device.SetCabinetFromDevice(mustBeInt(param.Value))
		}
	}
	return "unknown param", nil
}
