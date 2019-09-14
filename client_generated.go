package dac

import (
	"context"
	"github.com/paralin/go-underlords/events"
	"github.com/paralin/go-underlords/protocol"
)

// CheckFriendCode checks a friend code.
// Request ID: k_EMsgClientToGCCheckFriendCode
// Response ID: k_EMsgClientToGCCheckFriendCodeResponse
// Request type: CMsgClientToGCCheckFriendCode
// Response type: CMsgClientToGCCheckFriendCodeResponse
func (d *DAC) CheckFriendCode(
	ctx context.Context,
	friendCode uint64,
) (*protocol.CMsgClientToGCCheckFriendCodeResponse, error) {
	req := &protocol.CMsgClientToGCCheckFriendCode{
		FriendCode: &friendCode,
	}
	resp := &protocol.CMsgClientToGCCheckFriendCodeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCCheckFriendCode),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCCheckFriendCodeResponse),
		resp,
	)
}

// CreateFriendCode creates a friend code.
// Request ID: k_EMsgClientToGCCreateFriendCode
// Response ID: k_EMsgClientToGCCreateFriendCodeResponse
// Request type: CMsgClientToGCCreateFriendCode
// Response type: CMsgClientToGCCreateFriendCodeResponse
func (d *DAC) CreateFriendCode(
	ctx context.Context,
) (*protocol.CMsgClientToGCCreateFriendCodeResponse, error) {
	req := &protocol.CMsgClientToGCCreateFriendCode{}
	resp := &protocol.CMsgClientToGCCreateFriendCodeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCCreateFriendCode),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCCreateFriendCodeResponse),
		resp,
	)
}

// CreateParty creates a party.
// Request ID: k_EMsgClientToGCPartyCreate
// Response ID: k_EMsgClientToGCPartyCreateResponse
// Request type: CMsgClientToGCPartyCreate
// Response type: CMsgClientToGCPartyCreateResponse
func (d *DAC) CreateParty(
	ctx context.Context,
	partyMmInfo protocol.CMsgPartyMMInfo,
	startReady bool,
	inviteAccountID uint32,
	disablePartyCode bool,
) (*protocol.CMsgClientToGCPartyCreateResponse, error) {
	req := &protocol.CMsgClientToGCPartyCreate{
		PartyMmInfo:      &partyMmInfo,
		StartReady:       &startReady,
		InviteAccountId:  &inviteAccountID,
		DisablePartyCode: &disablePartyCode,
	}
	resp := &protocol.CMsgClientToGCPartyCreateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyCreate),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyCreateResponse),
		resp,
	)
}

// GetFriendCodes gets friend codes.
// Request ID: k_EMsgClientToGCGetFriendCodes
// Response ID: k_EMsgClientToGCGetFriendCodesResponse
// Request type: CMsgClientToGCGetFriendCodes
// Response type: CMsgClientToGCGetFriendCodesResponse
func (d *DAC) GetFriendCodes(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetFriendCodesResponse, error) {
	req := &protocol.CMsgClientToGCGetFriendCodes{}
	resp := &protocol.CMsgClientToGCGetFriendCodesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetFriendCodes),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetFriendCodesResponse),
		resp,
	)
}

// JoinParty joins a party.
// Request ID: k_EMsgClientToGCPartyJoin
// Response ID: k_EMsgClientToGCPartyJoinResponse
// Request type: CMsgClientToGCPartyJoin
// Response type: CMsgClientToGCPartyJoinResponse
func (d *DAC) JoinParty(
	ctx context.Context,
	partyID uint64,
	isRejoin bool,
	partyMmInfo protocol.CMsgPartyMMInfo,
	startReady bool,
) (*protocol.CMsgClientToGCPartyJoinResponse, error) {
	req := &protocol.CMsgClientToGCPartyJoin{
		PartyId:     &partyID,
		IsRejoin:    &isRejoin,
		PartyMmInfo: &partyMmInfo,
		StartReady:  &startReady,
	}
	resp := &protocol.CMsgClientToGCPartyJoinResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyJoin),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyJoinResponse),
		resp,
	)
}

