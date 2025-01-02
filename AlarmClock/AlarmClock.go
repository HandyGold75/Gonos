package AlarmClock

import (
	"Gonos/lib"
	"encoding/xml"
	"strconv"
	"time"
)

type (
	AlarmClock struct {
		Send func(action, body, targetTag string) (string, error)
	}

	getFormatResponse struct {
		XMLName           xml.Name `xml:"GetFormatResponse"`
		CurrentTimeFormat string
		CurrentDateFormat string
	}
	getTimeNowResponse struct {
		XMLName               xml.Name `xml:"GetTimeNowResponse"`
		CurrentUTCTime        string
		CurrentLocalTime      string
		CurrentTimeZone       string
		CurrentTimeGeneration int
	}
	getTimeZoneResponse struct {
		XMLName       xml.Name `xml:"GetTimeZoneResponse"`
		Index         int
		AutoAdjustDst bool
	}
	getTimeZoneAndRuleResponse struct {
		XMLName         xml.Name `xml:"GetTimeZoneAndRuleResponse"`
		Index           int
		AutoAdjustDst   bool
		CurrentTimeZone string
	}
	listAlarmsResponse struct {
		XMLName xml.Name `xml:"ListAlarmsResponse"`
		// xml string
		CurrentAlarmList        string
		CurrentAlarmListVersion string
	}
)

func New(send func(action, body, targetTag string) (string, error)) AlarmClock {
	return AlarmClock{Send: send}
}

// Prefer methods “
//
// `recurrence` should be one of `Gonos.RecurrenceModes.*`
//
// `playMode` should be one of `Gonos.PlayModes.*`
func (s *AlarmClock) CreateAlarm(startLocalTime time.Time, seconds int, recurrence string, enabled bool, roomUUID string, programURI string, programMetaData string, playMode string, volume int, includeLinkedZones bool) (assignedID string, err error) {
	return s.Send("CreateAlarm", "<StartLocalTime>"+startLocalTime.Format("15:04:05")+"</StartLocalTime><Duration>"+time.Time.Add(time.Time{}, time.Second*time.Duration(max(0, seconds))).Format("15:04:05")+"</Duration><Recurrence>"+recurrence+"</Recurrence><Enabled>"+lib.BoolTo10(enabled)+"</Enabled><RoomUUID>"+roomUUID+"</RoomUUID><ProgramURI>"+programURI+"</ProgramURI><ProgramMetaData>"+programMetaData+"</ProgramMetaData><PlayMode>"+playMode+"</PlayMode><Volume>"+strconv.Itoa(max(0, min(100, volume)))+"</Volume><IncludeLinkedZones>"+lib.BoolTo10(includeLinkedZones)+"</IncludeLinkedZones>", "AssignedID")
}

func (s *AlarmClock) DestroyAlarm(id int) error {
	_, err := s.Send("DestroyAlarm", "<ID>"+strconv.Itoa(id)+"</ID>", "")
	return err
}

func (s *AlarmClock) GetDailyIndexRefreshTime() (currentDailyIndexRefreshTime string, err error) {
	return s.Send("GetDailyIndexRefreshTime", "", "CurrentDailyIndexRefreshTime")
}

func (s *AlarmClock) GetFormat() (getFormatResponse, error) {
	res, err := s.Send("GetFormat", "", "s:Body")
	if err != nil {
		return getFormatResponse{}, err
	}
	data := getFormatResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *AlarmClock) GetHouseholdTimeAtStamp(timeStamp string) (householdUTCTime string, err error) {
	return s.Send("GetHouseholdTimeAtStamp", "<TimeStamp>"+timeStamp+"</TimeStamp>", "HouseholdUTCTime")
}

func (s *AlarmClock) GetTimeNow() (getTimeNowResponse, error) {
	res, err := s.Send("GetTimeNow", "", "s:Body")
	if err != nil {
		return getTimeNowResponse{}, err
	}
	data := getTimeNowResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *AlarmClock) GetTimeServer() (currentTimeServer string, err error) {
	return s.Send("GetTimeServer", "", "CurrentTimeServer")
}

