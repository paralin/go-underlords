package dac

import (
	"context"
	"errors"
	"sync"

	"github.com/faceit/go-steam"
	"github.com/faceit/go-steam/protocol/gamecoordinator"
	"github.com/golang/protobuf/proto"
	devents "github.com/paralin/go-underlords/events"
	"github.com/sirupsen/logrus"
	// gcmm "github.com/paralin/go-underlords/protocol"
	bgcm "github.com/paralin/go-underlords/protocol"
	gcsdkm "github.com/paralin/go-underlords/protocol"
	gcsm "github.com/paralin/go-underlords/protocol"
	"github.com/paralin/go-underlords/socache"
	"github.com/paralin/go-underlords/state"
)

// AppID is the ID for underlords
const AppID = 1046930

// ErrNotReady is returned when the client is not ready.
var ErrNotReady = errors.New("client not ready")

// handlerMap is the map of message types to handler functions.
type handlerMap map[uint32]func(packet *gamecoordinator.GCPacket) error

// DAC handles the dota game handler.
type DAC struct {
	le     *logrus.Entry
	client *steam.Client
	cache  *socache.SOCache

	connectionCtxMtx    sync.Mutex
	connectionCtx       context.Context
	connectionCtxCancel context.CancelFunc

	stateMtx sync.Mutex
	state    state.DACState

	handlers handlerMap

	pendReqMtx sync.Mutex
	pendReqID  uint32
	pendReq    map[uint32]map[uint32]responseHandler
}

// New builds a new DAC handler.
func New(client *steam.Client, le *logrus.Entry) *DAC {
	c := &DAC{
		le:      le,
		cache:   socache.NewSOCache(le),
		client:  client,
		pendReq: make(map[uint32]map[uint32]responseHandler),

		state: state.DACState{
			ConnectionStatus: gcsdkm.GCConnectionStatus_GCConnectionStatus_NO_SESSION,
		},
	}
	c.buildHandlerMap()
	client.GC.RegisterPacketHandler(c)
	return c
}

// GetCache returns the SO Cache.
func (d *DAC) GetCache() *socache.SOCache {
	return d.cache
}

// Close kills any ongoing calls.
func (d *DAC) Close() {
	d.connectionCtxMtx.Lock()
	if d.connectionCtxCancel != nil {
		d.connectionCtxCancel()
	}
	d.connectionCtxMtx.Unlock()
}

// buildHandlerMap builds the map of bound handler functions.
func (d *DAC) buildHandlerMap() {
	d.handlers = handlerMap{
		// Welcome and conn status
		uint32(gcsm.EGCBaseClientMsg_k_EMsgGCClientWelcome):          d.handleClientWelcome,
		uint32(gcsm.EGCBaseClientMsg_k_EMsgGCClientConnectionStatus): d.handleConnectionStatus,

		// Caching
		uint32(gcsm.ESOMsg_k_ESOMsg_CacheSubscribed):   d.handleCacheSubscribed,
		uint32(gcsm.ESOMsg_k_ESOMsg_UpdateMultiple):    d.handleCacheUpdateMultiple,
		uint32(gcsm.ESOMsg_k_ESOMsg_CacheUnsubscribed): d.handleCacheUnsubscribed,
		uint32(gcsm.ESOMsg_k_ESOMsg_Destroy):           d.handleCacheDestroy,

		// System events
		uint32(gcsm.EGCBaseClientMsg_k_EMsgGCPingRequest): d.handlePingRequest,

		// Invites
		uint32(bgcm.EGCBaseMsg_k_EMsgGCInvitationCreated): d.getEventEmitter(func() devents.Event {
			return &devents.InvitationCreated{}
		}),
	}

	d.registerGeneratedHandlers()
}

// write sends a message to the game coordinator.
func (d *DAC) write(messageType uint32, msg proto.Message) {
	d.client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppID, messageType, msg))
}

// emit emits an event.
func (d *DAC) emit(event interface{}) {
	d.client.Emit(event)
}

// accessState safely accesses the DAC state. return true if the state was changed / otherwise updated during the call.
func (d *DAC) accessState(cb func(nextState *state.DACState) (bool, error)) error {
	d.stateMtx.Lock()
	defer d.stateMtx.Unlock()

	lastState := d.state
	changed, err := cb(&d.state)
	if err != nil {
		return err
	}
	if changed {
		d.emit(devents.ClientStateChanged{
			OldState: lastState,
			NewState: d.state,
		})
	}
	return nil
}

// unmarshalBody attempts to unmarshal a packet body.
func (d *DAC) unmarshalBody(packet *gamecoordinator.GCPacket, msg proto.Message) (parseErr error) {
	defer func() {
		if parseErr != nil {
			d.le.WithError(parseErr).WithField("msgtype", packet.MsgType).Warn("unable to parse message")
		}
	}()

	return proto.Unmarshal(packet.Body, msg)
}

// HandleGCPacket handles an incoming game coordinator packet.
func (d *DAC) HandleGCPacket(packet *gamecoordinator.GCPacket) {
	if packet.AppId != AppID {
		return
	}

	le := d.le.WithField("msgtype", packet.MsgType)
	handler, ok := d.handlers[packet.MsgType]
	if ok && handler != nil {
		if err := handler(packet); err != nil {
			le.WithError(err).Warn("error handling gc msg")
			ok = false
		}
	}

	respHandled := d.handleResponsePacket(le, packet)
	if !ok && !respHandled {
		le.Debug("unhandled gc packet")
		d.emit(&devents.UnhandledGCPacket{
			Packet: packet,
		})
	}
}

// handlePingRequest handles an incoming ping request from the gc.
func (d *DAC) handlePingRequest(packet *gamecoordinator.GCPacket) error {
	d.write(uint32(gcsm.EGCBaseClientMsg_k_EMsgGCPingResponse), &gcsdkm.CMsgGCClientPing{})
	return nil
}

// getEventEmitter returns a handler that emits an event, used by the generated code.
func (d *DAC) getEventEmitter(ctor func() devents.Event) func(packet *gamecoordinator.GCPacket) error {
	return func(packet *gamecoordinator.GCPacket) error {
		obj := ctor()
		if err := d.unmarshalBody(packet, obj.GetEventBody()); err != nil {
			return err
		}

		d.emit(obj)
		return nil
	}
}