// JoinPartyViaCode joins a party via code.
// Request ID: k_EMsgClientToGCPartyJoinViaCode
// Response ID: k_EMsgClientToGCPartyJoinViaCodeResponse
// Request type: CMsgClientToGCPartyJoinViaCode
// Response type: CMsgClientToGCPartyJoinViaCodeResponse
func (d *DAC) JoinPartyViaCode(
	ctx context.Context,
	joinCode uint64,
	partyMmInfo protocol.CMsgPartyMMInfo,
	startReady bool,
) (*protocol.CMsgClientToGCPartyJoinViaCodeResponse, error) {
	req := &protocol.CMsgClientToGCPartyJoinViaCode{
		JoinCode:    &joinCode,
		PartyMmInfo: &partyMmInfo,
		StartReady:  &startReady,
	}
	resp := &protocol.CMsgClientToGCPartyJoinViaCodeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyJoinViaCode),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyJoinViaCodeResponse),
		resp,
	)
}

// LeaveLobby leaves a lobby.
// Request ID: k_EMsgClientToGCLeaveLobby
// Response ID: k_EMsgClientToGCLeaveLobbyResponse
// Request type: CMsgClientToGCLeaveLobby
// Response type: CMsgClientToGCLeaveLobbyResponse
func (d *DAC) LeaveLobby(
	ctx context.Context,
	lobbyID uint64,
) (*protocol.CMsgClientToGCLeaveLobbyResponse, error) {
	req := &protocol.CMsgClientToGCLeaveLobby{
		LobbyId: &lobbyID,
	}
	resp := &protocol.CMsgClientToGCLeaveLobbyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCLeaveLobby),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCLeaveLobbyResponse),
		resp,
	)
}

// LeaveParty leaves a party.
// Request ID: k_EMsgClientToGCPartyLeave
// Response ID: k_EMsgClientToGCPartyLeaveResponse
// Request type: CMsgClientToGCPartyLeave
// Response type: CMsgClientToGCPartyLeaveResponse
func (d *DAC) LeaveParty(
	ctx context.Context,
	partyID uint64,
) (*protocol.CMsgClientToGCPartyLeaveResponse, error) {
	req := &protocol.CMsgClientToGCPartyLeave{
		PartyId: &partyID,
	}
	resp := &protocol.CMsgClientToGCPartyLeaveResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyLeave),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyLeaveResponse),
		resp,
	)
}

// RecordLocalBotMatch records a local bot match.
// Request ID: k_EMsgClientToGCRecordLocalBotMatch
// Request type: CMsgClientToGCRecordLocalBotMatch
func (d *DAC) RecordLocalBotMatch(
	req *protocol.CMsgClientToGCRecordLocalBotMatch,
) {
	d.write(uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCRecordLocalBotMatch), req)
}

// RedeemFriendCode redeems a friend code.
// Request ID: k_EMsgClientToGCRedeemFriendCode
// Response ID: k_EMsgClientToGCRedeemFriendCodeResponse
// Request type: CMsgClientToGCRedeemFriendCode
// Response type: CMsgClientToGCRedeemFriendCodeResponse
func (d *DAC) RedeemFriendCode(
	ctx context.Context,
	friendCode uint64,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCRedeemFriendCodeResponse, error) {
	req := &protocol.CMsgClientToGCRedeemFriendCode{
		FriendCode:      &friendCode,
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCRedeemFriendCodeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCRedeemFriendCode),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCRedeemFriendCodeResponse),
		resp,
	)
}

// RequestEventInfo requests a event info.
// Request ID: k_EMsgClientToGCEventRequestInfo
// Request type: CMsgClientToGCEventRequestInfo
func (d *DAC) RequestEventInfo(
	accountID uint32,
	eventID uint32,
) {
	req := &protocol.CMsgClientToGCEventRequestInfo{
		AccountId: &accountID,
		EventId:   &eventID,
	}
	d.write(uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventRequestInfo), req)
}

// RequestEventInfoSync requests a event info sync.
// Request ID: k_EMsgClientToGCEventRequestInfoSync
// Response ID: k_EMsgClientToGCEventRequestInfoSyncResponse
// Request type: CMsgClientToGCEventRequestInfoSync
// Response type: CMsgClientToGCEventRequestInfoSyncResponse
func (d *DAC) RequestEventInfoSync(
	ctx context.Context,
	accountID uint32,
	eventID uint32,
) (*protocol.CMsgClientToGCEventRequestInfoSyncResponse, error) {
	req := &protocol.CMsgClientToGCEventRequestInfoSync{
		AccountId: &accountID,
		EventId:   &eventID,
	}
	resp := &protocol.CMsgClientToGCEventRequestInfoSyncResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventRequestInfoSync),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventRequestInfoSyncResponse),
		resp,
	)
}