func (s *AlarmClock) GetTimeZone() (getTimeZoneResponse, error) {
	res, err := s.Send("GetTimeZone", "", "s:Body")
	if err != nil {
		return getTimeZoneResponse{}, err
	}
	data := getTimeZoneResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *AlarmClock) GetTimeZoneAndRule() (getTimeZoneAndRuleResponse, error) {
	res, err := s.Send("GetTimeZoneAndRule", "", "")
	if err != nil {
		return getTimeZoneAndRuleResponse{}, err
	}
	data := getTimeZoneAndRuleResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *AlarmClock) GetTimeZoneRule(index int) (timeZone string, err error) {
	return s.Send("GetTimeZoneRule", "<Index>"+strconv.Itoa(index)+"</Index>", "TimeZone")
}

func (s *AlarmClock) ListAlarms() (listAlarmsResponse, error) {
	res, err := s.Send("ListAlarms", "", "")
	if err != nil {
		return listAlarmsResponse{}, err
	}
	data := listAlarmsResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *AlarmClock) SetDailyIndexRefreshTime(desiredDailyIndexRefreshTime string) error {
	_, err := s.Send("SetDailyIndexRefreshTime", "<DesiredDailyIndexRefreshTime>"+desiredDailyIndexRefreshTime+"</DesiredDailyIndexRefreshTime>", "")
	return err
}

func (s *AlarmClock) SetFormat(desiredTimeFormat string, desiredDateFormat string) error {
	_, err := s.Send("SetFormat", "<DesiredTimeFormat>"+desiredTimeFormat+"</DesiredTimeFormat><DesiredDateFormat>"+desiredDateFormat+"</DesiredDateFormat>", "")
	return err
}

func (s *AlarmClock) SetTimeNow(timeZoneForDesiredTime string, desiredTime string) error {
	_, err := s.Send("SetTimeNow", "<DesiredTime>"+desiredTime+"</DesiredTime><TimeZoneForDesiredTime>"+timeZoneForDesiredTime+"</TimeZoneForDesiredTime>", "")
	return err
}

func (s *AlarmClock) SetTimeServer(desiredTimeServer string) error {
	_, err := s.Send("SetTimeServer", "<DesiredTimeServer>"+desiredTimeServer+"</DesiredTimeServer>", "")
	return err
}

func (s *AlarmClock) SetTimeZone(index int, autoAdjustDst bool) error {
	_, err := s.Send("SetTimeZone", "<Index>"+strconv.Itoa(index)+"</Index><AutoAdjustDst>"+lib.BoolTo10(autoAdjustDst)+"</AutoAdjustDst>", "")

	return err
}

// `recurrence` should be one of `Gonos.RecurrenceModes.*`
// `playMode` should be one of `Gonos.PlayModes.*`
func (s *AlarmClock) UpdateAlarm(id int, startLocalTime time.Time, seconds int, recurrence string, enabled bool, roomUUID string, programURI string, programMetaData string, playMode string, volume int, includeLinkedZones bool) error {
	_, err := s.Send("UpdateAlarm", "<ID>"+strconv.Itoa(id)+"</ID><StartLocalTime>"+startLocalTime.Format("15:04:05")+"</StartLocalTime><Duration>"+time.Time.Add(time.Time{}, time.Second*time.Duration(max(0, seconds))).Format("15:04:05")+"</Duration><Recurrence>"+recurrence+"</Recurrence><Enabled>"+lib.BoolTo10(enabled)+"</Enabled><RoomUUID>"+roomUUID+"</RoomUUID><ProgramURI>"+programURI+"</ProgramURI><ProgramMetaData>"+programMetaData+"</ProgramMetaData><PlayMode>"+playMode+"</PlayMode><Volume>"+strconv.Itoa(max(0, min(100, volume)))+"</Volume><IncludeLinkedZones>"+lib.BoolTo10(includeLinkedZones)+"</IncludeLinkedZones>", "")
	return err
}
