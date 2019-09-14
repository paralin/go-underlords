package dac

import (
	"github.com/faceit/go-steam/protocol/gamecoordinator"
	gcsdkm "github.com/paralin/go-underlords/protocol"
	gcsm "github.com/paralin/go-underlords/protocol"
)

// RequestCacheSubscriptionRefresh requests a subscription refresh for a specific cache ID.
func (d *DAC) RequestCacheSubscriptionRefresh(ownerSoid *gcsdkm.CMsgSOIDOwner) {
	d.write(uint32(gcsm.ESOMsg_k_ESOMsg_CacheSubscriptionRefresh), &gcsdkm.CMsgSOCacheSubscriptionRefresh{
		OwnerSoid: ownerSoid,
	})
}

// handleCacheSubscribed handles a CacheSubscribed packet.
func (d *DAC) handleCacheSubscribed(packet *gamecoordinator.GCPacket) error {
	sub := &gcsdkm.CMsgSOCacheSubscribed{}
	if err := d.unmarshalBody(packet, sub); err != nil {
		return err
	}

	if err := d.cache.HandleSubscribed(sub); err != nil {
		d.le.WithError(err).Debug("unhandled cache issue (ignore)")
	}

	return nil
}

// handleCacheUnsubscribed handles a CacheUnsubscribed packet.
func (d *DAC) handleCacheUnsubscribed(packet *gamecoordinator.GCPacket) error {
	sub := &gcsdkm.CMsgSOCacheUnsubscribed{}
	if err := d.unmarshalBody(packet, sub); err != nil {
		return err
	}

	if err := d.cache.HandleUnsubscribed(sub); err != nil {
		d.le.WithError(err).Debug("unhandled cache issue (ignore)")
	}

	return nil
}

// handleCacheUpdateMultiple handles when one or more object(s) in a cache is/are updated.
func (d *DAC) handleCacheUpdateMultiple(packet *gamecoordinator.GCPacket) error {
	sub := &gcsdkm.CMsgSOMultipleObjects{}
	if err := d.unmarshalBody(packet, sub); err != nil {
		return err
	}

	if err := d.cache.HandleUpdateMultiple(sub); err != nil {
		d.le.WithError(err).Debug("unhandled cache issue (ignore)")
	}

	return nil
}

// handleCacheDestroy handles when an object in a cache is destroyed.
func (d *DAC) handleCacheDestroy(packet *gamecoordinator.GCPacket) error {
	sub := &gcsdkm.CMsgSOSingleObject{}
	if err := d.unmarshalBody(packet, sub); err != nil {
		return err
	}

	if err := d.cache.HandleDestroy(sub); err != nil {
		d.le.WithError(err).Debug("unhandled cache issue (ignore)")
	}

	return nil
}
