package Gonos

// Short for `zp.SystemProperties.EnableRDM`.
func (zp *ZonePlayer) EnableRDM(state bool) error {
	return zp.SystemProperties.EnableRDM(state)
}

// Short for `zp.SystemProperties.GetRDM`.
func (zp *ZonePlayer) GetRDM() (bool, error) {
	return zp.SystemProperties.GetRDM()
}
