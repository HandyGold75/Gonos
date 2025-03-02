package Helper

// Short for `zp.SystemProperties.EnableRDM`.
func (h *Helper) EnableRDM(state bool) error {
	return h.systemProperties.EnableRDM(state)
}

// Short for `zp.SystemProperties.GetRDM`.
func (h *Helper) GetRDM() (bool, error) {
	return h.systemProperties.GetRDM()
}
