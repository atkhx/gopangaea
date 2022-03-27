package device

import (
	"log"
	"time"

	"github.com/atkhx/gopangaea/internal/pkg/commands/change-preset"
	get_bank "github.com/atkhx/gopangaea/internal/pkg/commands/get-bank"
	get_device "github.com/atkhx/gopangaea/internal/pkg/commands/get-device"
	get_impulse "github.com/atkhx/gopangaea/internal/pkg/commands/get-impulse"
	get_impulse_name "github.com/atkhx/gopangaea/internal/pkg/commands/get-impulse-name"
	get_impulse_names "github.com/atkhx/gopangaea/internal/pkg/commands/get-impulse-names"
	get_mode "github.com/atkhx/gopangaea/internal/pkg/commands/get-mode"
	get_settings "github.com/atkhx/gopangaea/internal/pkg/commands/get-settings"
	get_version "github.com/atkhx/gopangaea/internal/pkg/commands/get-version"
	reset_preset "github.com/atkhx/gopangaea/internal/pkg/commands/reset-preset"
	save_preset "github.com/atkhx/gopangaea/internal/pkg/commands/save-preset"
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
	set_impulse "github.com/atkhx/gopangaea/internal/pkg/commands/set-impulse"
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
	set_settings "github.com/atkhx/gopangaea/internal/pkg/commands/set-settings"
	"github.com/atkhx/gopangaea/internal/pkg/device/deviceio"
	"github.com/atkhx/gopangaea/internal/pkg/library/impulse"
	"github.com/pkg/errors"
)

type Command interface {
	GetCommand() string
	GetResponseLength() int
}

type CommandSetter interface {
	GetCommand() string
}

type Validator interface {
	Validate() error
}

type Connection interface {
	IsConnected() bool
	Connect() error
	Disconnect() error
	Exec(command string, responseLength int) ([]byte, error)
}

type Response interface {
	GetLength() int
	Parse([]byte) error
}

type device struct {
	connection Connection
}

func New(connection Connection) *device {
	return &device{connection: connection}
}

func (d *device) IsConnected() bool {
	return d.connection.IsConnected()
}

