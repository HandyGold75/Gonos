package Helper

// Short for `zp.HTControl.GetIRRepeaterState'`.
func (h *Helper) GetIRRepeaterState() (bool, error) {
	return h.hTControl.GetIRRepeaterState()
}

// Short for `zp.HTControl.SetLEDFeedbackState'`.
func (h *Helper) SetLEDFeedbackState(state bool) error {
	return h.hTControl.SetLEDFeedbackState(state)
}
