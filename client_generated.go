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

// GetDuosRanks gets duos ranks.
// Request ID: k_EMsgClientToGCGetDuosRanks
// Response ID: k_EMsgClientToGCGetDuosRanksResponse
// Request type: CMsgClientToGCGetDuosRanks
// Response type: CMsgClientToGCGetDuosRanksResponse
func (d *DAC) GetDuosRanks(
	ctx context.Context,
	accountID uint32,
	cursorValue uint32,
	batchSize uint32,
	sortOrder protocol.CMsgClientToGCGetDuosRanks_ESortOrder,
) (*protocol.CMsgClientToGCGetDuosRanksResponse, error) {
	req := &protocol.CMsgClientToGCGetDuosRanks{
		AccountId:   &accountID,
		CursorValue: &cursorValue,
		BatchSize:   &batchSize,
		SortOrder:   &sortOrder,
	}
	resp := &protocol.CMsgClientToGCGetDuosRanksResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetDuosRanks),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetDuosRanksResponse),
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

// GetFriendRanks gets friend ranks.
// Request ID: k_EMsgClientToGCGetFriendRanks
// Response ID: k_EMsgClientToGCGetFriendRanksResponse
// Request type: CMsgClientToGCGetFriendRanks
// Response type: CMsgClientToGCGetFriendRanksResponse
func (d *DAC) GetFriendRanks(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetFriendRanksResponse, error) {
	req := &protocol.CMsgClientToGCGetFriendRanks{}
	resp := &protocol.CMsgClientToGCGetFriendRanksResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetFriendRanks),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetFriendRanksResponse),
		resp,
	)
}

// GetMatchHistory gets a match history.
// Request ID: k_EMsgClientToGCGetMatchHistory
// Response ID: k_EMsgClientToGCGetMatchHistoryResponse
// Request type: CMsgClientToGCGetMatchHistory
// Response type: CMsgClientToGCGetMatchHistoryResponse
func (d *DAC) GetMatchHistory(
	ctx context.Context,
	accountID uint32,
	requestRows uint32,
	matchIDCursor uint64,
) (*protocol.CMsgClientToGCGetMatchHistoryResponse, error) {
	req := &protocol.CMsgClientToGCGetMatchHistory{
		AccountId:     &accountID,
		RequestRows:   &requestRows,
		MatchIdCursor: &matchIDCursor,
	}
	resp := &protocol.CMsgClientToGCGetMatchHistoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetMatchHistory),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetMatchHistoryResponse),
		resp,
	)
}

// GetPostMatchStats gets post match stats.
// Request ID: k_EMsgClientToGCGetPostMatchStats
// Response ID: k_EMsgClientToGCGetPostMatchStatsResponse
// Request type: CMsgClientToGCGetPostMatchStats
// Response type: CMsgClientToGCGetPostMatchStatsResponse
func (d *DAC) GetPostMatchStats(
	ctx context.Context,
	matchID uint64,
) (*protocol.CMsgClientToGCGetPostMatchStatsResponse, error) {
	req := &protocol.CMsgClientToGCGetPostMatchStats{
		MatchId: &matchID,
	}
	resp := &protocol.CMsgClientToGCGetPostMatchStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetPostMatchStats),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetPostMatchStatsResponse),
		resp,
	)
}

// GetProfile gets a profile.
// Request ID: k_EMsgClientToGCGetProfile
// Response ID: k_EMsgClientToGCGetProfileResponse
// Request type: CMsgClientToGCGetProfile
// Response type: CMsgClientToGCGetProfileResponse
func (d *DAC) GetProfile(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCGetProfileResponse, error) {
	req := &protocol.CMsgClientToGCGetProfile{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCGetProfileResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetProfile),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetProfileResponse),
		resp,
	)
}

