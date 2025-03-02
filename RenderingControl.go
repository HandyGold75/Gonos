package Gonos

import (
	"Gonos/lib"
	"strconv"
)

func (zp *ZonePlayer) GetEQDialogLevel() (bool, error) {
	res, err := zp.RenderingControl.GetEQ("EQDialogLevel")
	return res == "1", err
}

func (zp *ZonePlayer) GetEQMusicSurroundLevel() (int, error) {
	res, err := zp.RenderingControl.GetEQ("EQMusicSurroundLevel")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (zp *ZonePlayer) GetEQNightMode() (bool, error) {
	res, err := zp.RenderingControl.GetEQ("EQNightMode")
	return res == "1", err
}

func (zp *ZonePlayer) GetEQSubGain() (int, error) {
	res, err := zp.RenderingControl.GetEQ("EQSubGain")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (zp *ZonePlayer) GetEQSurroundEnable() (bool, error) {
	res, err := zp.RenderingControl.GetEQ("EQSurroundEnable")
	return res == "1", err
}

func (zp *ZonePlayer) GetEQSurroundLevel() (int, error) {
	res, err := zp.RenderingControl.GetEQ("EQSurroundLevel")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (zp *ZonePlayer) GetEQSurroundMode() (bool, error) {
	res, err := zp.RenderingControl.GetEQ("EQSurroundMode")
	return res == "1", err
}

func (zp *ZonePlayer) GetEQHeightChannelLevel() (int, error) {
	res, err := zp.RenderingControl.GetEQ("EQHeightChannelLevel")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (zp *ZonePlayer) RampToVolumeSleepTimer(volume int, resetVolumeAfter bool, programURI string) (int, error) {
	return zp.RenderingControl.RampToVolume("SleepTimer", volume, resetVolumeAfter, programURI)
}

func (zp *ZonePlayer) RampToVolumeAlarm(volume int, resetVolumeAfter bool, programURI string) (int, error) {
	return zp.RenderingControl.RampToVolume("Alarm", volume, resetVolumeAfter, programURI)
}

func (zp *ZonePlayer) RampToVolumeAutoPlay(volume int, resetVolumeAfter bool, programURI string) (int, error) {
	return zp.RenderingControl.RampToVolume("AutoPlay", volume, resetVolumeAfter, programURI)
}

func (zp *ZonePlayer) ResetExtEQDialogLevel() error {
	return zp.RenderingControl.ResetExtEQ("EQDialogLevel")
}

func (zp *ZonePlayer) ResetExtEQMusicSurroundLevel() error {
	return zp.RenderingControl.ResetExtEQ("EQMusicSurroundLevel")
}

func (zp *ZonePlayer) ResetExtEQNightMode() error {
	return zp.RenderingControl.ResetExtEQ("EQNightMode")
}

func (zp *ZonePlayer) ResetExtEQSubGain() error {
	return zp.RenderingControl.ResetExtEQ("EQSubGain")
}

func (zp *ZonePlayer) ResetExtEQSurroundEnable() error {
	return zp.RenderingControl.ResetExtEQ("EQSurroundEnable")
}

func (zp *ZonePlayer) ResetExtEQSurroundLevel() error {
	return zp.RenderingControl.ResetExtEQ("EQSurroundLevel")
}

func (zp *ZonePlayer) ResetExtEQSurroundMode() error {
	return zp.RenderingControl.ResetExtEQ("EQSurroundMode")
}

func (zp *ZonePlayer) ResetExtEQHeightChannelLevel() error {
	return zp.RenderingControl.ResetExtEQ("EQHeightChannelLevel")
}

func (zp *ZonePlayer) SetEQDialogLevel(state bool) error {
	return zp.RenderingControl.SetEQ("DialogLevel", lib.BoolTo10(state))
}

func (zp *ZonePlayer) SetEQMusicSurroundLevel(volume int) error {
	return zp.RenderingControl.SetEQ("MusicSurroundLevel", strconv.Itoa(max(-15, min(15, volume))))
}

func (zp *ZonePlayer) SetEQNightMode(state bool) error {
	return zp.RenderingControl.SetEQ("NightMode", lib.BoolTo10(state))
}

func (zp *ZonePlayer) SetEQSubGain(volume int) error {
	return zp.RenderingControl.SetEQ("SubGain", strconv.Itoa(max(-10, min(10, volume))))
}

func (zp *ZonePlayer) SetEQSurroundEnable(state bool) error {
	return zp.RenderingControl.SetEQ("SurroundEnable", lib.BoolTo10(state))
}

func (zp *ZonePlayer) SetEQSurroundLevel(volume int) error {
	return zp.RenderingControl.SetEQ("SurroundLevel", strconv.Itoa(max(-15, min(15, volume))))
}

func (zp *ZonePlayer) SetEQSurroundMode(full bool) error {
	return zp.RenderingControl.SetEQ("SurroundMode", lib.BoolTo10(full))
}

func (zp *ZonePlayer) SetEQHeightChannelLevel(volume int) error {
	return zp.RenderingControl.SetEQ("HeightChannelLevel", strconv.Itoa(max(-10, min(10, volume))))
}

// Short for `zp.RenderingControl.SetLoudness`.
func (zp *ZonePlayer) SetLoudness(volume bool) error {
	return zp.RenderingControl.SetLoudness(volume)
}

// Short for `zp.RenderingControl.GetLoudness`.
func (zp *ZonePlayer) GetLoudness() (bool, error) {
	return zp.RenderingControl.GetLoudness()
}

// Short for `zp.RenderingControl.SetMute`.
func (zp *ZonePlayer) SetMute(volume bool) error {
	return zp.RenderingControl.SetMute(volume)
}

// Short for `zp.RenderingControl.GetMute`.
func (zp *ZonePlayer) GetMute() (bool, error) {
	return zp.RenderingControl.GetMute()
}

// Short for `zp.RenderingControl.SetBass`.
func (zp *ZonePlayer) SetBass(volume int) error {
	return zp.RenderingControl.SetBass(volume)
}

// Short for `zp.RenderingControl.GetBass`.
func (zp *ZonePlayer) GetBass() (int, error) {
	return zp.RenderingControl.GetBass()
}

// Short for `zp.RenderingControl.SetTreble`.
func (zp *ZonePlayer) SetTreble(volume int) error {
	return zp.RenderingControl.SetTreble(volume)
}

// Short for `zp.RenderingControl.GetTreble`.
func (zp *ZonePlayer) GetTreble() (int, error) {
	return zp.RenderingControl.GetTreble()
}

// Short for `zp.RenderingControl.SetVolume`.
func (zp *ZonePlayer) SetVolume(volume int) error {
	return zp.RenderingControl.SetVolume(volume)
}

// Short for `zp.RenderingControl.SetRelativeVolume`.
func (zp *ZonePlayer) SetVolumeDelta(volume int) (int, error) {
	return zp.RenderingControl.SetRelativeVolume(volume)
}

// Short for `zp.RenderingControl.GetVolume`.
func (zp *ZonePlayer) GetVolume() (int, error) {
	return zp.RenderingControl.GetVolume()
}

// Short for `zp.RenderingControl.SetVolumeDB`.
func (zp *ZonePlayer) SetVolumeDB(volume int) error {
	return zp.RenderingControl.SetVolumeDB(volume)
}

// Short for `zp.RenderingControl.GetVolumeDB`.
func (zp *ZonePlayer) GetVolumeDB() (int, error) {
	return zp.RenderingControl.GetVolumeDB()
}
