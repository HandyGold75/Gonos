package Gonos

import (
	"time"

	"github.com/HandyGold75/Gonos/lib"
)

// Set a alarm.
func (zp *ZonePlayer) CreateAlarm(startLocalTime time.Time, seconds int, roomUUID string, programURI string, volume int) (string, error) {
	return zp.AlarmClock.CreateAlarm(startLocalTime, seconds, lib.RecurrenceModes.Once, true, roomUUID, programURI, "", lib.PlayModes.Normal, volume, true)
}

// Update a alarm.
func (zp *ZonePlayer) UpdateAlarm(id int, startLocalTime time.Time, seconds int, roomUUID string, programURI string, volume int) error {
	return zp.AlarmClock.UpdateAlarm(id, startLocalTime, seconds, lib.RecurrenceModes.Once, true, roomUUID, programURI, "", lib.PlayModes.Normal, volume, true)
}

// Short for `zp.AlarmClock.DestroyAlarm`
func (zp *ZonePlayer) DestroyAlarm(id int) error {
	return zp.AlarmClock.DestroyAlarm(id)
}