// GetPuzzleLeaderboards gets puzzle leaderboards.
// Request ID: k_EMsgClientToGCGetPuzzleLeaderboards
// Response ID: k_EMsgClientToGCGetPuzzleLeaderboardsResponse
// Request type: CMsgClientToGCGetPuzzleLeaderboards
// Response type: CMsgClientToGCGetPuzzleLeaderboardsResponse
func (d *DAC) GetPuzzleLeaderboards(
	ctx context.Context,
	puzzleID uint32,
	count uint32,
) (*protocol.CMsgClientToGCGetPuzzleLeaderboardsResponse, error) {
	req := &protocol.CMsgClientToGCGetPuzzleLeaderboards{
		PuzzleId: &puzzleID,
		Count:    &count,
	}
	resp := &protocol.CMsgClientToGCGetPuzzleLeaderboardsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetPuzzleLeaderboards),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetPuzzleLeaderboardsResponse),
		resp,
	)
}

// GetRegionModeInfo gets a region mode info.
// Request ID: k_EMsgClientToGCGetRegionModeInfo
// Response ID: k_EMsgClientToGCGetRegionModeInfoResponse
// Request type: CMsgClientToGCGetRegionModeInfo
// Response type: CMsgClientToGCGetRegionModeInfoResponse
func (d *DAC) GetRegionModeInfo(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetRegionModeInfoResponse, error) {
	req := &protocol.CMsgClientToGCGetRegionModeInfo{}
	resp := &protocol.CMsgClientToGCGetRegionModeInfoResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetRegionModeInfo),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCGetRegionModeInfoResponse),
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

// PurchaseCanItem purchases a can item.
// Request ID: k_EMsgClientToGCCanPurchaseItem
// Response ID: k_EMsgClientToGCCanPurchaseItemResponse
// Request type: CMsgClientToGCCanPurchaseItem
// Response type: CMsgClientToGCCanPurchaseItemResponse
func (d *DAC) PurchaseCanItem(
	ctx context.Context,
	req *protocol.CMsgClientToGCCanPurchaseItem,
) (*protocol.CMsgClientToGCCanPurchaseItemResponse, error) {
	resp := &protocol.CMsgClientToGCCanPurchaseItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCCanPurchaseItem),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCCanPurchaseItemResponse),
		resp,
	)
}

// PurchaseClearReserve purchases a clear reserve.
// Request ID: k_EMsgClientToGCClearPurchaseReserve
// Response ID: k_EMsgClientToGCClearPurchaseReserveResponse
// Request type: CMsgClientToGCClearPurchaseReserve
// Response type: CMsgClientToGCClearPurchaseReserveResponse
func (d *DAC) PurchaseClearReserve(
	ctx context.Context,
	defIndex uint32,
	storeID protocol.EDACStoreID,
	deviceID uint64,
) (*protocol.CMsgClientToGCClearPurchaseReserveResponse, error) {
	req := &protocol.CMsgClientToGCClearPurchaseReserve{
		DefIndex: &defIndex,
		StoreId:  &storeID,
		DeviceId: &deviceID,
	}
	resp := &protocol.CMsgClientToGCClearPurchaseReserveResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCClearPurchaseReserve),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCClearPurchaseReserveResponse),
		resp,
	)
}

// PurchaseEvent purchases a event.
// Request ID: k_EMsgClientToGCEventPurchase
// Response ID: k_EMsgClientToGCEventPurchaseResponse
// Request type: CMsgClientToGCEventPurchase
// Response type: CMsgClientToGCEventPurchaseResponse
func (d *DAC) PurchaseEvent(
	ctx context.Context,
	eventID uint32,
	expectedCredits uint32,
) (*protocol.CMsgClientToGCEventPurchaseResponse, error) {
	req := &protocol.CMsgClientToGCEventPurchase{
		EventId:         &eventID,
		ExpectedCredits: &expectedCredits,
	}
	resp := &protocol.CMsgClientToGCEventPurchaseResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventPurchase),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventPurchaseResponse),
		resp,
	)
}

