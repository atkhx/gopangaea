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
	"github.com/pkg/errors"
)

type Command interface {
	GetCommand() string
	GetResponseLength() int
}

type Connection interface {
	IsConnected() bool
	Connect() error
	Disconnect() error
	WriteCommand(command string) error
	ReadResponse(command string, length int) ([]byte, error)
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

func (d *device) ExecCommand(command Command) ([]byte, error) {
	return d.execCommand(command.GetCommand(), command.GetResponseLength())
}

func (d *device) execCommand(command string, responseLength int) ([]byte, error) {
	if !d.IsConnected() {
		if err := d.connection.Connect(); err != nil {
			return nil, err
		}
	}

	if err := d.connection.WriteCommand(command); err != nil {
		return nil, err
	}
	t := time.Now()
	defer func() {
		log.Println("time to read response:", time.Now().Sub(t))
	}()
	return d.connection.ReadResponse(command, responseLength)
}

func (d *device) GetDevice() (get_device.Response, error) {
	b, err := d.ExecCommand(get_device.Command{})
	if err != nil {
		return get_device.Response{}, err
	}
	return get_device.ParseResponse(b)
}

func (d *device) GetVersion() (get_version.Response, error) {
	b, err := d.ExecCommand(get_version.Command{})
	if err != nil {
		return get_version.Response{}, err
	}
	return get_version.ParseResponse(b)
}

func (d *device) GetBank() (get_bank.Response, error) {
	b, err := d.ExecCommand(get_bank.Command{})
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

	log.Println("settings", settings)
	log.Println("currentBank", bank)
	log.Println("bankIndex", bankIndex)
	log.Println("presetIndex", presetIndex)

	if _, err := d.ChangePreset(bankIndex, presetIndex); err != nil {
		return false, errors.Wrap(err, "change preset failed")
	}

	impulse, err := d.GetImpulse()
	if err != nil {
		return false, errors.Wrap(err, "get impulse failed")
	}

	log.Println("target impulse:", impulse.Name)

	if _, err := d.ChangePreset(bank.Bank, bank.Preset); err != nil {
		return false, errors.Wrap(err, "change preset bak failed")
	}

	if _, err := d.SetImpulse(impulse.Name, impulse.Impulse); err != nil {
		return false, errors.Wrap(err, "set impulse failed")
	}

	if _, err := d.SetSettings(settings.Bytes); err != nil {
		return false, errors.Wrap(err, "set settings failed")
	}

	return true, nil
}

func (d *device) GetImpulseName() (get_impulse_name.Response, error) {
	b, err := d.ExecCommand(get_impulse_name.Command{})
	if err != nil {
		return get_impulse_name.Response{}, err
	}
	return get_impulse_name.ParseResponse(b)
}

func (d *device) GetImpulse() (get_impulse.Response, error) {
	b, err := d.ExecCommand(get_impulse.Command{})
	if err != nil {
		return get_impulse.Response{}, err
	}
	return get_impulse.ParseResponse(b)
}

func (d *device) SetImpulse(name string, impulse []byte) (set_impulse.Response, error) {
	cmd := set_impulse.Command{Name: name, Impulse: impulse}
	b, err := d.ExecCommand(cmd)
	if err != nil {
		return set_impulse.Response{}, err
	}
	return set_impulse.ParseResponse(b)
}

func (d *device) GetMode() (get_mode.Response, error) {
	b, err := d.ExecCommand(get_mode.Command{})
	if err != nil {
		return get_mode.Response{}, err
	}
	return get_mode.ParseResponse(b)
}

func (d *device) GetSettings() (get_settings.Response, error) {
	b, err := d.ExecCommand(get_settings.Command{})
	if err != nil {
		return get_settings.Response{}, err
	}
	return get_settings.ParseResponse(b)
}

func (d *device) SetSettings(settings []byte) (set_settings.Response, error) {
	b, err := d.ExecCommand(set_settings.Command{Settings: settings})
	if err != nil {
		return set_settings.Response{}, err
	}
	return set_settings.ParseResponse(b)
}

var impulsesCacheResponse *get_impulse_names.Response = nil

func (d *device) GetImpulseNames() (get_impulse_names.Response, error) {
	if impulsesCacheResponse != nil {
		return *impulsesCacheResponse, nil
	}

	b, err := d.ExecCommand(get_impulse_names.Command{})
	if err != nil {
		return get_impulse_names.Response{}, err
	}

	r, err := get_impulse_names.ParseResponse(b)
	if err != nil {
		return r, err
	}
	return r, err

	//impulsesCacheResponse = &r
	//return *impulsesCacheResponse, nil
}

func (d *device) ChangePreset(bank, preset int) (change_preset.Response, error) {
	command, err := change_preset.NewWithArgs(bank, preset)
	if err != nil {
		return change_preset.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return change_preset.Response{}, err
	}

	return change_preset.ParseResponse(b)
}

func (d *device) SetMode(value int) (set_mode.Response, error) {
	command, err := set_mode.NewWithArgs(value)
	if err != nil {
		return set_mode.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_mode.Response{}, err
	}

	return set_mode.ParseResponse(b)
}

func (d *device) SetMasterVolume(value int) (set_master_volume.Response, error) {
	command, err := set_master_volume.NewWithArgs(value)
	if err != nil {
		return set_master_volume.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_master_volume.Response{}, err
	}

	return set_master_volume.ParseResponse(b)
}

func (d *device) SetReverbState(value bool) (set_reverb_state.Response, error) {
	command := set_reverb_state.NewWithArgs(value)
	b, err := d.ExecCommand(command)
	if err != nil {
		return set_reverb_state.Response{}, err
	}

	return set_reverb_state.ParseResponse(b)
}

func (d *device) SetReverbType(value int) (set_reverb_type.Response, error) {
	command, err := set_reverb_type.NewWithArgs(value)
	if err != nil {
		return set_reverb_type.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_reverb_type.Response{}, err
	}

	return set_reverb_type.ParseResponse(b)
}

func (d *device) SetReverbVolume(value int) (set_reverb_volume.Response, error) {
	command, err := set_reverb_volume.NewWithArgs(value)
	if err != nil {
		return set_reverb_volume.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_reverb_volume.Response{}, err
	}

	return set_reverb_volume.ParseResponse(b)
}

func (d *device) SetPresenceState(value bool) (set_presence_state.Response, error) {
	command := set_presence_state.NewWithArgs(value)
	b, err := d.ExecCommand(command)
	if err != nil {
		return set_presence_state.Response{}, err
	}

	return set_presence_state.ParseResponse(b)
}

func (d *device) SetPresenceVolume(value int) (set_presence_value.Response, error) {
	command, err := set_presence_value.NewWithArgs(value)
	if err != nil {
		return set_presence_value.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_presence_value.Response{}, err
	}

	return set_presence_value.ParseResponse(b)
}

func (d *device) SetLPFilterState(value bool) (set_lp_filter_state.Response, error) {
	command := set_lp_filter_state.NewWithArgs(value)
	b, err := d.ExecCommand(command)
	if err != nil {
		return set_lp_filter_state.Response{}, err
	}

	return set_lp_filter_state.ParseResponse(b)
}

func (d *device) SetLPFilterValue(value int) (set_lp_filter_value.Response, error) {
	command, err := set_lp_filter_value.NewWithArgs(value)
	if err != nil {
		return set_lp_filter_value.Response{}, err
	}
	b, err := d.ExecCommand(command)
	if err != nil {
		return set_lp_filter_value.Response{}, err
	}

	return set_lp_filter_value.ParseResponse(b)
}

func (d *device) SetHPFilterState(value bool) (set_hp_filter_state.Response, error) {
	command := set_hp_filter_state.NewWithArgs(value)
	b, err := d.ExecCommand(command)
	if err != nil {
		return set_hp_filter_state.Response{}, err
	}

	return set_hp_filter_state.ParseResponse(b)
}

func (d *device) SetHPFilterValue(value int) (set_hp_filter_value.Response, error) {
	command, err := set_hp_filter_value.NewWithArgs(value)
	if err != nil {
		return set_hp_filter_value.Response{}, err
	}
	b, err := d.ExecCommand(command)
	if err != nil {
		return set_hp_filter_value.Response{}, err
	}

	return set_hp_filter_value.ParseResponse(b)
}

func (d *device) SetImpulseState(value bool) (set_impulse_state.Response, error) {
	command := set_impulse_state.NewWithArgs(value)
	b, err := d.ExecCommand(command)
	if err != nil {
		return set_impulse_state.Response{}, err
	}

	return set_impulse_state.ParseResponse(b)
}

func (d *device) SetPreampState(value bool) (set_preamp_state.Response, error) {
	command := set_preamp_state.NewWithArgs(value)
	b, err := d.ExecCommand(command)
	if err != nil {
		return set_preamp_state.Response{}, err
	}

	return set_preamp_state.ParseResponse(b)
}

func (d *device) SetPreampVolume(volume int) (set_preamp_volume.Response, error) {
	command, err := set_preamp_volume.NewWithArgs(volume)
	if err != nil {
		return set_preamp_volume.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_preamp_volume.Response{}, err
	}

	return set_preamp_volume.ParseResponse(b)
}

func (d *device) SetPreampHigh(high int) (set_preamp_high.Response, error) {
	command, err := set_preamp_high.NewWithArgs(high)
	if err != nil {
		return set_preamp_high.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_preamp_high.Response{}, err
	}

	return set_preamp_high.ParseResponse(b)
}

func (d *device) SetPreampMid(mid int) (set_preamp_mid.Response, error) {
	command, err := set_preamp_mid.NewWithArgs(mid)
	if err != nil {
		return set_preamp_mid.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_preamp_mid.Response{}, err
	}

	return set_preamp_mid.ParseResponse(b)
}

func (d *device) SetPreampLow(low int) (set_preamp_low.Response, error) {
	command, err := set_preamp_low.NewWithArgs(low)
	if err != nil {
		return set_preamp_low.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_preamp_low.Response{}, err
	}

	return set_preamp_low.ParseResponse(b)
}

func (d *device) SetPowerAmpState(value bool) (set_poweramp_state.Response, error) {
	command := set_poweramp_state.NewWithArgs(value)
	b, err := d.ExecCommand(command)
	if err != nil {
		return set_poweramp_state.Response{}, err
	}

	return set_poweramp_state.ParseResponse(b)
}

func (d *device) SetPowerAmpType(value int) (set_poweramp_type.Response, error) {
	command, err := set_poweramp_type.NewWithArgs(value)
	if err != nil {
		return set_poweramp_type.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_poweramp_type.Response{}, err
	}

	return set_poweramp_type.ParseResponse(b)
}

func (d *device) SetPowerAmpVolume(volume int) (set_poweramp_volume.Response, error) {
	command, err := set_poweramp_volume.NewWithArgs(volume)
	if err != nil {
		return set_poweramp_volume.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_poweramp_volume.Response{}, err
	}

	return set_poweramp_volume.ParseResponse(b)
}

func (d *device) SetPowerAmpSlave(slave int) (set_poweramp_slave.Response, error) {
	command, err := set_poweramp_slave.NewWithArgs(slave)
	if err != nil {
		return set_poweramp_slave.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_poweramp_slave.Response{}, err
	}

	return set_poweramp_slave.ParseResponse(b)
}

func (d *device) SetCompressorState(value bool) (set_compressor_state.Response, error) {
	command := set_compressor_state.NewWithArgs(value)
	b, err := d.ExecCommand(command)
	if err != nil {
		return set_compressor_state.Response{}, err
	}

	return set_compressor_state.ParseResponse(b)
}

func (d *device) SetCompressorSustain(value int) (set_compressor_sustain.Response, error) {
	command, err := set_compressor_sustain.NewWithArgs(value)
	if err != nil {
		return set_compressor_sustain.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_compressor_sustain.Response{}, err
	}

	return set_compressor_sustain.ParseResponse(b)
}

func (d *device) SetCompressorVolume(value int) (set_compressor_volume.Response, error) {
	command, err := set_compressor_volume.NewWithArgs(value)
	if err != nil {
		return set_compressor_volume.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_compressor_volume.Response{}, err
	}

	return set_compressor_volume.ParseResponse(b)
}

func (d *device) SetNoiseGateState(value bool) (set_noisegate_state.Response, error) {
	command := set_noisegate_state.NewWithArgs(value)
	b, err := d.ExecCommand(command)
	if err != nil {
		return set_noisegate_state.Response{}, err
	}

	return set_noisegate_state.ParseResponse(b)
}

func (d *device) SetNoiseGateThresh(value int) (set_noisegate_thresh.Response, error) {
	command, err := set_noisegate_thresh.NewWithArgs(value)
	if err != nil {
		return set_noisegate_thresh.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_noisegate_thresh.Response{}, err
	}

	return set_noisegate_thresh.ParseResponse(b)
}

func (d *device) SetNoiseGateDecay(value int) (set_noisegate_decay.Response, error) {
	command, err := set_noisegate_decay.NewWithArgs(value)
	if err != nil {
		return set_noisegate_decay.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_noisegate_decay.Response{}, err
	}

	return set_noisegate_decay.ParseResponse(b)
}

func (d *device) SetEqualizerState(value bool) (set_equalizer_state.Response, error) {
	command := set_equalizer_state.NewWithArgs(value)
	b, err := d.ExecCommand(command)
	if err != nil {
		return set_equalizer_state.Response{}, err
	}

	return set_equalizer_state.ParseResponse(b)
}

func (d *device) SetEqualizerPosition(value bool) (set_equalizer_position.Response, error) {
	command := set_equalizer_position.NewWithArgs(value)
	b, err := d.ExecCommand(command)
	if err != nil {
		return set_equalizer_position.Response{}, err
	}

	return set_equalizer_position.ParseResponse(b)
}

func (d *device) SetEqualizerQuantityFactor(idx, value int) (set_equalizer_quantity_factor.Response, error) {
	command, err := set_equalizer_quantity_factor.NewWithArgs(idx, value)
	if err != nil {
		return set_equalizer_quantity_factor.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_equalizer_quantity_factor.Response{}, err
	}

	return set_equalizer_quantity_factor.ParseResponse(b)
}

func (d *device) SetEqualizerFrequencies(idx, value int) (set_equalizer_frequencies.Response, error) {
	command, err := set_equalizer_frequencies.NewWithArgs(idx, value)
	if err != nil {
		return set_equalizer_frequencies.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_equalizer_frequencies.Response{}, err
	}

	return set_equalizer_frequencies.ParseResponse(b)
}

func (d *device) SetEqualizerMixer(idx, value int) (set_equalizer_mixer.Response, error) {
	command, err := set_equalizer_mixer.NewWithArgs(idx, value)
	if err != nil {
		return set_equalizer_mixer.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return set_equalizer_mixer.Response{}, err
	}

	return set_equalizer_mixer.ParseResponse(b)
}

func (d *device) SavePreset() (save_preset.Response, error) {
	command := save_preset.New()

	b, err := d.ExecCommand(command)
	if err != nil {
		return save_preset.Response{}, err
	}

	return save_preset.ParseResponse(b)
}
