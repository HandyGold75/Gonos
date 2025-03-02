package ConnectionManager

import "encoding/xml"

type (
	ConnectionManager struct {
		Send func(action, body, targetTag string) (string, error)
	}

	getCurrentConnectionInfoResponse struct {
		XMLName               xml.Name `xml:"GetCurrentConnectionInfoResponse"`
		RcsID                 int
		AVTransportID         int
		ProtocolInfo          string
		PeerConnectionManager string
		PeerConnectionID      int
		// Possible values: `Input`, `Output`
		Direction string
		// Possible values: `OK`, `ContentFormatMismatch`, `InsufficientBandwidth`, `UnreliableChannel`, `Unknown`
		Status string
	}
	getProtocolInfoResponse struct {
		XMLName xml.Name `xml:"GetProtocolInfoResponse"`
		Source  string
		Sink    string
	}
)

func New(send func(action, body, targetTag string) (string, error)) ConnectionManager {
	return ConnectionManager{Send: send}
}

// Prefer method `zp.GetCurrentConnectionIDs`.
func (s *ConnectionManager) GetCurrentConnectionIDs() (CurrentConnectionIDs string, err error) {
	return s.Send("GetCurrentConnectionIDs", "", "CurrentConnectionIDs")
}

// Prefer method `zp.GetCurrentConnectionInfo`.
func (s *ConnectionManager) GetCurrentConnectionInfo(connectionID string) (getCurrentConnectionInfoResponse, error) {
	res, err := s.Send("GetCurrentConnectionInfo", "<ConnectionID>"+connectionID+"</ConnectionID>", "")
	if err != nil {
		return getCurrentConnectionInfoResponse{}, err
	}
	data := getCurrentConnectionInfoResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

// Prefer method `zp.GetProtocolInfo`.
func (s *ConnectionManager) GetProtocolInfo() (getProtocolInfoResponse, error) {
	res, err := s.Send("GetProtocolInfo", "", "")
	if err != nil {
		return getProtocolInfoResponse{}, err
	}
	data := getProtocolInfoResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}