// PurchaseEventCurrency purchases a event currency.
// Request ID: k_EMsgClientToGCEventPurchaseCurrency
// Response ID: k_EMsgClientToGCEventPurchaseCurrencyResponse
// Request type: CMsgClientToGCEventPurchaseCurrency
// Response type: CMsgClientToGCEventPurchaseCurrencyResponse
func (d *DAC) PurchaseEventCurrency(
	ctx context.Context,
	eventID uint32,
	purchaseQuantity uint32,
	expectedCost uint32,
	currencyID uint32,
	startingBalance uint32,
) (*protocol.CMsgClientToGCEventPurchaseCurrencyResponse, error) {
	req := &protocol.CMsgClientToGCEventPurchaseCurrency{
		EventId:          &eventID,
		PurchaseQuantity: &purchaseQuantity,
		ExpectedCost:     &expectedCost,
		CurrencyId:       &currencyID,
		StartingBalance:  &startingBalance,
	}
	resp := &protocol.CMsgClientToGCEventPurchaseCurrencyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventPurchaseCurrency),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventPurchaseCurrencyResponse),
		resp,
	)
}

// PurchaseItemHandle purchases a item handle.
// Request ID: k_EMsgClientToGCHandleItemPurchase
// Response ID: k_EMsgClientToGCHandleItemPurchaseResponse
// Request type: CMsgClientToGCHandleItemPurchase
// Response type: CMsgClientToGCHandleItemPurchaseResponse
func (d *DAC) PurchaseItemHandle(
	ctx context.Context,
	req *protocol.CMsgClientToGCHandleItemPurchase,
) (*protocol.CMsgClientToGCHandleItemPurchaseResponse, error) {
	resp := &protocol.CMsgClientToGCHandleItemPurchaseResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCHandleItemPurchase),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCHandleItemPurchaseResponse),
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

// SendAckSupportCredits sends ack support credits.
// Request ID: k_EMsgClientToGCAckSupportCredits
// Response ID: k_EMsgClientToGCAckSupportCreditsResponse
// Request type: CMsgClientToGCAckSupportCredits
// Response type: CMsgClientToGCAckSupportCreditsResponse
func (d *DAC) SendAckSupportCredits(
	ctx context.Context,
	itemID uint64,
) (*protocol.CMsgClientToGCAckSupportCreditsResponse, error) {
	req := &protocol.CMsgClientToGCAckSupportCredits{
		ItemId: &itemID,
	}
	resp := &protocol.CMsgClientToGCAckSupportCreditsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCAckSupportCredits),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCAckSupportCreditsResponse),
		resp,
	)
}

// SendAssociateDevice sends a associate device.
// Request ID: k_EMsgClientToGCAssociateDevice
// Response ID: k_EMsgClientToGCAssociateDeviceResponse
// Request type: CMsgClientToGCAssociateDevice
// Response type: CMsgClientToGCAssociateDeviceResponse
func (d *DAC) SendAssociateDevice(
	ctx context.Context,
	deviceID uint64,
	platform protocol.EDACPlatform,
	salt uint32,
) (*protocol.CMsgClientToGCAssociateDeviceResponse, error) {
	req := &protocol.CMsgClientToGCAssociateDevice{
		DeviceId: &deviceID,
		Platform: &platform,
		Salt:     &salt,
	}
	resp := &protocol.CMsgClientToGCAssociateDeviceResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCAssociateDevice),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCAssociateDeviceResponse),
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
	req *protocol.CMsgClientToGCDevEventOperation,
) (*protocol.CMsgClientToGCDevEventOperationResponse, error) {
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

// SendDevOperation sends a dev operation.
// Request ID: k_EMsgClientToGCDevOperation
// Response ID: k_EMsgClientToGCDevOperationResponse
// Request type: CMsgClientToGCDevOperation
// Response type: CMsgClientToGCDevOperationResponse
func (d *DAC) SendDevOperation(
	ctx context.Context,
	op protocol.CMsgClientToGCDevOperation_EOperation,
	uintValue uint64,
	strValue string,
	uintValue2 uint64,
) (*protocol.CMsgClientToGCDevOperationResponse, error) {
	req := &protocol.CMsgClientToGCDevOperation{
		Op:          &op,
		UintValue:   &uintValue,
		StrValue:    &strValue,
		UintValue_2: &uintValue2,
	}
	resp := &protocol.CMsgClientToGCDevOperationResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCDevOperation),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCDevOperationResponse),
		resp,
	)
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

