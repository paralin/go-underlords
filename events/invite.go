package events

import (
	"github.com/golang/protobuf/proto"
	bgcm "github.com/paralin/go-underlords/protocol"
)

// InvitationCreated confirms that an invitation has been created.
type InvitationCreated struct {
	bgcm.CMsgInvitationCreated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *InvitationCreated) GetDotaEventMsgID() bgcm.EGCDACClientMessages {
	return bgcm.EGCDACClientMessages(bgcm.EGCBaseMsg_k_EMsgGCInvitationCreated)
}

// GetEventBody returns the event body.
func (e *InvitationCreated) GetEventBody() proto.Message {
	return &e.CMsgInvitationCreated
}

// GetEventName returns the event name.
func (e *InvitationCreated) GetEventName() string {
	return "InvitationCreated"
}
