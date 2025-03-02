package Helper



// Get ZoneGroupAttribute CurrentZoneGroupName 
func (h *Helper) GetCurrentZoneGroupName () (string, error) {
	res, err := h.zoneGroupTopology.GetZoneGroupAttributes() 
	return res.CurrentZoneGroupName , err
}

// Get ZoneGroupAttribute CurrentZoneGroupID 
func (h *Helper) GetCurrentZoneGroupID () (string, error) {
	res, err := h.zoneGroupTopology.GetZoneGroupAttributes() 
	return res.CurrentZoneGroupID , err
}

// Get ZoneGroupAttribute CurrentZonePlayerUUIDsInGroup 
func (h *Helper) GetCurrentZonePlayerUUIDsInGroup () (string, error) {
	res, err := h.zoneGroupTopology.GetZoneGroupAttributes() 
	return res.CurrentZonePlayerUUIDsInGroup , err
}

// Get ZoneGroupAttribute CurrentMuseHouseholdId
func (h *Helper) GetCurrentMuseHouseholdId() (string, error) {
	res, err := h.zoneGroupTopology.GetZoneGroupAttributes() 
	return res.CurrentMuseHouseholdId, err
}