// SendEventChallengeProgress sends event challenge progress.
// Request ID: k_EMsgClientToGCEventChallengeProgress
// Response ID: k_EMsgClientToGCEventChallengeProgressResponse
// Request type: CMsgClientToGCEventChallengeProgress
// Response type: CMsgClientToGCEventChallengeProgressResponse
func (d *DAC) SendEventChallengeProgress(
	ctx context.Context,
	eventID uint32,
	slotID uint32,
	sequenceID uint32,
	progress uint32,
	auditData uint64,
) (*protocol.CMsgClientToGCEventChallengeProgressResponse, error) {
	req := &protocol.CMsgClientToGCEventChallengeProgress{
		EventId:    &eventID,
		SlotId:     &slotID,
		SequenceId: &sequenceID,
		Progress:   &progress,
		AuditData:  &auditData,
	}
	resp := &protocol.CMsgClientToGCEventChallengeProgressResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventChallengeProgress),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventChallengeProgressResponse),
		resp,
	)
}

// SendEventChallengeProgressWithCurrency sends a event challenge progress with currency.
// Request ID: k_EMsgClientToGCEventChallengeProgressWithCurrency
// Response ID: k_EMsgClientToGCEventChallengeProgressWithCurrencyResponse
// Request type: CMsgClientToGCEventChallengeProgressWithCurrency
// Response type: CMsgClientToGCEventChallengeProgressWithCurrencyResponse
func (d *DAC) SendEventChallengeProgressWithCurrency(
	ctx context.Context,
	req *protocol.CMsgClientToGCEventChallengeProgressWithCurrency,
) (*protocol.CMsgClientToGCEventChallengeProgressWithCurrencyResponse, error) {
	resp := &protocol.CMsgClientToGCEventChallengeProgressWithCurrencyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventChallengeProgressWithCurrency),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventChallengeProgressWithCurrencyResponse),
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

// SendEventEquipVirtualItems sends event equip virtual items.
// Request ID: k_EMsgClientToGCEventEquipVirtualItems
// Response ID: k_EMsgClientToGCEventEquipVirtualItemsResponse
// Request type: CMsgClientToGCEventEquipVirtualItems
// Response type: CMsgClientToGCEventEquipVirtualItemsResponse
func (d *DAC) SendEventEquipVirtualItems(
	ctx context.Context,
	eventID uint32,
	equipList []*protocol.CMsgClientToGCEventEquipVirtualItems_Equip,
) (*protocol.CMsgClientToGCEventEquipVirtualItemsResponse, error) {
	req := &protocol.CMsgClientToGCEventEquipVirtualItems{
		EventId:   &eventID,
		EquipList: equipList,
	}
	resp := &protocol.CMsgClientToGCEventEquipVirtualItemsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventEquipVirtualItems),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCEventEquipVirtualItemsResponse),
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
	updateChallengeSlots []uint32,
	updateClaims []uint32,
) (*protocol.CMsgClientToGCPerformAutoActionsResponse, error) {
	req := &protocol.CMsgClientToGCPerformAutoActions{
		EventId:              &eventID,
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

// SendUpdateAccountSync sends a update account sync.
// Request ID: k_EMsgClientToGCUpdateAccountSync
// Response ID: k_EMsgClientToGCUpdateAccountSyncResponse
// Request type: CMsgClientToGCUpdateAccountSync
// Response type: CMsgClientToGCUpdateAccountSyncResponse
func (d *DAC) SendUpdateAccountSync(
	ctx context.Context,
	ids []uint32,
	values []uint32,
) (*protocol.CMsgClientToGCUpdateAccountSyncResponse, error) {
	req := &protocol.CMsgClientToGCUpdateAccountSync{
		Ids:    ids,
		Values: values,
	}
	resp := &protocol.CMsgClientToGCUpdateAccountSyncResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCUpdateAccountSync),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCUpdateAccountSyncResponse),
		resp,
	)
}

