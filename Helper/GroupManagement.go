package Helper

// Short for `zp.GroupManagement.AddMember'`.
func (h *Helper) AddMember(memberID string, bootSeq int) (string, error) {
	out, err := h.groupManagement.AddMember(memberID, bootSeq)
	return out.CurrentURI, err
}

// Short for `zp.GroupManagement.RemoveMember'`.
func (h *Helper) RemoveMember(memberID string) error {
	return h.groupManagement.RemoveMember(memberID)
}
