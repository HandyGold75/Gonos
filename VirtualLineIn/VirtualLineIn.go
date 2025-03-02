package VirtualLineIn

import "strconv"

type (
	VirtualLineIn struct {
		Send func(action, body, targetTag string) (string, error)
		// Play speed usually `1`, can be a fraction of 1 Allowed values: `1`
		Speed int
	}
)

func New(send func(action, body, targetTag string) (string, error)) VirtualLineIn {
	return VirtualLineIn{Send: send, Speed: 1}
}

func (s *VirtualLineIn) Next() error {
	_, err := s.Send("Next", "", "")
	return err
}

func (s *VirtualLineIn) Pause() error {
	_, err := s.Send("Pause", "", "")
	return err
}

func (s *VirtualLineIn) Play() error {
	_, err := s.Send("Play", "<Speed>"+strconv.Itoa(s.Speed)+"</Speed>", "")
	return err
}

func (s *VirtualLineIn) Previous() error {
	_, err := s.Send("Previous", "", "")
	return err
}

func (s *VirtualLineIn) SetVolume(volume int) error {
	_, err := s.Send("SetVolume", "<DesiredVolume>"+strconv.Itoa(max(0, min(100, volume)))+"</DesiredVolume>", "")
	return err
}

func (s *VirtualLineIn) StartTransmission(coordinatorID string) (string, error) {
	return s.Send("StartTransmission", "<CoordinatorID>"+coordinatorID+"</CoordinatorID>", "CurrentTransportSettings")
}

func (s *VirtualLineIn) Stop() error {
	_, err := s.Send("Stop", "", "")
	return err
}

func (s *VirtualLineIn) StopTransmission(coordinatorID string) error {
	_, err := s.Send("StopTransmission", "<CoordinatorID>"+coordinatorID+"</CoordinatorID>", "")
	return err
}
