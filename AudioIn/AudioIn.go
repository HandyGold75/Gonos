package AudioIn

import (
	"encoding/xml"
	"strconv"
)

type (
	AudioIn struct {
		Send func(action, body, targetTag string) (string, error)
	}

	getAudioInputAttributesResponse struct {
		XMLName     xml.Name `xml:"GetAudioInputAttributesResponse"`
		CurrentName string
		CurrentIcon string
	}
	getLineInLevelResponse struct {
		XMLName                 xml.Name `xml:"GetLineInLevelResponse"`
		CurrentLeftLineInLevel  int
		CurrentRightLineInLevel int
	}
)

func New(send func(action, body, targetTag string) (string, error)) AudioIn {
	return AudioIn{Send: send}
}

func (s *AudioIn) GetAudioInputAttributes() (getAudioInputAttributesResponse, error) {
	res, err := s.Send("GetAudioInputAttributes", "", "")
	if err != nil {
		return getAudioInputAttributesResponse{}, err
	}
	data := getAudioInputAttributesResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

// Prefer method `h.GetLineInLevel`, `h.GetLineInLevelLeft`, `h.GetLineInLevelRight`.
func (s *AudioIn) GetLineInLevel() (getLineInLevelResponse, error) {
	res, err := s.Send("GetLineInLevel", "", "")
	if err != nil {
		return getLineInLevelResponse{}, err
	}
	data := getLineInLevelResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

// `objectID` may be one of `Gonos.ContentTypes.*` or a custom id
func (s *AudioIn) SelectAudio(objectID string) error {
	_, err := s.Send("SelectAudio", "<ObjectID>"+objectID+"</ObjectID>", "")
	return err
}

func (s *AudioIn) SetAudioInputAttributes(desiredName, desiredIcon string) error {
	_, err := s.Send("SetAudioInputAttributes", "<DesiredName>"+desiredName+"</DesiredName><DesiredIcon>"+desiredIcon+"</DesiredIcon>", "")
	return err
}

// Prefer methods `h.SetLineInLevel`, `h.SetLineInLevelLeft`, `h.SetLineInLevelRight`.
func (s *AudioIn) SetLineInLevel(desiredLeftLineInLevel, desiredRightLineInLevel int) error {
	_, err := s.Send("SetLineInLevel", "<DesiredLeftLineInLevel>"+strconv.Itoa(max(0, min(100, desiredLeftLineInLevel)))+"</DesiredLeftLineInLevel><DesiredRightLineInLevel>"+strconv.Itoa(max(0, min(100, desiredRightLineInLevel)))+"</DesiredRightLineInLevel>", "")
	return err
}

func (s *AudioIn) StartTransmissionToGroup(coordinatorID string) (CurrentTransportSettings string, err error) {
	return s.Send("StartTransmissionToGroup", "<CoordinatorID>"+coordinatorID+"</CoordinatorID>", "CurrentTransportSettings")
}

func (s *AudioIn) StopTransmissionToGroup(coordinatorID string) error {
	_, err := s.Send("StopTransmissionToGroup", "<CoordinatorID>"+coordinatorID+"</CoordinatorID>", "")
	return err
}
