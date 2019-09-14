package state

import (
	gcmm "github.com/paralin/go-underlords/protocol"
	gcsdkm "github.com/paralin/go-underlords/protocol"
)

// DACState is a snapshot of the client state at a point in time.
type DACState struct {
	// ConnectionStatus is the status of the connection to the GC.
	ConnectionStatus gcsdkm.GCConnectionStatus
	// Lobby is the current lobby object.
	Lobby *gcmm.CSODACLobby
	// Party is the current party object.
	Party *gcmm.CSODACLobby
	// PartyInvite is the active incoming party invite.
	// PartyInvite *gcmm.CSODOTAPartyInvite
	// LastConnectionStatusUpdate is the last connection state update we received.
	LastConnectionStatusUpdate *gcsdkm.CMsgConnectionStatus
}

// ClearState clears everything.
func (s *DACState) ClearState() {
	*s = DACState{ConnectionStatus: gcsdkm.GCConnectionStatus_GCConnectionStatus_NO_SESSION}
}

// IsReady checks if the client is ready to receive requests.
func (s *DACState) IsReady() bool {
	return s.ConnectionStatus == gcsdkm.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION
}