// RerollChallenge rerolls a challenge.
// Request ID: k_EMsgClientToGCRerollChallenge
// Response ID: k_EMsgClientToGCRerollChallengeResponse
// Request type: CMsgClientToGCRerollChallenge
// Response type: CMsgClientToGCRerollChallengeResponse
func (d *DAC) RerollChallenge(
	ctx context.Context,
	eventID uint32,
	slotID uint32,
	sequenceID uint32,
) (*protocol.CMsgClientToGCRerollChallengeResponse, error) {
	req := &protocol.CMsgClientToGCRerollChallenge{
		EventId:    &eventID,
		SlotId:     &slotID,
		SequenceId: &sequenceID,
	}
	resp := &protocol.CMsgClientToGCRerollChallengeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCRerollChallenge),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCRerollChallengeResponse),
		resp,
	)
}

// RevokeFriendCode revokes a friend code.
// Request ID: k_EMsgClientToGCRevokeFriendCode
// Response ID: k_EMsgClientToGCRevokeFriendCodeResponse
// Request type: CMsgClientToGCRevokeFriendCode
// Response type: CMsgClientToGCRevokeFriendCodeResponse
func (d *DAC) RevokeFriendCode(
	ctx context.Context,
	friendCode uint64,
) (*protocol.CMsgClientToGCRevokeFriendCodeResponse, error) {
	req := &protocol.CMsgClientToGCRevokeFriendCode{
		FriendCode: &friendCode,
	}
	resp := &protocol.CMsgClientToGCRevokeFriendCodeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCRevokeFriendCode),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCRevokeFriendCodeResponse),
		resp,
	)
}

// SendClaimChallengeReward sends a claim challenge reward.
// Request ID: k_EMsgClientToGCClaimChallengeReward
// Response ID: k_EMsgClientToGCClaimChallengeRewardResponse
// Request type: CMsgClientToGCClaimChallengeReward
// Response type: CMsgClientToGCClaimChallengeRewardResponse
func (d *DAC) SendClaimChallengeReward(
	ctx context.Context,
	eventID uint32,
	slotID uint32,
	sequenceID uint32,
	startingClaimed uint32,
) (*protocol.CMsgClientToGCClaimChallengeRewardResponse, error) {
	req := &protocol.CMsgClientToGCClaimChallengeReward{
		EventId:         &eventID,
		SlotId:          &slotID,
		SequenceId:      &sequenceID,
		StartingClaimed: &startingClaimed,
	}
	resp := &protocol.CMsgClientToGCClaimChallengeRewardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCClaimChallengeReward),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCClaimChallengeRewardResponse),
		resp,
	)
}

// SendDevEventOperation sends a dev event operation.
// Request ID: k_EMsgClientToGCDevEventOperation
// Response ID: k_EMsgClientToGCDevEventOperationResponse
// Request type: CMsgClientToGCDevEventOperation
// Response type: CMsgClientToGCDevEventOperationResponse
func (d *DAC) SendDevEventOperation(
	ctx context.Context,
	operation protocol.CMsgClientToGCDevEventOperation_EOperation,
	eventID uint32,
	iD uint32,
	amount uint32,
) (*protocol.CMsgClientToGCDevEventOperationResponse, error) {
	req := &protocol.CMsgClientToGCDevEventOperation{
		Operation: &operation,
		EventId:   &eventID,
		Id:        &iD,
		Amount:    &amount,
	}
	resp := &protocol.CMsgClientToGCDevEventOperationResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCDevEventOperation),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCDevEventOperationResponse),
		resp,
	)
}

// SendDevForceMatchFormation sends a dev force match formation.
// Request ID: k_EMsgClientToGCDevForceMatchFormation
// Request type: CMsgClientToGCDevForceMatchFormation
func (d *DAC) SendDevForceMatchFormation() {
	req := &protocol.CMsgClientToGCDevForceMatchFormation{}
	d.write(uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCDevForceMatchFormation), req)
}

