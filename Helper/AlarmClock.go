package Helper

import (
	"Gonos/lib"
	"time"
)

// Set a alarm.
func (h *Helper) CreateAlarm(startLocalTime time.Time, seconds int, roomUUID string, programURI string, volume int) (string, error) {
	return h.alarmClock.CreateAlarm(startLocalTime, seconds, lib.RecurrenceModes.Once, true, roomUUID, programURI, "", lib.PlayModes.Normal, volume, true)
}

// Update a alarm.
func (h *Helper) UpdateAlarm(id int, startLocalTime time.Time, seconds int, roomUUID string, programURI string, volume int) error {
	return h.alarmClock.UpdateAlarm(id, startLocalTime, seconds, lib.RecurrenceModes.Once, true, roomUUID, programURI, "", lib.PlayModes.Normal, volume, true)
}

// Short for `zp.AlarmClock.DestroyAlarm`
func (h *Helper) DestroyAlarm(id int) error {
	return h.alarmClock.DestroyAlarm(id)
}
