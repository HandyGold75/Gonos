package Gonos

// Get ZoneGroupAttribute ZoneGroupName
func (zp *ZonePlayer) GetZoneGroupName() (string, error) {
	res, err := zp.ZoneGroupTopology.GetZoneGroupAttributes()
	return res.CurrentZoneGroupName, err
}

// Get ZoneGroupAttribute ZoneGroupID
func (zp *ZonePlayer) GetZoneGroupID() (string, error) {
	res, err := zp.ZoneGroupTopology.GetZoneGroupAttributes()
	return res.CurrentZoneGroupID, err
}

// Get ZoneGroupAttribute ZonePlayerUUIDsInGroup
func (zp *ZonePlayer) GetZonePlayerUUIDsInGroup() (string, error) {
	res, err := zp.ZoneGroupTopology.GetZoneGroupAttributes()
	return res.CurrentZonePlayerUUIDsInGroup, err
}

// Get ZoneGroupAttribute MuseHouseholdId
func (zp *ZonePlayer) GetMuseHouseholdId() (string, error) {
	res, err := zp.ZoneGroupTopology.GetZoneGroupAttributes()
	return res.CurrentMuseHouseholdId, err
}
