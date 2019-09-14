package main

import (
	"github.com/golang/protobuf/proto"
	dm "github.com/paralin/go-underlords/protocol"
)

// msgSenderOverrides overrides the heuristic-generated sender parties for each message
// Most of the MsgSenderNone messages are not GC<->Client messages.
var msgSenderOverrides = map[dm.EGCDACClientMessages]MsgSender{
	// dm.EGCDACClientMessages_k_EMsgGCLiveScoreboardUpdate: MsgSenderNone,
	// dm.EGCDACClientMessages_k_EMsgClientsRejoinChatChannels: MsgSenderClient,
}

// msgMethodNameOverrides overrides the generated client method names.
var msgMethodNameOverrides = map[dm.EGCDACClientMessages]string{}

// msgResponseOverrides maps request message IDs to response message IDs.
// Setting zero as the value indicates it is an action and not a query
var msgResponseOverrides = map[dm.EGCDACClientMessages]dm.EGCDACClientMessages{
	// Example:
	// dm.EGCDACClientMessages_k_EMsgClientToGCCreatePlayerCardPack: dm.EGCDACClientMessages_k_EMsgClientToGCCreatePlayerCardPackResponse,
}

// msgProtoTypeOverrides overrides the GC message to proto mapping.
var msgProtoTypeOverrides = map[dm.EGCDACClientMessages]proto.Message{
	// dm.EGCDACClientMessages_k_EMsgGCToClientTeamInfo: &dm.CMsgDOTATeamInfo{},
}

var msgArgAsParameterOverrides = map[dm.EGCDACClientMessages]bool{
	// dm.EGCDACClientMessages_k_EMsgGCPracticeLobbySetDetails: true,
}

var msgEventNameOverrides = map[dm.EGCDACClientMessages]string{
	// dm.EGCDACClientMessages_k_EMsgGCTeamInvite_GCRequestToInvitee:  "TeamInviteReceived",
}
