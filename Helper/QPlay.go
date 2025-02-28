package Helper

// Short for `zp.QPlay.QPlayAuth`.
func (h *Helper) QPlayAuth(seed string) error {
	_, err := h.qPlay.QPlayAuth(seed)
	return err
}