// SendEquipLoadout sends a equip loadout.
// Request ID: k_EMsgClientToGCEquipLoadout
// Response ID: k_EMsgClientToGCEquipLoadoutResponse
// Request type: CMsgClientToGCEquipLoadout
// Response type: CMsgClientToGCEquipLoadoutResponse
func (d *DAC) SendEquipLoadout(
	ctx context.Context,
	loadoutSlot []uint32,
	itemID []uint64,
	loadoutSubSlot []uint32,
) (*protocol.CMsgClientToGCEquipLoadoutResponse, error) {
	req := &protocol.CMsgClientToGCEquipLoadout{
		LoadoutSlot:    loadoutSlot,
		ItemId:         itemID,
		LoadoutSubSlot: loadoutSubSlot,
	}
	resp := &protocol.CMsgClientToGCEquipLoadoutResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEquipLoadout),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEquipLoadoutResponse),
		resp,
	)
}

// SendEventClaim sends a event claim.
// Request ID: k_EMsgClientToGCEventClaim
// Response ID: k_EMsgClientToGCEventClaimResponse
// Request type: CMsgClientToGCEventClaim
// Response type: CMsgClientToGCEventClaimResponse
func (d *DAC) SendEventClaim(
	ctx context.Context,
	eventID uint32,
	claimID uint32,
	startingValue uint32,
) (*protocol.CMsgClientToGCEventClaimResponse, error) {
	req := &protocol.CMsgClientToGCEventClaim{
		EventId:       &eventID,
		ClaimId:       &claimID,
		StartingValue: &startingValue,
	}
	resp := &protocol.CMsgClientToGCEventClaimResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventClaim),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventClaimResponse),
		resp,
	)
}

// SendEventEquipVirtualItem sends a event equip virtual item.
// Request ID: k_EMsgClientToGCEventEquipVirtualItem
// Response ID: k_EMsgClientToGCEventEquipVirtualItemResponse
// Request type: CMsgClientToGCEventEquipVirtualItem
// Response type: CMsgClientToGCEventEquipVirtualItemResponse
func (d *DAC) SendEventEquipVirtualItem(
	ctx context.Context,
	eventID uint32,
	defIndex uint32,
	equipSlot uint32,
	equipSubSlot uint32,
) (*protocol.CMsgClientToGCEventEquipVirtualItemResponse, error) {
	req := &protocol.CMsgClientToGCEventEquipVirtualItem{
		EventId:      &eventID,
		DefIndex:     &defIndex,
		EquipSlot:    &equipSlot,
		EquipSubSlot: &equipSubSlot,
	}
	resp := &protocol.CMsgClientToGCEventEquipVirtualItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventEquipVirtualItem),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventEquipVirtualItemResponse),
		resp,
	)
}

// SendIsInMatchmaking sends a is in matchmaking.
// Request ID: k_EMsgClientToGCIsInMatchmaking
// Response ID: k_EMsgClientToGCIsInMatchmakingResponse
// Request type: CMsgClientToGCIsInMatchmaking
// Response type: CMsgClientToGCIsInMatchmakingResponse
func (d *DAC) SendIsInMatchmaking(
	ctx context.Context,
) (*protocol.CMsgClientToGCIsInMatchmakingResponse, error) {
	req := &protocol.CMsgClientToGCIsInMatchmaking{}
	resp := &protocol.CMsgClientToGCIsInMatchmakingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCIsInMatchmaking),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCIsInMatchmakingResponse),
		resp,
	)
}

// SendPartyAction sends a party action.
// Request ID: k_EMsgClientToGCPartyAction
// Response ID: k_EMsgClientToGCPartyActionResponse
// Request type: CMsgClientToGCPartyAction
// Response type: CMsgClientToGCPartyActionResponse
func (d *DAC) SendPartyAction(
	ctx context.Context,
	partyID uint64,
	targetAccountID uint32,
	actionID protocol.CMsgClientToGCPartyAction_EAction,
	uintValue uint64,
	boolValue bool,
) (*protocol.CMsgClientToGCPartyActionResponse, error) {
	req := &protocol.CMsgClientToGCPartyAction{
		PartyId:         &partyID,
		TargetAccountId: &targetAccountID,
		ActionId:        &actionID,
		UintValue:       &uintValue,
		BoolValue:       &boolValue,
	}
	resp := &protocol.CMsgClientToGCPartyActionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyAction),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyActionResponse),
		resp,
	)
}

