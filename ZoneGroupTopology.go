package Gonos

// Get current zone group name.
func (zp *ZonePlayer) GetZoneGroupName() (string, error) {
	res, err := zp.ZoneGroupTopology.GetZoneGroupAttributes()
	return res.CurrentZoneGroupName, err
}

// Get current zone group ID.
func (zp *ZonePlayer) GetZoneGroupID() (string, error) {
	res, err := zp.ZoneGroupTopology.GetZoneGroupAttributes()
	return res.CurrentZoneGroupID, err
}

// Get current UUIDs of all zone players in the current group.
func (zp *ZonePlayer) GetZonePlayerUUIDsInGroup() (string, error) {
	res, err := zp.ZoneGroupTopology.GetZoneGroupAttributes()
	return res.CurrentZonePlayerUUIDsInGroup, err
}

// Get current Muse household ID.
func (zp *ZonePlayer) GetMuseHouseholdId() (string, error) {
	res, err := zp.ZoneGroupTopology.GetZoneGroupAttributes()
	return res.CurrentMuseHouseholdId, err
}
