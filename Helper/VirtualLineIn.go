package Helper

// Short for `zp.VirtualLineIn.Next`.
func (h *Helper) VLINext() error { return h.virtualLineIn.Next() }

// Short for `zp.VirtualLineIn.Pause`.
func (h *Helper) VLIPause() error { return h.virtualLineIn.Pause() }

// Short for `zp.VirtualLineIn.Play`.
func (h *Helper) VLIPlay() error { return h.virtualLineIn.Play() }

// Short for `zp.VirtualLineIn.Previous`.
func (h *Helper) VLIPrevious() error { return h.virtualLineIn.Previous() }

// Short for `zp.VirtualLineIn.SetVolume`.
func (h *Helper) VLISetVolume(volume int) error { return h.virtualLineIn.SetVolume(volume) }

// Short for `zp.VirtualLineIn.Stop`.
func (h *Helper) VLIStop() error { return h.virtualLineIn.Stop() }
