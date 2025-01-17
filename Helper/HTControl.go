package Helper

func (h *Helper) GetIRRepeaterState() (bool, error) {
	return h.hTControl.GetIRRepeaterState()
}

func (h *Helper) SetLEDFeedbackState(state bool) error {
	return h.hTControl.SetLEDFeedbackState(state)
}