func (d *device) ValidateCommand(command interface{}) error {
	if v, ok := command.(Validator); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (d *device) execSetter(command CommandSetter, suffix ...[]byte) (bool, error) {
	response := deviceio.NewResponse(suffix...)
	b, err := d.execCommand(NewSetter(command, response))
	if err != nil {
		return false, err
	}

	if err := response.Parse(b); err != nil {
		return false, err
	}

	return response.Success(), nil
}

func (d *device) execCommand(command Command) ([]byte, error) {
	if !d.IsConnected() {
		if err := d.connection.Connect(); err != nil {
			return nil, err
		}
	}

	if v, ok := command.(Validator); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	t := time.Now()
	defer func() {
		log.Println(string([]byte(command.GetCommand())[:3]), "to read response:", time.Now().Sub(t))
	}()

	return d.connection.Exec(command.GetCommand(), command.GetResponseLength())
}

func (d *device) GetDevice() (get_device.Response, error) {
	b, err := d.execCommand(get_device.Command{})
	if err != nil {
		return get_device.Response{}, err
	}
	return get_device.ParseResponse(b)
}

func (d *device) GetVersion() (get_version.Response, error) {
	b, err := d.execCommand(get_version.Command{})
	if err != nil {
		return get_version.Response{}, err
	}
	return get_version.ParseResponse(b)
}

func (d *device) GetBank() (get_bank.Response, error) {
	b, err := d.execCommand(get_bank.Command{})
	if err != nil {
		return get_bank.Response{}, err
	}
	return get_bank.ParseResponse(b)
}

func (d *device) SetCabinetFromDevice(value int) (bool, error) {
	settings, err := d.GetSettings()
	if err != nil {
		return false, errors.Wrap(err, "get settings failed")
	}

	bank, err := d.GetBank()
	if err != nil {
		return false, errors.Wrap(err, "get bank failed")
	}

	bankIndex := value / 10
	presetIndex := value - 10*bankIndex

	if _, err := d.ChangePreset(bankIndex, presetIndex); err != nil {
		return false, errors.Wrap(err, "change preset failed")
	}

	if _, err := d.SetSettings(settings.Bytes); err != nil {
		return false, errors.Wrap(err, "set settings failed")
	}

	impulse, err := d.GetImpulse()
	if err != nil {
		return false, errors.Wrap(err, "get impulse failed")
	}

	if _, err := d.ChangePreset(bank.Bank, bank.Preset); err != nil {
		return false, errors.Wrap(err, "change preset bak failed")
	}

	if _, err := d.SetImpulse(impulse.Name, false, impulse.Impulse); err != nil {
		return false, errors.Wrap(err, "set impulse failed")
	}

	if _, err := d.SetSettings(settings.Bytes); err != nil {
		return false, errors.Wrap(err, "set settings failed")
	}

	return true, nil
}

func (d *device) GetImpulseName() (get_impulse_name.Response, error) {
	b, err := d.execCommand(get_impulse_name.Command{})
	if err != nil {
		return get_impulse_name.Response{}, err
	}
	return get_impulse_name.ParseResponse(b)
}

func (d *device) GetImpulse() (get_impulse.Response, error) {
	b, err := d.execCommand(get_impulse.Command{})
	if err != nil {
		return get_impulse.Response{}, err
	}
	return get_impulse.ParseResponse(b)
}

func (d *device) SetImpulse(name string, persist bool, impulse *impulse.Impulse) (bool, error) {
	return d.execSetter(set_impulse.New(name, persist, impulse), set_impulse.ResponseSuffix)
}

func (d *device) GetMode() (get_mode.Response, error) {
	b, err := d.execCommand(get_mode.Command{})
	if err != nil {
		return get_mode.Response{}, err
	}
	return get_mode.ParseResponse(b)
}

func (d *device) GetSettings() (get_settings.Response, error) {
	b, err := d.execCommand(get_settings.Command{})
	if err != nil {
		return get_settings.Response{}, err
	}
	return get_settings.ParseResponse(b)
}

func (d *device) SetSettings(settings []byte) (bool, error) {
	return d.execSetter(set_settings.Command{Settings: settings}, set_settings.ResponseSuffix)
}

var impulsesCacheResponse *get_impulse_names.Response = nil

func (d *device) GetImpulseNames() (get_impulse_names.Response, error) {
	if impulsesCacheResponse != nil {
		return *impulsesCacheResponse, nil
	}

	b, err := d.execCommand(get_impulse_names.Command{})
	if err != nil {
		return get_impulse_names.Response{}, err
	}

	r, err := get_impulse_names.ParseResponse(b)
	if err != nil {
		return r, err
	}
	//return r, err

	impulsesCacheResponse = &r
	return *impulsesCacheResponse, nil
}

func (d *device) ResetPreset() (bool, error) {
	return d.execSetter(reset_preset.New())
}

func (d *device) ChangePreset(bank, preset int) (bool, error) {
	return d.execSetter(change_preset.New(bank, preset))
}

func (d *device) SetMode(value int) (bool, error) {
	return d.execSetter(set_mode.New(value), []byte{})
}

func (d *device) SetMasterVolume(value int) (bool, error) {
	return d.execSetter(set_master_volume.New(value))
}

func (d *device) SetReverbState(value bool) (bool, error) {
	return d.execSetter(set_reverb_state.New(value))
}

func (d *device) SetReverbType(value int) (bool, error) {
	return d.execSetter(set_reverb_type.New(value))
}

func (d *device) SetReverbVolume(value int) (bool, error) {
	return d.execSetter(set_reverb_volume.New(value))
}

func (d *device) SetPresenceState(value bool) (bool, error) {
	return d.execSetter(set_presence_state.New(value))
}

func (d *device) SetPresenceVolume(value int) (bool, error) {
	return d.execSetter(set_presence_value.New(value))
}

func (d *device) SetLPFilterState(value bool) (bool, error) {
	return d.execSetter(set_lp_filter_state.New(value))
}

func (d *device) SetLPFilterValue(value int) (bool, error) {
	return d.execSetter(set_lp_filter_value.New(value))
}

func (d *device) SetHPFilterState(value bool) (bool, error) {
	return d.execSetter(set_hp_filter_state.New(value))
}

func (d *device) SetHPFilterValue(value int) (bool, error) {
	return d.execSetter(set_hp_filter_value.New(value))
}

func (d *device) SetImpulseState(value bool) (bool, error) {
	return d.execSetter(set_impulse_state.New(value))
}

func (d *device) SetPreampState(value bool) (bool, error) {
	return d.execSetter(set_preamp_state.New(value))
}

func (d *device) SetPreampVolume(volume int) (bool, error) {
	return d.execSetter(set_preamp_volume.New(volume))
}

func (d *device) SetPreampHigh(high int) (bool, error) {
	return d.execSetter(set_preamp_high.New(high))
}

func (d *device) SetPreampMid(mid int) (bool, error) {
	return d.execSetter(set_preamp_mid.New(mid))
}

func (d *device) SetPreampLow(low int) (bool, error) {
	return d.execSetter(set_preamp_low.New(low))
}

func (d *device) SetPowerAmpState(value bool) (bool, error) {
	return d.execSetter(set_poweramp_state.New(value))
}

func (d *device) SetPowerAmpType(value int) (bool, error) {
	return d.execSetter(set_poweramp_type.New(value))
}

func (d *device) SetPowerAmpVolume(volume int) (bool, error) {
	return d.execSetter(set_poweramp_volume.New(volume))
}

func (d *device) SetPowerAmpSlave(slave int) (bool, error) {
	return d.execSetter(set_poweramp_slave.New(slave))
}

func (d *device) SetCompressorState(value bool) (bool, error) {
	return d.execSetter(set_compressor_state.New(value))
}

func (d *device) SetCompressorSustain(value int) (bool, error) {
	return d.execSetter(set_compressor_sustain.New(value))
}

func (d *device) SetCompressorVolume(value int) (bool, error) {
	return d.execSetter(set_compressor_volume.New(value))
}

func (d *device) SetNoiseGateState(value bool) (bool, error) {
	return d.execSetter(set_noisegate_state.New(value))
}

func (d *device) SetNoiseGateThresh(value int) (bool, error) {
	return d.execSetter(set_noisegate_thresh.New(value))
}

func (d *device) SetNoiseGateDecay(value int) (bool, error) {
	return d.execSetter(set_noisegate_decay.New(value))
}

func (d *device) SetEqualizerState(value bool) (bool, error) {
	return d.execSetter(set_equalizer_state.New(value), []byte{})
}

func (d *device) SetEqualizerPosition(value bool) (bool, error) {
	return d.execSetter(set_equalizer_position.New(value))
}

func (d *device) SetEqualizerQuantityFactor(idx, value int) (bool, error) {
	return d.execSetter(set_equalizer_quantity_factor.New(idx, value), []byte{})
}

func (d *device) SetEqualizerFrequencies(idx, value int) (bool, error) {
	return d.execSetter(set_equalizer_frequencies.New(idx, value), []byte{})
}

func (d *device) SetEqualizerMixer(idx, value int) (bool, error) {
	return d.execSetter(set_equalizer_mixer.New(idx, value), []byte{})
}

func (d *device) SavePreset() (bool, error) {
	return d.execSetter(save_preset.New())
}
