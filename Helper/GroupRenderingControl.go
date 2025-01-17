package Helper

// Short for `zp.GroupRenderingControl.GetGroupMute'`.
func (h *Helper) GetGroupMute() (bool, error) { return h.groupRenderingControl.GetGroupMute() }

// Short for `zp.GroupRenderingControl.GetGroupVolume'`.
func (h *Helper) GetGroupVolume() (int, error) { return h.groupRenderingControl.GetGroupVolume() }

// Short for `zp.GroupRenderingControl.SetGroupMute'`.
func (h *Helper) SetGroupMute(state bool) error { return h.groupRenderingControl.SetGroupMute(state) }

// Short for `zp.GroupRenderingControl.SetGroupVolume'`.
func (h *Helper) SetGroupVolume(volume int) error {
	return h.groupRenderingControl.SetGroupVolume(volume)
}

// Short for `zp.GroupRenderingControl.SetRelativeGroupVolume'`.
func (h *Helper) SetRelativeGroupVolume(volume int) (int, error) {
	return h.groupRenderingControl.SetRelativeGroupVolume(volume)
}
