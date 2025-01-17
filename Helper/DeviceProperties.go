package Helper

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
func (h *Helper) GetZoneInfo() (ZoneInfo, error) {
	info, err := h.deviceProperties.GetZoneInfo()
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
func (h *Helper) GetZoneName() (string, error) {
	res, err := h.deviceProperties.GetZoneAttributes()
	return res.CurrentZoneName, err
}

// Get ZoneAttribute Icon
func (h *Helper) GetIcon() (string, error) {
	res, err := h.deviceProperties.GetZoneAttributes()
	return res.CurrentIcon, err
}

// Get ZoneAttribute Configuration
func (h *Helper) GetConfiguration() (string, error) {
	res, err := h.deviceProperties.GetZoneAttributes()
	return res.CurrentConfiguration, err
}

// Get ZoneAttribute TargetRoomName
func (h *Helper) GetTargetRoomName() (string, error) {
	res, err := h.deviceProperties.GetZoneAttributes()
	return res.CurrentTargetRoomName, err
}

// Set ZoneAttribute ZoneName
func (h *Helper) SetZoneName(zoneName string) error {
	_, err := h.deviceProperties.Send("SetZoneAttributes", "<DesiredZoneName>"+zoneName+"</DesiredZoneName>", "")
	return err
}

// Set ZoneAttribute Icon
func (h *Helper) SetIcon(icon string) error {
	_, err := h.deviceProperties.Send("SetZoneAttributes", "<DesiredIcon>"+icon+"</DesiredIcon>", "")
	return err
}

// Set ZoneAttribute Configuration
func (h *Helper) SetConfiguration(configuration string) error {
	_, err := h.deviceProperties.Send("SetZoneAttributes", "<DesiredConfiguration>"+configuration+"</DesiredConfiguration>", "")
	return err
}

// Set ZoneAttribute TargetRoomName
func (h *Helper) SetTargetRoomName(targetRoomName string) error {
	_, err := h.deviceProperties.Send("SetZoneAttributes", "<DesiredTargetRoomName>"+targetRoomName+"</DesiredTargetRoomName>", "")
	return err
}

// Short for `zp.DeviceProperties.GetLEDState`.
func (h *Helper) GetLED() (bool, error) {
	return h.deviceProperties.GetLEDState()
}

// Short for `zp.DeviceProperties.SetLEDState`.
func (h *Helper) SetLED(state bool) error {
	return h.deviceProperties.SetLEDState(state)
}
