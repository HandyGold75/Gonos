package Gonos

// Short for `zp.Queue.RemoveTrackRange`.
func (zp *ZonePlayer) RemoveTrackRange(start int, count int) error {
	_, err := zp.Queue.RemoveTrackRange(start, count)
	return err
}

// Short for `zp.Queue.ReorderTracks`.
func (zp *ZonePlayer) ReorderTracks(start int, count int, insertBefore int) error {
	_, err := zp.Queue.ReorderTracks(start, count, insertBefore)
	return err
}