// SendPartyClientVersion sends a party client version.
// Request ID: k_EMsgClientToGCPartyClientVersion
// Request type: CMsgClientToGCPartyClientVersion
func (d *DAC) SendPartyClientVersion(
	partyID uint64,
	partyMmInfo protocol.CMsgPartyMMInfo,
) {
	req := &protocol.CMsgClientToGCPartyClientVersion{
		PartyId:     &partyID,
		PartyMmInfo: &partyMmInfo,
	}
	d.write(uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyClientVersion), req)
}

// SendPartyInviteUser sends a party invite user.
// Request ID: k_EMsgClientToGCPartyInviteUser
// Response ID: k_EMsgClientToGCPartyInviteUserResponse
// Request type: CMsgClientToGCPartyInviteUser
// Response type: CMsgClientToGCPartyInviteUserResponse
func (d *DAC) SendPartyInviteUser(
	ctx context.Context,
	partyID uint64,
	inviteAccountID uint32,
) (*protocol.CMsgClientToGCPartyInviteUserResponse, error) {
	req := &protocol.CMsgClientToGCPartyInviteUser{
		PartyId:         &partyID,
		InviteAccountId: &inviteAccountID,
	}
	resp := &protocol.CMsgClientToGCPartyInviteUserResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyInviteUser),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyInviteUserResponse),
		resp,
	)
}

// SendPerformAutoActions sends perform auto actions.
// Request ID: k_EMsgClientToGCPerformAutoActions
// Response ID: k_EMsgClientToGCPerformAutoActionsResponse
// Request type: CMsgClientToGCPerformAutoActions
// Response type: CMsgClientToGCPerformAutoActionsResponse
func (d *DAC) SendPerformAutoActions(
	ctx context.Context,
	eventID uint32,
	updateAutoClaims bool,
	updateChallengeSlots []uint32,
	updateClaims []uint32,
) (*protocol.CMsgClientToGCPerformAutoActionsResponse, error) {
	req := &protocol.CMsgClientToGCPerformAutoActions{
		EventId:              &eventID,
		UpdateAutoClaims:     &updateAutoClaims,
		UpdateChallengeSlots: updateChallengeSlots,
		UpdateClaims:         updateClaims,
	}
	resp := &protocol.CMsgClientToGCPerformAutoActionsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPerformAutoActions),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPerformAutoActionsResponse),
		resp,
	)
}

// SendReplacementSDRTicket sends a replacement sdr ticket.
// Request ID: k_EMsgClientToGCReplacementSDRTicket
// Response ID: k_EMsgClientToGCReplacementSDRTicketResponse
// Request type: CMsgClientToGCReplacementSDRTicket
// Response type: CMsgClientToGCReplacementSDRTicketResponse
func (d *DAC) SendReplacementSDRTicket(
	ctx context.Context,
	lobbyID uint64,
) (*protocol.CMsgClientToGCReplacementSDRTicketResponse, error) {
	req := &protocol.CMsgClientToGCReplacementSDRTicket{
		LobbyId: &lobbyID,
	}
	resp := &protocol.CMsgClientToGCReplacementSDRTicketResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCReplacementSDRTicket),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCReplacementSDRTicketResponse),
		resp,
	)
}

// SetActiveUnderlord sets a active underlord.
// Request ID: k_EMsgClientToGCSetActiveUnderlord
// Response ID: k_EMsgClientToGCSetActiveUnderlordResponse
// Request type: CMsgClientToGCSetActiveUnderlord
// Response type: CMsgClientToGCSetActiveUnderlordResponse
func (d *DAC) SetActiveUnderlord(
	ctx context.Context,
	underlordID uint32,
) (*protocol.CMsgClientToGCSetActiveUnderlordResponse, error) {
	req := &protocol.CMsgClientToGCSetActiveUnderlord{
		UnderlordId: &underlordID,
	}
	resp := &protocol.CMsgClientToGCSetActiveUnderlordResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCSetActiveUnderlord),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCSetActiveUnderlordResponse),
		resp,
	)
}

// SpectateUser spectates a user.
// Request ID: k_EMsgClientToGCSpectateUser
// Response ID: k_EMsgClientToGCSpectateUserResponse
// Request type: CMsgClientToGCSpectateUser
// Response type: CMsgClientToGCSpectateUserResponse
func (d *DAC) SpectateUser(
	ctx context.Context,
	spectateAccountID uint32,
) (*protocol.CMsgClientToGCSpectateUserResponse, error) {
	req := &protocol.CMsgClientToGCSpectateUser{
		SpectateAccountId: &spectateAccountID,
	}
	resp := &protocol.CMsgClientToGCSpectateUserResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCSpectateUser),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCSpectateUserResponse),
		resp,
	)
}

