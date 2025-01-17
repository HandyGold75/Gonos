package Helper

type AvailableServices struct {
	DescriptorList string
	TypeList       string
	ListVersion    string
}

// Short for `zp.MusicServices.ListAvailableServices'`.
func (h *Helper) ListAvailableServices() (AvailableServices, error) {
	res, err := h.musicServices.ListAvailableServices()
	return AvailableServices{
		DescriptorList: res.AvailableServiceDescriptorList,
		TypeList:       res.AvailableServiceTypeList,
		ListVersion:    res.AvailableServiceListVersion,
	}, err
}

// Short for `zp.MusicServices.UpdateAvailableServices'`.
func (h *Helper) UpdateAvailableServices() error {
	return h.musicServices.UpdateAvailableServices()
}
