package pfcp

import (
	"bytes"

	"github.com/dot-5g/pfcp/ie"
)

type SessionContext struct {
	CPFSEID   ie.FSEID
	CreatePDR ie.CreatePDR
	CreateFAR ie.CreateFAR
}

type PFCPAssociation struct {
	NodeID   ie.NodeID
	Sessions []SessionContext
}

type UPFContext struct {
	NodeID           ie.NodeID
	PFCPAssociations []*PFCPAssociation
}

func (pfcpAssociation *PFCPAssociation) AddPFCPSession(session SessionContext) {
	pfcpAssociation.Sessions = append(pfcpAssociation.Sessions, session)
}

func (upfContext *UPFContext) AddPFCPAssociation(association PFCPAssociation) {
	upfContext.PFCPAssociations = append(upfContext.PFCPAssociations, &association)
}

func (upfContext *UPFContext) GetPFCPAssociation(nodeID ie.NodeID) *PFCPAssociation {
	for _, id := range upfContext.PFCPAssociations {
		if id.NodeID.Type == nodeID.Type && bytes.Equal(id.NodeID.Value, nodeID.Value) {
			return id
		}
	}
	return nil
}

func (upfContext *UPFContext) RemovePFCPAssociation(nodeID ie.NodeID) {
	for i, id := range upfContext.PFCPAssociations {
		if id.NodeID.Type == nodeID.Type && bytes.Equal(id.NodeID.Value, nodeID.Value) {
			upfContext.PFCPAssociations = append(upfContext.PFCPAssociations[:i], upfContext.PFCPAssociations[i+1:]...)
		}
	}
}

func (upfContext *UPFContext) IsKnownPFCPAssociation(nodeID ie.NodeID) bool {
	for _, id := range upfContext.PFCPAssociations {
		if id.NodeID.Type == nodeID.Type && bytes.Equal(id.NodeID.Value, nodeID.Value) {
			return true
		}
	}
	return false
}
