package dac

import (
	bgcm "github.com/paralin/go-underlords/protocol"
)

// LeavePartyGCBase attempts to leave the current party using the GC base party system.
func (d *DAC) LeavePartyGCBase() {
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCLeaveParty), &bgcm.CMsgLeaveParty{})
}
