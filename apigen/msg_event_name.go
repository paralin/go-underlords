package main

import (
	"strings"

	"github.com/paralin/go-underlords/protocol"
)

// GetMessageEventName returns the event name for the message.
func GetMessageEventName(msg protocol.EGCDACClientMessages) string {
	if over, ok := msgEventNameOverrides[msg]; ok {
		return over
	}

	msgName := msg.String()
	msgName = strings.TrimPrefix(msgName, "k_EMsg")
	msgName = strings.TrimPrefix(msgName, "GC")
	msgName = strings.TrimPrefix(msgName, "ToClient")
	msgName = strings.Replace(msgName, "_", "", -1)
	return msgName
}
