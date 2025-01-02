package RenderingControl

// https://sonos.svrooij.io/services/device-properties

import (
	"Gonos/lib"
	"encoding/xml"
	"strconv"
)

type (
	RenderingControl struct {
		Send func(action, body, targetTag string) (string, error)
		// Channel should be one of: `Master`, `LF` or `RF`
		Channel string
	}

	getRoomCalibrationStatusResponse struct {
		XMLName                  xml.Name `xml:"GetRoomCalibrationStatusResponse"`
		RoomCalibrationEnabled   bool
		RoomCalibrationAvailable bool
	}

	getVolumeDBRangeResponse struct {
		XMLName  xml.Name `xml:"GetVolumeDBRangeResponse"`
		MinValue int
		MaxValue int
	}

	resetBasicEQResponse struct {
		XMLName     xml.Name `xml:"ResetBasicEQResponse"`
		Bass        int
		Treble      int
		Loudness    bool
		LeftVolume  int
		RightVolume int
	}
)

func New(send func(action, body, targetTag string) (string, error)) RenderingControl {
	return RenderingControl{Send: send, Channel: "Master"}
}

func (s *RenderingControl) GetBass() (CurrentBass int, err error) {
	res, err := s.Send("GetBass", "", "CurrentBass")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (s *RenderingControl) GetEQ(eQType string) (CurrentValue string, err error) {
	return s.Send("GetEQ", "<EQType>"+eQType+"</EQType>", "CurrentValue")
}

func (s *RenderingControl) GetHeadphoneConnected() (CurrentHeadphoneConnected bool, err error) {
	res, err := s.Send("GetHeadphoneConnected", "", "CurrentHeadphoneConnected")
	return res == "1", err
}

func (s *RenderingControl) GetLoudness() (CurrentLoudness bool, err error) {
	res, err := s.Send("GetLoudness", "<Channel>"+s.Channel+"</Channel>", "CurrentLoudness")
	return res == "1", err
}

func (s *RenderingControl) GetMute() (CurrentMute bool, err error) {
	res, err := s.Send("GetMute", "<Channel>"+s.Channel+"</Channel>", "CurrentMute")
	return res == "1", err
}

func (s *RenderingControl) GetOutputFixed() (CurrentFixed bool, err error) {
	res, err := s.Send("GetOutputFixed", "", "CurrentFixed")
	return res == "1", err
}

func (s *RenderingControl) GetRoomCalibrationStatus() (getRoomCalibrationStatusResponse, error) {
	res, err := s.Send("GetRoomCalibrationStatus", "", "s:Body")
	if err != nil {
		return getRoomCalibrationStatusResponse{}, err
	}
	data := getRoomCalibrationStatusResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *RenderingControl) GetSupportsOutputFixed() (CurrentSupportsFixed bool, err error) {
	res, err := s.Send("GetSupportsOutputFixed", "", "CurrentSupportsFixed")
	return res == "1", err
}

func (s *RenderingControl) GetTreble() (CurrentTreble int, err error) {
	res, err := s.Send("GetTreble", "", "CurrentTreble")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (s *RenderingControl) GetVolume() (CurrentVolume int, err error) {
	res, err := s.Send("GetVolume", "<Channel>"+s.Channel+"</Channel>", "CurrentVolume")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (s *RenderingControl) GetVolumeDB() (CurrentVolume int, err error) {
	res, err := s.Send("GetVolumeDB", "<Channel>"+s.Channel+"</Channel>", "CurrentVolume")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (s *RenderingControl) GetVolumeDBRange() (getVolumeDBRangeResponse, error) {
	res, err := s.Send("GetVolumeDBRange", "<Channel>"+s.Channel+"</Channel>", "s:Body")
	if err != nil {
		return getVolumeDBRangeResponse{}, err
	}
	data := getVolumeDBRangeResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *RenderingControl) RampToVolume(rampType string, volume int, resetVolumeAfter bool, programURI string) (RampTime int, err error) {
	res, err := s.Send("RampToVolume", "<Channel>"+s.Channel+"</Channel><RampType>"+rampType+"</RampType><DesiredVolume>"+strconv.Itoa(max(0, min(100, volume)))+"</DesiredVolume><ResetVolumeAfter>"+lib.BoolTo10(resetVolumeAfter)+"</ResetVolumeAfter><ProgramURI>"+programURI+"</ProgramURI>", "RampTime")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (s *RenderingControl) ResetBasicEQ() (resetBasicEQResponse, error) {
	res, err := s.Send("ResetBasicEQ", "", "s:Body")
	if err != nil {
		return resetBasicEQResponse{}, err
	}
	data := resetBasicEQResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *RenderingControl) ResetExtEQ(eQType string) error {
	_, err := s.Send("ResetExtEQ", "<EQType>"+eQType+"</EQType>", "")
	return err
}

func (s *RenderingControl) RestoreVolumePriorToRamp() error {
	_, err := s.Send("RestoreVolumePriorToRamp", "<Channel>"+s.Channel+"</Channel>", "")
	return err
}

func (s *RenderingControl) SetBass(volume int) error {
	_, err := s.Send("SetBass", "<DesiredBass>"+strconv.Itoa(max(-10, min(10, volume)))+"</DesiredBass>", "")
	return err
}

func (s *RenderingControl) SetChannelMap(channelMap string) error {
	_, err := s.Send("SetChannelMap", "<ChannelMap>"+channelMap+"</ChannelMap>", "")
	return err
}

func (s *RenderingControl) SetEQ(eQType string, state string) error {
	_, err := s.Send("SetEQ", "<EQType>"+eQType+"</EQType><DesiredValue>"+state+"</DesiredValue>", "")
	return err
}

func (s *RenderingControl) SetLoudness(state bool) error {
	_, err := s.Send("SetLoudness", "<Channel>"+s.Channel+"</Channel><DesiredLoudness>"+lib.BoolTo10(state)+"</DesiredLoudness>", "")
	return err
}

func (s *RenderingControl) SetMute(state bool) error {
	_, err := s.Send("SetMute", "<Channel>"+s.Channel+"</Channel><DesiredMute>"+lib.BoolTo10(state)+"</DesiredMute>", "")
	return err
}

func (s *RenderingControl) SetOutputFixed(state bool) error {
	_, err := s.Send("SetOutputFixed", "<DesiredFixed>"+lib.BoolTo10(state)+"</DesiredFixed>", "")
	return err
}

func (s *RenderingControl) SetRelativeVolume(volume int) (NewVolume int, err error) {
	res, err := s.Send("SetRelativeVolume", "<Channel>"+s.Channel+"</Channel><Adjustment>"+strconv.Itoa(max(-100, min(100, volume)))+"</Adjustment>", "NewVolume")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (s *RenderingControl) SetRoomCalibrationStatus(state bool) error {
	_, err := s.Send("SetRoomCalibrationStatus", "<RoomCalibrationEnabled>"+lib.BoolTo10(state)+"</RoomCalibrationEnabled>", "")
	return err
}

func (s *RenderingControl) SetRoomCalibrationX(calibrationID string, coeddicients string, calibrationMode string) error {
	_, err := s.Send("SetRoomCalibrationX", "<CalibrationID>"+calibrationID+"</CalibrationID><Coefficients>"+coeddicients+"</Coefficients><CalibrationMode>"+calibrationMode+"</CalibrationMode>", "")
	return err
}

func (s *RenderingControl) SetTreble(volume int) error {
	_, err := s.Send("SetTreble", "<DesiredTreble>"+strconv.Itoa(max(-10, min(10, volume)))+"</DesiredTreble>", "")
	return err
}

func (s *RenderingControl) SetVolume(volume int) error {
	_, err := s.Send("SetVolume", "<Channel>"+s.Channel+"</Channel><DesiredVolume>"+strconv.Itoa(max(0, min(100, volume)))+"</DesiredVolume>", "")
	return err
}

func (s *RenderingControl) SetVolumeDB(volume int) error {
	_, err := s.Send("SetVolumeDB", "<Channel>"+s.Channel+"</Channel><DesiredVolume>"+strconv.Itoa(max(0, min(100, volume)))+"</DesiredVolume>", "")
	return err
}
