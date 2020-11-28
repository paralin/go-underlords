package events

import (
	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-underlords/protocol"
)

// AcquireRegionModeInfo event.
// MessageID: k_EMsgGCToClientAcquireRegionModeInfo
type AcquireRegionModeInfo struct {
	protocol.CMsgGCToClientAcquireRegionModeInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *AcquireRegionModeInfo) GetDotaEventMsgID() protocol.EGCDACClientMessages {
	return protocol.EGCDACClientMessages_k_EMsgGCToClientAcquireRegionModeInfo
}

// GetEventBody returns the event body.
func (e *AcquireRegionModeInfo) GetEventBody() proto.Message {
	return &e.CMsgGCToClientAcquireRegionModeInfo
}

// GetEventName returns the event name.
func (e *AcquireRegionModeInfo) GetEventName() string {
	return "AcquireRegionModeInfo"
}

// CanRejoinParty event.
// MessageID: k_EMsgGCToClientCanRejoinParty
type CanRejoinParty struct {
	protocol.CMsgGCToClientCanRejoinParty
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *CanRejoinParty) GetDotaEventMsgID() protocol.EGCDACClientMessages {
	return protocol.EGCDACClientMessages_k_EMsgGCToClientCanRejoinParty
}

// GetEventBody returns the event body.
func (e *CanRejoinParty) GetEventBody() proto.Message {
	return &e.CMsgGCToClientCanRejoinParty
}

// GetEventName returns the event name.
func (e *CanRejoinParty) GetEventName() string {
	return "CanRejoinParty"
}

// DevMMStatus event.
// MessageID: k_EMsgGCToClientDevMMStatus
type DevMMStatus struct {
	protocol.CMsgGCToClientDevMMStatus
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *DevMMStatus) GetDotaEventMsgID() protocol.EGCDACClientMessages {
	return protocol.EGCDACClientMessages_k_EMsgGCToClientDevMMStatus
}

// GetEventBody returns the event body.
func (e *DevMMStatus) GetEventBody() proto.Message {
	return &e.CMsgGCToClientDevMMStatus
}

// GetEventName returns the event name.
func (e *DevMMStatus) GetEventName() string {
	return "DevMMStatus"
}

// DurationControlWarning event.
// MessageID: k_EMsgGCToClientDurationControlWarning
type DurationControlWarning struct {
	protocol.CMsgGCToClientDurationControlWarning
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *DurationControlWarning) GetDotaEventMsgID() protocol.EGCDACClientMessages {
	return protocol.EGCDACClientMessages_k_EMsgGCToClientDurationControlWarning
}

// GetEventBody returns the event body.
func (e *DurationControlWarning) GetEventBody() proto.Message {
	return &e.CMsgGCToClientDurationControlWarning
}

// GetEventName returns the event name.
func (e *DurationControlWarning) GetEventName() string {
	return "DurationControlWarning"
}

// EventInfo event.
// MessageID: k_EMsgGCToClientEventInfo
type EventInfo struct {
	protocol.CMsgGCToClientEventInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *EventInfo) GetDotaEventMsgID() protocol.EGCDACClientMessages {
	return protocol.EGCDACClientMessages_k_EMsgGCToClientEventInfo
}

// GetEventBody returns the event body.
func (e *EventInfo) GetEventBody() proto.Message {
	return &e.CMsgGCToClientEventInfo
}

// GetEventName returns the event name.
func (e *EventInfo) GetEventName() string {
	return "EventInfo"
}

// GameModifiersUpdated event.
// MessageID: k_EMsgGCToClientGameModifiersUpdated
type GameModifiersUpdated struct {
	protocol.CMsgGCToClientGameModifiersUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *GameModifiersUpdated) GetDotaEventMsgID() protocol.EGCDACClientMessages {
	return protocol.EGCDACClientMessages_k_EMsgGCToClientGameModifiersUpdated
}

// GetEventBody returns the event body.
func (e *GameModifiersUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientGameModifiersUpdated
}

// GetEventName returns the event name.
func (e *GameModifiersUpdated) GetEventName() string {
	return "GameModifiersUpdated"
}

// MatchmakingStopped event.
// MessageID: k_EMsgGCToClientMatchmakingStopped
type MatchmakingStopped struct {
	protocol.CMsgGCToClientMatchmakingStopped
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *MatchmakingStopped) GetDotaEventMsgID() protocol.EGCDACClientMessages {
	return protocol.EGCDACClientMessages_k_EMsgGCToClientMatchmakingStopped
}

// GetEventBody returns the event body.
func (e *MatchmakingStopped) GetEventBody() proto.Message {
	return &e.CMsgGCToClientMatchmakingStopped
}

// GetEventName returns the event name.
func (e *MatchmakingStopped) GetEventName() string {
	return "MatchmakingStopped"
}

// PartyEvent event.
// MessageID: k_EMsgGCToClientPartyEvent
type PartyEvent struct {
	protocol.CMsgGCToClientPartyEvent
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PartyEvent) GetDotaEventMsgID() protocol.EGCDACClientMessages {
	return protocol.EGCDACClientMessages_k_EMsgGCToClientPartyEvent
}

// GetEventBody returns the event body.
func (e *PartyEvent) GetEventBody() proto.Message {
	return &e.CMsgGCToClientPartyEvent
}

// GetEventName returns the event name.
func (e *PartyEvent) GetEventName() string {
	return "PartyEvent"
}

// SDRTicket event.
// MessageID: k_EMsgGCToClientSDRTicket
type SDRTicket struct {
	protocol.CMsgGCToClientSDRTicket
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *SDRTicket) GetDotaEventMsgID() protocol.EGCDACClientMessages {
	return protocol.EGCDACClientMessages_k_EMsgGCToClientSDRTicket
}

// GetEventBody returns the event body.
func (e *SDRTicket) GetEventBody() proto.Message {
	return &e.CMsgGCToClientSDRTicket
}

// GetEventName returns the event name.
func (e *SDRTicket) GetEventName() string {
	return "SDRTicket"
}

// UpdateConsoleCommands event.
// MessageID: k_EMsgGCToClientUpdateConsoleCommands
type UpdateConsoleCommands struct {
	protocol.CMsgGCToClientUpdateConsoleCommands
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *UpdateConsoleCommands) GetDotaEventMsgID() protocol.EGCDACClientMessages {
	return protocol.EGCDACClientMessages_k_EMsgGCToClientUpdateConsoleCommands
}

// GetEventBody returns the event body.
func (e *UpdateConsoleCommands) GetEventBody() proto.Message {
	return &e.CMsgGCToClientUpdateConsoleCommands
}

// GetEventName returns the event name.
func (e *UpdateConsoleCommands) GetEventName() string {
	return "UpdateConsoleCommands"
}
