package GroupRenderingControl

import (
	"strconv"

	"github.com/HandyGold75/Gonos/lib"
)

type (
	GroupRenderingControl struct {
		Send func(action, body, targetTag string) (string, error)
	}
)

func New(send func(action, body, targetTag string) (string, error)) GroupRenderingControl {
	return GroupRenderingControl{Send: send}
}

func (s *GroupRenderingControl) GetGroupMute() (bool, error) {
	res, err := s.Send("GetGroupMute", "", "CurrentMute")
	return res == "1", err
}

func (s *GroupRenderingControl) GetGroupVolume() (int, error) {
	res, err := s.Send("GetGroupVolume", "", "CurrentVolume")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (s *GroupRenderingControl) SetGroupMute(state bool) error {
	_, err := s.Send("SetGroupMute", "<DesiredMute>"+lib.BoolTo10(state)+"</DesiredMute>", "")
	return err
}

func (s *GroupRenderingControl) SetGroupVolume(volume int) error {
	_, err := s.Send("SetGroupVolume", "<DesiredVolume>"+strconv.Itoa(max(0, min(100, volume)))+"</DesiredVolume>", "")
	return err
}

func (s *GroupRenderingControl) SetRelativeGroupVolume(volume int) (int, error) {
	res, err := s.Send("SetRelativeGroupVolume", "<Adjustment>"+strconv.Itoa(max(-100, min(100, volume)))+"</Adjustment>", "")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (s *GroupRenderingControl) SnapshotGroupVolume() error {
	_, err := s.Send("SnapshotGroupVolume", "", "")
	return err
}
