package Gonos

// GetZoneGroupName returns the current zone group name.
func (zp *ZonePlayer) GetZoneGroupName() (string, error) {
	res, err := zp.ZoneGroupTopology.GetZoneGroupAttributes()
	return res.CurrentZoneGroupName, err
}

// GetZoneGroupID returns the current zone group ID.
func (zp *ZonePlayer) GetZoneGroupID() (string, error) {
	res, err := zp.ZoneGroupTopology.GetZoneGroupAttributes()
	return res.CurrentZoneGroupID, err
}

// GetZonePlayerUUIDsInGroup returns the UUIDs of all zone players in the current group.
func (zp *ZonePlayer) GetZonePlayerUUIDsInGroup() (string, error) {
	res, err := zp.ZoneGroupTopology.GetZoneGroupAttributes()
	return res.CurrentZonePlayerUUIDsInGroup, err
}

// GetMuseHouseholdId returns the Muse household ID.
func (zp *ZonePlayer) GetMuseHouseholdId() (string, error) {
	res, err := zp.ZoneGroupTopology.GetZoneGroupAttributes()
	return res.CurrentMuseHouseholdId, err
}
