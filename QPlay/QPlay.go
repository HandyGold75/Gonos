package QPlay

import "encoding/xml"

type (
	QPlay struct {
		Send func(action, body, targetTag string) (string, error)
	}

	qPlayAuthResponse struct {
		XMLName xml.Name `xml:"QPlayAuthResponse "`
		Code    string
		MID     string
		DID     string
	}
)

func New(send func(action, body, targetTag string) (string, error)) QPlay {
	return QPlay{Send: send}
}

func (s *QPlay) QPlayAuth(seed string) (qPlayAuthResponse, error) {
	res, err := s.Send("GetSessionId", "<Seed>"+seed+"</Seed>", "s:Body")
	if err != nil {
		return qPlayAuthResponse{}, err
	}
	data := qPlayAuthResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}