// SetPartyGameModifier sets a party game modifier.
// Request ID: k_EMsgClientToGCPartySetGameModifier
// Response ID: k_EMsgClientToGCPartySetGameModifierResponse
// Request type: CMsgClientToGCPartySetGameModifier
// Response type: CMsgClientToGCPartySetGameModifierResponse
func (d *DAC) SetPartyGameModifier(
	ctx context.Context,
	partyID uint64,
	useCustomModifier bool,
	customModifier protocol.CMsgGameModifiers,
) (*protocol.CMsgClientToGCPartySetGameModifierResponse, error) {
	req := &protocol.CMsgClientToGCPartySetGameModifier{
		PartyId:           &partyID,
		UseCustomModifier: &useCustomModifier,
		CustomModifier:    &customModifier,
	}
	resp := &protocol.CMsgClientToGCPartySetGameModifierResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartySetGameModifier),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCPartySetGameModifierResponse),
		resp,
	)
}

// SetTeamName sets a team name.
// Request ID: k_EMsgClientToGCSetTeamName
// Response ID: k_EMsgClientToGCSetTeamNameResponse
// Request type: CMsgClientToGCSetTeamName
// Response type: CMsgClientToGCSetTeamNameResponse
func (d *DAC) SetTeamName(
	ctx context.Context,
	routingID uint32,
	otherAccountID uint32,
	teamName string,
) (*protocol.CMsgClientToGCSetTeamNameResponse, error) {
	req := &protocol.CMsgClientToGCSetTeamName{
		RoutingId:      &routingID,
		OtherAccountId: &otherAccountID,
		TeamName:       &teamName,
	}
	resp := &protocol.CMsgClientToGCSetTeamNameResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCSetTeamName),
		req,
		uint32(protocol.EGCDACClientMessages_k_EMsgClientToGCSetTeamNameResponse),
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
	regionMode protocol.EDACRegionMode,
) (*protocol.CMsgClientToGCSpectateUserResponse, error) {
	req := &protocol.CMsgClientToGCSpectateUser{
		SpectateAccountId: &spectateAccountID,
		RegionMode:        &regionMode,
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
	req *protocol.CMsgClientToGCPartyStartMatch,
) (*protocol.CMsgClientToGCPartyStartMatchResponse, error) {
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
	d.handlers[uint32(protocol.EGCDACClientMessages_k_EMsgGCToClientAcquireRegionModeInfo)] = d.getEventEmitter(func() events.Event {
		return &events.AcquireRegionModeInfo{}
	})
	d.handlers[uint32(protocol.EGCDACClientMessages_k_EMsgGCToClientCanRejoinParty)] = d.getEventEmitter(func() events.Event {
		return &events.CanRejoinParty{}
	})
	d.handlers[uint32(protocol.EGCDACClientMessages_k_EMsgGCToClientDevMMStatus)] = d.getEventEmitter(func() events.Event {
		return &events.DevMMStatus{}
	})
	d.handlers[uint32(protocol.EGCDACClientMessages_k_EMsgGCToClientDurationControlWarning)] = d.getEventEmitter(func() events.Event {
		return &events.DurationControlWarning{}
	})
	d.handlers[uint32(protocol.EGCDACClientMessages_k_EMsgGCToClientEventInfo)] = d.getEventEmitter(func() events.Event {
		return &events.EventInfo{}
	})
	d.handlers[uint32(protocol.EGCDACClientMessages_k_EMsgGCToClientGameModifiersUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.GameModifiersUpdated{}
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
