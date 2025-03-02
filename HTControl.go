package Gonos

// Short for `zp.HTControl.GetIRRepeaterState'`.
func (zp *ZonePlayer) GetIRRepeaterState() (bool, error) {
	return zp.HTControl.GetIRRepeaterState()
}

// Short for `zp.HTControl.SetLEDFeedbackState'`.
func (zp *ZonePlayer) SetLEDFeedbackState(state bool) error {
	return zp.HTControl.SetLEDFeedbackState(state)
}
