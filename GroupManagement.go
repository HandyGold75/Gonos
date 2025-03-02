package Gonos

// Short for `zp.GroupManagement.AddMember'`.
func (zp *ZonePlayer) AddMember(memberID string, bootSeq int) (string, error) {
	out, err := zp.GroupManagement.AddMember(memberID, bootSeq)
	return out.CurrentURI, err
}

// Short for `zp.GroupManagement.RemoveMember'`.
func (zp *ZonePlayer) RemoveMember(memberID string) error {
	return zp.GroupManagement.RemoveMember(memberID)
}
