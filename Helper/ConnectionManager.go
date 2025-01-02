package Helper

type (
	ConnectionInfo struct {
		RcsID                 int
		AVTransportID         int
		ProtocolInfo          string
		PeerConnectionManager string
		PeerConnectionID      int
		// Possible values: `Input`, `Output`
		Direction string
		// Possible values: `OK`, `ContentFormatMismatch`, `InsufficientBandwidth`, `UnreliableChannel`, `Unknown`
		Status string
	}
)

// Short for `zp.ConnectionManager.GetCurrentConnectionIDs`.
func (h *Helper) GetCurrentConnectionIDs() (string, error) {
	return h.connectionManager.GetCurrentConnectionIDs()
}

// Short for `zp.ConnectionManager.GetCurrentConnectionInfo`.
func (h *Helper) GetCurrentConnectionInfo(connectionID string) (ConnectionInfo, error) {
	info, err := h.connectionManager.GetCurrentConnectionInfo(connectionID)
	if err != nil {
		return ConnectionInfo{}, err
	}
	return ConnectionInfo{
		RcsID:                 info.RcsID,
		AVTransportID:         info.AVTransportID,
		ProtocolInfo:          info.ProtocolInfo,
		PeerConnectionManager: info.PeerConnectionManager,
		PeerConnectionID:      info.PeerConnectionID,
		Direction:             info.Direction,
		Status:                info.Status,
	}, nil
}

// Short for `zp.ConnectionManager.GetProtocolInfo`.
func (h *Helper) GetProtocolInfo() (string, string, error) {
	res, err := h.connectionManager.GetProtocolInfo()
	if err != nil {
		return "", "", nil
	}
	return res.Source, res.Sink, err
}
