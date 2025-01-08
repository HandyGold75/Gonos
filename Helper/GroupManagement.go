package Helper

// Short for `zp.ConnectionManager.GetCurrentConnectionIDs`.
func (h *Helper) AddMember(memberID string, bootSeq int) (string, error) {
	out, err := h.groupManagement.AddMember(memberID, bootSeq)
	return out.CurrentURI, err
}

func (h *Helper) RemoveMember(memberID string) error {
	return h.groupManagement.RemoveMember(memberID)
}
