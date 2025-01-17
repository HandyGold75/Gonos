package MusicServices

import (
	"encoding/xml"
	"strconv"
)

type (
	MusicServices struct {
		Send func(action, body, targetTag string) (string, error)
	}

	listAvailableServicesResponse struct {
		XMLName                        xml.Name `xml:"ListAvailableServicesResponse"`
		AvailableServiceDescriptorList string
		AvailableServiceTypeList       string
		AvailableServiceListVersion    string
	}
)

func New(send func(action, body, targetTag string) (string, error)) MusicServices {
	return MusicServices{Send: send}
}

func (s *MusicServices) GetSessionId(serviceId int, username string) (string, error) {
	return s.Send("GetSessionId", "<ServiceId>"+strconv.Itoa(serviceId)+"</ServiceId><Username>"+username+"</Username>", "SessionId")
}

func (s *MusicServices) ListAvailableServices() (listAvailableServicesResponse, error) {
	res, err := s.Send("ListAvailableServices", "", "s:Body")
	if err != nil {
		return listAvailableServicesResponse{}, err
	}
	data := listAvailableServicesResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *MusicServices) UpdateAvailableServices() error {
	_, err := s.Send("UpdateAvailableServices", "", "")
	return err
}
