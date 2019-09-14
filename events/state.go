package events

import (
	"github.com/paralin/go-underlords/state"
)

// ClientStateChanged is emitted whenever anything about the client state changes.
type ClientStateChanged struct {
	// OldState is the old state.
	OldState state.DACState
	// NewState is the new state.
	NewState state.DACState
}
