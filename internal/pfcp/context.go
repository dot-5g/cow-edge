package pfcp

import (
	"bytes"

	"github.com/dot-5g/pfcp/ie"
)

type UPFContext struct {
	NodeID       ie.NodeID
	KnownNodeIDs []ie.NodeID
}

func (upfContext *UPFContext) GetKnownNodeIDs() []ie.NodeID {
	return upfContext.KnownNodeIDs
}

func (upfContext *UPFContext) AddKnownNodeID(nodeID ie.NodeID) {
	upfContext.KnownNodeIDs = append(upfContext.KnownNodeIDs, nodeID)
}

func (upfContext *UPFContext) RemoveKnownNodeID(nodeID ie.NodeID) {
	for i, id := range upfContext.KnownNodeIDs {
		if id.Type == nodeID.Type && bytes.Equal(id.Value, nodeID.Value) {
			upfContext.KnownNodeIDs = append(upfContext.KnownNodeIDs[:i], upfContext.KnownNodeIDs[i+1:]...)
			break
		}
	}
}

func (upfContext *UPFContext) IsKnownNodeID(nodeID ie.NodeID) bool {
	for _, id := range upfContext.KnownNodeIDs {
		if id.Type == nodeID.Type && bytes.Equal(id.Value, nodeID.Value) {
			return true
		}
	}
	return false
}
