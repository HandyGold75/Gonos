package DeviceProperties

// https://sonos.svrooij.io/services/device-properties

import (
	"Gonos/lib"
	"encoding/xml"
	"strconv"
)

type (
	DeviceProperties struct {
		Send   func(action, body, targetTag string) (string, error)
		Source string
	}

	getZoneAttributesResponse struct {
		XMLName               xml.Name `xml:"GetZoneAttributesResponse"`
		CurrentZoneName       string
		CurrentIcon           string
		CurrentConfiguration  string
		CurrentTargetRoomName string
	}
	getZoneInfoResponse struct {
		XMLName                xml.Name `xml:"GetZoneInfoResponse"`
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

func New(send func(action, body, targetTag string) (string, error)) DeviceProperties {
	return DeviceProperties{Send: send, Source: ""}
}

func (s *DeviceProperties) AddBondedZones(channelMapSet string) error {
	_, err := s.Send("AddBondedZones", "<ChannelMapSet>"+channelMapSet+"</ChannelMapSet>", "")
	return err
}

func (s *DeviceProperties) AddHTSatellite(hTSatChanMapSet string) error {
	_, err := s.Send("AddHTSatellite", "<HTSatChanMapSet>"+hTSatChanMapSet+"</HTSatChanMapSet>", "")
	return err
}

func (s *DeviceProperties) CreateStereoPair(channelMapSet string) error {
	_, err := s.Send("CreateStereoPair", "<ChannelMapSet>"+channelMapSet+"</ChannelMapSet>", "")
	return err
}

func (s *DeviceProperties) EnterConfigMode(mode string, options string) (State string, err error) {
	res, err := s.Send("EnterConfigMode", "<Mode>"+mode+"</Mode><Options>"+options+"</Options>", "State")
	return res, err
}

func (s *DeviceProperties) ExitConfigMode(options string) error {
	_, err := s.Send("ExitConfigMode", "<Options>"+options+"</Options>", "")
	return err
}

func (s *DeviceProperties) GetAutoplayLinkedZones() (IncludeLinkedZones bool, err error) {
	res, err := s.Send("GetAutoplayLinkedZones", "<Source>"+s.Source+"</Source>", "IncludeLinkedZones")
	return res == "1", err
}

func (s *DeviceProperties) GetAutoplayRoomUUID() (RoomUUID string, err error) {
	res, err := s.Send("GetAutoplayRoomUUID", "<Source>"+s.Source+"</Source>", "RoomUUID")
	return res, err
}

func (s *DeviceProperties) GetAutoplayVolume() (CurrentVolume int, err error) {
	res, err := s.Send("GetAutoplayVolume", "<Source>"+s.Source+"</Source>", "CurrentVolume")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (s *DeviceProperties) GetButtonLockState() (CurrentButtonLockState bool, err error) {
	res, err := s.Send("GetButtonLockState", "", "CurrentButtonLockState")
	return res == "On", err
}

func (s *DeviceProperties) GetButtonState() (State string, err error) {
	res, err := s.Send("GetButtonState", "", "State")
	return res, err
}

func (s *DeviceProperties) GetHouseholdID() (CurrentHouseholdID string, err error) {
	res, err := s.Send("GetHouseholdID", "", "CurrentHouseholdID")
	return res, err
}

func (s *DeviceProperties) GetHTForwardState() (IsHTForwardEnabled bool, err error) {
	res, err := s.Send("GetHTForwardState", "", "IsHTForwardEnabled")
	return res == "1", err
}

// Prefer method `h.GetLEDState`.
func (s *DeviceProperties) GetLEDState() (CurrentLEDState bool, err error) {
	res, err := s.Send("GetLEDState", "", "CurrentLEDState")
	return res == "On", err
}

func (s *DeviceProperties) GetUseAutoplayVolume() (UseVolume bool, err error) {
	res, err := s.Send("GetUseAutoplayVolume", "<Source>"+s.Source+"</Source>", "UseVolume")
	return res == "1", err
}

// Prefer methods `h.GetZoneName`, `h.GetIcon`, `h.GetConfiguration`, `h.GetTargetRoomName`.
func (s *DeviceProperties) GetZoneAttributes() (getZoneAttributesResponse, error) {
	res, err := s.Send("GetZoneAttributes", "", "s:Body")
	if err != nil {
		return getZoneAttributesResponse{}, err
	}
	data := getZoneAttributesResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

// Prefer method `h.GetZoneInfo`.
func (s *DeviceProperties) GetZoneInfo() (getZoneInfoResponse, error) {
	res, err := s.Send("GetZoneInfo", "", "s:Body")
	if err != nil {
		return getZoneInfoResponse{}, err
	}
	data := getZoneInfoResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *DeviceProperties) RemoveBondedZones(channelMapSet string, keepGrouped bool) error {
	_, err := s.Send("RemoveBondedZones", "<ChannelMapSet>"+channelMapSet+"</ChannelMapSet><KeepGrouped>"+lib.BoolTo10(keepGrouped)+"</KeepGrouped>", "")
	return err
}

func (s *DeviceProperties) RemoveHTSatellite(satRoomUUID string) error {
	_, err := s.Send("RemoveHTSatellite", "<SatRoomUUID>"+satRoomUUID+"</SatRoomUUID>", "")
	return err
}

func (s *DeviceProperties) RoomDetectionStartChirping(channel int, milliseconds int, chirpIfPlayingSwappableAudio bool) (PlayId int, err error) {
	res, err := s.Send("RoomDetectionStartChirping", "<Channel>"+strconv.Itoa(channel)+"</Channel><DurationMilliseconds>"+strconv.Itoa(milliseconds)+"</DurationMilliseconds><ChirpIfPlayingSwappableAudio>"+lib.BoolTo10(chirpIfPlayingSwappableAudio)+"</ChirpIfPlayingSwappableAudio>", "PlayId")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (s *DeviceProperties) RoomDetectionStopChirping(playId int) error {
	_, err := s.Send("RoomDetectionStopChirping", "<PlayId>"+strconv.Itoa(playId)+"</PlayId>", "")
	return err
}

func (s *DeviceProperties) SeparateStereoPair(channelMapSet string) error {
	_, err := s.Send("SeparateStereoPair", "<ChannelMapSet>"+channelMapSet+"</ChannelMapSet>", "")
	return err
}

func (s *DeviceProperties) SetAutoplayLinkedZones(includeLinkedZones bool) error {
	_, err := s.Send("SetAutoplayLinkedZones", "<IncludeLinkedZones>"+lib.BoolTo10(includeLinkedZones)+"</IncludeLinkedZones><Source>"+s.Source+"</Source>", "")
	return err
}

func (s *DeviceProperties) SetAutoplayRoomUUID(roomUUID string) error {
	_, err := s.Send("SetAutoplayRoomUUID", "<RoomUUID>"+roomUUID+"</RoomUUID><Source>"+s.Source+"</Source>", "")
	return err
}

func (s *DeviceProperties) SetAutoplayVolume(volume int) error {
	_, err := s.Send("SetAutoplayVolume", "<Volume>"+strconv.Itoa(max(0, min(100, volume)))+"</Volume><Source>"+s.Source+"</Source>", "")
	return err
}

func (s *DeviceProperties) SetButtonLockState(state bool) error {
	_, err := s.Send("SetButtonLockState", "<DesiredButtonLockState>"+lib.BoolToOnOff(state)+"</DesiredButtonLockState>", "")
	return err
}

// Prefer method `h.SetLEDState`.
func (s *DeviceProperties) SetLEDState(state bool) error {
	_, err := s.Send("SetLEDState", "<DesiredLEDState>"+lib.BoolToOnOff(state)+"</DesiredLEDState>", "")
	return err
}

func (s *DeviceProperties) SetUseAutoplayVolume(state bool) error {
	_, err := s.Send("SetUseAutoplayVolume", "<UseVolume>"+lib.BoolTo10(state)+"</UseVolume><Source>"+s.Source+"</Source>", "")
	return err
}

// Prefer methods `h.SetZoneName`, `h.SetIcon`, `h.SetConfiguration`, `h.SetTargetRoomName`.
func (s *DeviceProperties) SetZoneAttributes(zoneName string, icon string, configuration string, targetRoomName string) error {
	_, err := s.Send("SetZoneAttributes", "<DesiredZoneName>"+zoneName+"</DesiredZoneName><DesiredIcon>"+icon+"</DesiredIcon><DesiredConfiguration>"+configuration+"</DesiredConfiguration><DesiredTargetRoomName>"+targetRoomName+"</DesiredTargetRoomName>", "")
	return err
}
