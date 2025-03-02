package Gonos

type (
	ZoneInfo struct {
		SerialNumber           string
		SoftwareVersion        string
		DisplaySoftwareVersion string
		HardwareVersion        string
		IPAddress              string
		MACAddress             string
		CopyrightInfo          string
		ExtraInfo              string
		// SPDIF input, `0` not connected, `2` stereo, `7` Dolby 2.0, `18` dolby 5.1, `21` not listening, `22` silence
		HTAudioIn int
		Flags     int
	}
)

// Short for `zp.DeviceProperties.GetZoneInfo`.
func (zp *ZonePlayer) GetZoneInfo() (ZoneInfo, error) {
	info, err := zp.DeviceProperties.GetZoneInfo()
	return ZoneInfo{
		SerialNumber:           info.SerialNumber,
		SoftwareVersion:        info.SoftwareVersion,
		DisplaySoftwareVersion: info.DisplaySoftwareVersion,
		HardwareVersion:        info.HardwareVersion,
		IPAddress:              info.IPAddress,
		MACAddress:             info.MACAddress,
		CopyrightInfo:          info.CopyrightInfo,
		ExtraInfo:              info.ExtraInfo,
		HTAudioIn:              info.HTAudioIn,
		Flags:                  info.Flags,
	}, err
}

// Get ZoneAttribute ZoneName
func (zp *ZonePlayer) GetZoneName() (string, error) {
	res, err := zp.DeviceProperties.GetZoneAttributes()
	return res.CurrentZoneName, err
}

// Get ZoneAttribute Icon
func (zp *ZonePlayer) GetIcon() (string, error) {
	res, err := zp.DeviceProperties.GetZoneAttributes()
	return res.CurrentIcon, err
}

// Get ZoneAttribute Configuration
func (zp *ZonePlayer) GetConfiguration() (string, error) {
	res, err := zp.DeviceProperties.GetZoneAttributes()
	return res.CurrentConfiguration, err
}

// Get ZoneAttribute TargetRoomName
func (zp *ZonePlayer) GetTargetRoomName() (string, error) {
	res, err := zp.DeviceProperties.GetZoneAttributes()
	return res.CurrentTargetRoomName, err
}

// Set ZoneAttribute ZoneName
func (zp *ZonePlayer) SetZoneName(zoneName string) error {
	_, err := zp.DeviceProperties.Send("SetZoneAttributes", "<DesiredZoneName>"+zoneName+"</DesiredZoneName>", "")
	return err
}

// Set ZoneAttribute Icon
func (zp *ZonePlayer) SetIcon(icon string) error {
	_, err := zp.DeviceProperties.Send("SetZoneAttributes", "<DesiredIcon>"+icon+"</DesiredIcon>", "")
	return err
}

// Set ZoneAttribute Configuration
func (zp *ZonePlayer) SetConfiguration(configuration string) error {
	_, err := zp.DeviceProperties.Send("SetZoneAttributes", "<DesiredConfiguration>"+configuration+"</DesiredConfiguration>", "")
	return err
}

// Set ZoneAttribute TargetRoomName
func (zp *ZonePlayer) SetTargetRoomName(targetRoomName string) error {
	_, err := zp.DeviceProperties.Send("SetZoneAttributes", "<DesiredTargetRoomName>"+targetRoomName+"</DesiredTargetRoomName>", "")
	return err
}

// Short for `zp.DeviceProperties.GetLEDState`.
func (zp *ZonePlayer) GetLED() (bool, error) {
	return zp.DeviceProperties.GetLEDState()
}

// Short for `zp.DeviceProperties.SetLEDState`.
func (zp *ZonePlayer) SetLED(state bool) error {
	return zp.DeviceProperties.SetLEDState(state)
}