// StartMatchmaking starts a matchmaking.
// Request ID: k_EMsgClientToGCStartMatchmaking
// Response ID: k_EMsgClientToGCStartMatchmakingResponse
// Request type: CMsgClientToGCStartMatchmaking
// Response type: CMsgClientToGCStartMatchmakingResponse
func (d *DAC) StartMatchmaking(
	ctx context.Context,
	matchInfo protocol.CMsgStartFindingMatchInfo,
	pingTimes protocol.CMsgRegionPingTimesClient,
) (*protocol.CMsgClientToGCStartMatchmakingResponse, error) {
	req := &protocol.CMsgClientToGCStartMatchmaking{
		MatchInfo: &matchInfo,
		PingTimes: &pingTimes,
	}
	resp := &protocol.CMsgClientToGCStartMatchmakingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCStartMatchmaking),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCStartMatchmakingResponse),
		resp,
	)
}

// StartPartyMatch starts a party match.
// Request ID: k_EMsgClientToGCPartyStartMatch
// Response ID: k_EMsgClientToGCPartyStartMatchResponse
// Request type: CMsgClientToGCPartyStartMatch
// Response type: CMsgClientToGCPartyStartMatchResponse
func (d *DAC) StartPartyMatch(
	ctx context.Context,
	partyID uint64,
	privateMatch bool,
	serverSearchKey string,
	matchMode protocol.EDACMatchMode,
) (*protocol.CMsgClientToGCPartyStartMatchResponse, error) {
	req := &protocol.CMsgClientToGCPartyStartMatch{
		PartyId:         &partyID,
		PrivateMatch:    &privateMatch,
		ServerSearchKey: &serverSearchKey,
		MatchMode:       &matchMode,
	}
	resp := &protocol.CMsgClientToGCPartyStartMatchResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyStartMatch),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartyStartMatchResponse),
		resp,
	)
}

// StopMatchmaking stops a matchmaking.
// Request ID: k_EMsgClientToGCStopMatchmaking
// Response ID: k_EMsgClientToGCStopMatchmakingResponse
// Request type: CMsgClientToGCStopMatchmaking
// Response type: CMsgClientToGCStopMatchmakingResponse
func (d *DAC) StopMatchmaking(
	ctx context.Context,
) (*protocol.CMsgClientToGCStopMatchmakingResponse, error) {
	req := &protocol.CMsgClientToGCStopMatchmaking{}
	resp := &protocol.CMsgClientToGCStopMatchmakingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCStopMatchmaking),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCStopMatchmakingResponse),
		resp,
	)
}

// registerGeneratedHandlers registers the auto-generated event handlers.
func (d *DAC) registerGeneratedHandlers() {
	d.handlers[uint32(protocol.EGCDACClientMessages_k_EMsgGCToClientCanRejoinParty)] = d.getEventEmitter(func() events.Event {
		return &events.CanRejoinParty{}
	})
	d.handlers[uint32(protocol.EGCDACClientMessages_k_EMsgGCToClientDevMMStatus)] = d.getEventEmitter(func() events.Event {
		return &events.DevMMStatus{}
	})
	d.handlers[uint32(protocol.EGCDACClientMessages_k_EMsgGCToClientEventInfo)] = d.getEventEmitter(func() events.Event {
		return &events.EventInfo{}
	})
	d.handlers[uint32(protocol.EGCDACClientMessages_k_EMsgGCToClientMatchmakingStopped)] = d.getEventEmitter(func() events.Event {
		return &events.MatchmakingStopped{}
	})
	d.handlers[uint32(protocol.EGCDACClientMessages_k_EMsgGCToClientPartyEvent)] = d.getEventEmitter(func() events.Event {
		return &events.PartyEvent{}
	})
	d.handlers[uint32(protocol.EGCDACClientMessages_k_EMsgGCToClientSDRTicket)] = d.getEventEmitter(func() events.Event {
		return &events.SDRTicket{}
	})
	d.handlers[uint32(protocol.EGCDACClientMessages_k_EMsgGCToClientUpdateConsoleCommands)] = d.getEventEmitter(func() events.Event {
		return &events.UpdateConsoleCommands{}
	})
}
