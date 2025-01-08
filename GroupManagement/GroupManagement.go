package GroupManagement

import (
	"encoding/xml"
	"strconv"
)

type (
	GroupManagement struct {
		Send func(action, body, targetTag string) (string, error)
	}

	addMemberResponse struct {
		XMLName                  xml.Name `xml:"AddMemberResponse"`
		CurrentTransportSettings string
		CurrentURI               string
		GroupUUIDJoined          string
		ResetVolumeAfter         bool
		VolumeAVTransportURI     string
	}
)

func New(send func(action, body, targetTag string) (string, error)) GroupManagement {
	return GroupManagement{Send: send}
}

func (s *GroupManagement) AddMember(memberID string, bootSeq int) (addMemberResponse, error) {
	res, err := s.Send("AddMember", "<MemberID>"+memberID+"</MemberID><BootSeq>"+strconv.Itoa(bootSeq)+"</BootSeq>", "s:Body")
	if err != nil {
		return addMemberResponse{}, err
	}
	data := addMemberResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *GroupManagement) RemoveMember(memberID string) error {
	_, err := s.Send("RemoveMember", "<MemberID>"+memberID+"</MemberID>", "")
	return err
}

func (s *GroupManagement) ReportTrackBufferingResult(memberID string, resultCode int) error {
	_, err := s.Send("ReportTrackBufferingResult", "<MemberID>"+memberID+"</MemberID><ResultCode>"+strconv.Itoa(resultCode)+"</ResultCode>", "")
	return err
}

func (s *GroupManagement) SetSourceAreaIds(desiredSourceAreaIds string) error {
	_, err := s.Send("SetSourceAreaIds", "<DesiredSourceAreaIds>"+desiredSourceAreaIds+"</DesiredSourceAreaIds>", "")
	return err
}
