package Gonos

// Short for `zp.VirtualLineIn.Next`.
func (zp *ZonePlayer) VLINext() error {
	return zp.VirtualLineIn.Next()
}

// Short for `zp.VirtualLineIn.Pause`.
func (zp *ZonePlayer) VLIPause() error {
	return zp.VirtualLineIn.Pause()
}

// Short for `zp.VirtualLineIn.Play`.
func (zp *ZonePlayer) VLIPlay() error {
	return zp.VirtualLineIn.Play()
}

// Short for `zp.VirtualLineIn.Previous`.
func (zp *ZonePlayer) VLIPrevious() error {
	return zp.VirtualLineIn.Previous()
}

// Short for `zp.VirtualLineIn.SetVolume`.
func (zp *ZonePlayer) VLISetVolume(volume int) error {
	return zp.VirtualLineIn.SetVolume(volume)
}

// Short for `zp.VirtualLineIn.Stop`.
func (zp *ZonePlayer) VLIStop() error {
	return zp.VirtualLineIn.Stop()
}
