package Gonos

// Short for `zp.GroupRenderingControl.GetGroupMute'`.
func (zp *ZonePlayer) GetGroupMute() (bool, error) { return zp.GroupRenderingControl.GetGroupMute() }

// Short for `zp.GroupRenderingControl.GetGroupVolume'`.
func (zp *ZonePlayer) GetGroupVolume() (int, error) { return zp.GroupRenderingControl.GetGroupVolume() }

// Short for `zp.GroupRenderingControl.SetGroupMute'`.
func (zp *ZonePlayer) SetGroupMute(state bool) error {
	return zp.GroupRenderingControl.SetGroupMute(state)
}

// Short for `zp.GroupRenderingControl.SetGroupVolume'`.
func (zp *ZonePlayer) SetGroupVolume(volume int) error {
	return zp.GroupRenderingControl.SetGroupVolume(volume)
}

// Short for `zp.GroupRenderingControl.SetRelativeGroupVolume'`.
func (zp *ZonePlayer) SetRelativeGroupVolume(volume int) (int, error) {
	return zp.GroupRenderingControl.SetRelativeGroupVolume(volume)
}
