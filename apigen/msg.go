package main

import (
	"sort"

	"github.com/paralin/go-underlords/protocol"
)

// IsValidMsg checks if the message is valid.
func IsValidMsg(msg protocol.EGCDACClientMessages) bool {
	_, ok := protocol.EGCDACClientMessages_name[int32(msg)]
	return ok && msg > 9000
}

func getSortedMsgIDs() []protocol.EGCDACClientMessages {
	var msgIds []protocol.EGCDACClientMessages
	for msgIDNum := range protocol.EGCDACClientMessages_name {
		msgID := protocol.EGCDACClientMessages(msgIDNum)
		msgIds = append(msgIds, msgID)
	}

	sort.Slice(msgIds, func(i int, j int) bool {
		return msgIds[i] < msgIds[j]
	})
	return msgIds
}
