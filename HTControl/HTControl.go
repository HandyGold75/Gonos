package HTControl

import (
	"strconv"

	"github.com/HandyGold75/Gonos/lib"
)

type (
	HTControl struct {
		Send func(action, body, targetTag string) (string, error)
	}
)

func New(send func(action, body, targetTag string) (string, error)) HTControl {
	return HTControl{Send: send}
}

func (s *HTControl) CommitLearnedIRCodes(name string) error {
	_, err := s.Send("CommitLearnedIRCodes", "<Name>"+name+"</Name>", "")
	return err
}

func (s *HTControl) GetIRRepeaterState() (bool, error) {
	res, err := s.Send("GetIRRepeaterState", "", "CurrentIRRepeaterState")
	return res == "On", err
}

func (s *HTControl) GetLEDFeedbackState() (bool, error) {
	res, err := s.Send("GetLEDFeedbackState", "", "LEDFeedbackState")
	return res == "On", err
}

func (s *HTControl) IdentifyIRRemote(timeout int) error {
	_, err := s.Send("IdentifyIRRemote", "<Timeout>"+strconv.Itoa(timeout)+"</Timeout>", "")
	return err
}

func (s *HTControl) IsRemoteConfigured(timeout int) (bool, error) {
	res, err := s.Send("IsRemoteConfigured", "", "RemoteConfigured")
	return res == "1", err
}

func (s *HTControl) LearnIRCode(iRCode string, timeout int) error {
	_, err := s.Send("LearnIRCode", "<IRCode>"+iRCode+"</IRCode><Timeout>"+strconv.Itoa(timeout)+"</Timeout>", "")
	return err
}

func (s *HTControl) SetIRRepeaterState(state bool) error {
	_, err := s.Send("SetIRRepeaterState", "<DesiredIRRepeaterState>"+lib.BoolToOnOff(state)+"</DesiredIRRepeaterState>", "")
	return err
}

func (s *HTControl) SetLEDFeedbackState(state bool) error {
	_, err := s.Send("SetLEDFeedbackState", "<LEDFeedbackState>"+lib.BoolToOnOff(state)+"</LEDFeedbackState>", "")
	return err
}
