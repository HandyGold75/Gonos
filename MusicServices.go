package Gonos

type AvailableServices struct {
	DescriptorList string
	TypeList       string
	ListVersion    string
}

// Short for `zp.MusicServices.ListAvailableServices'`.
func (zp *ZonePlayer) ListAvailableServices() (AvailableServices, error) {
	res, err := zp.MusicServices.ListAvailableServices()
	return AvailableServices{
		DescriptorList: res.AvailableServiceDescriptorList,
		TypeList:       res.AvailableServiceTypeList,
		ListVersion:    res.AvailableServiceListVersion,
	}, err
}

// Short for `zp.MusicServices.UpdateAvailableServices'`.
func (zp *ZonePlayer) UpdateAvailableServices() error {
	return zp.MusicServices.UpdateAvailableServices()
}
