package Helper

// Short for `zp.Queue.RemoveTrackRange`.
func (h *Helper) RemoveTrackRange(start int, count int) error {
	_, err := h.queue.RemoveTrackRange(start, count)
	return err
}

// Short for `zp.Queue.ReorderTracks`.
func (h *Helper) ReorderTracks(start int, count int, insertBefore int) error {
	_, err := h.queue.ReorderTracks(start, count, insertBefore)
	return err
}
