package ZoneGroupTopology

import (
	"Gonos/lib"
	"encoding/xml"
	"strconv"
)

type (
	ZoneGroupTopology struct {
		Send func(action, body, targetTag string) (string, error)
	}

	getZoneGroupAttributesResponse struct {
		XMLName                       xml.Name `xml:"GetZoneGroupAttributesResponse"`
		CurrentZoneGroupName          string
		CurrentZoneGroupID            string
		CurrentZonePlayerUUIDsInGroup string
		CurrentMuseHouseholdId        string
	}
)

func New(send func(action, body, targetTag string) (string, error)) ZoneGroupTopology {
	return ZoneGroupTopology{Send: send}
}

func (s *ZoneGroupTopology) BeginSoftwareUpdate(updateURL string, flags int, extraOptions string) error {
	_, err := s.Send("BeginSoftwareUpdate", "<UpdateURL>"+updateURL+"</UpdateURL><Flags>"+strconv.Itoa(flags)+"</Flags><ExtraOptions>"+extraOptions+"</ExtraOptions>", "")
	return err
}

// `updateType` may be one of `Gonos.UpdateTypes.*`.
func (s *ZoneGroupTopology) CheckForUpdate(updateType string, cachedOnly bool, version string) (string, error) {
	return s.Send("CheckForUpdate", "<UpdateType>"+updateType+"</UpdateType><CachedOnly>"+lib.BoolTo10(cachedOnly)+"</CachedOnly><Version>"+version+"</Version>", "UpdateItem")
}

func (s *ZoneGroupTopology) GetZoneGroupAttributes() (getZoneGroupAttributesResponse, error) {
	res, err := s.Send("GetZoneGroupAttributes", "", "s:Body")
	if err != nil {
		return getZoneGroupAttributesResponse{}, err
	}
	data := getZoneGroupAttributesResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *ZoneGroupTopology) GetZoneGroupState() (string, error) {
	return s.Send("GetZoneGroupState", "", "ZoneGroupState")
}

func (s *ZoneGroupTopology) RegisterMobileDevice(mobileDeviceName string, mobileDeviceUDN string, mobileIPAndPort string) error {
	_, err := s.Send("RegisterMobileDevice", "<MobileDeviceName>"+mobileDeviceName+"</MobileDeviceName><MobileDeviceUDN>"+mobileDeviceUDN+"</MobileDeviceUDN><MobileIPAndPort>"+mobileIPAndPort+"</MobileIPAndPort>", "")
	return err
}

func (s *ZoneGroupTopology) ReportAlarmStartedRunning() error {
	_, err := s.Send("ReportAlarmStartedRunning", "", "")
	return err
}

// `desiredAction` Allowed values: `Remove` / `TopologyMonitorProbe` / `VerifyThenRemoveSystemwide`
func (s *ZoneGroupTopology) ReportUnresponsiveDevice(deviceUUID string, desiredAction string) error {
	_, err := s.Send("ReportUnresponsiveDevice", "<DeviceUUID>"+deviceUUID+"</DeviceUUID><DesiredAction>"+desiredAction+"</DesiredAction>", "")
	return err
}

func (s *ZoneGroupTopology) SubmitDiagnostics(includeControllers bool, typ string) (string, error) {
	return s.Send("SubmitDiagnostics", "<IncludeControllers>"+lib.BoolTo10(includeControllers)+"</IncludeControllers><Type>"+typ+"</Type>", "DiagnosticID")
}
