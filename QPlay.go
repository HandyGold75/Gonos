package Gonos

// Short for `zp.QPlay.QPlayAuth`.
func (zp *ZonePlayer) QPlayAuth(seed string) error {
	_, err := zp.QPlay.QPlayAuth(seed)
	return err
}
