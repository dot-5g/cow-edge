package pfcp_test

import (
	"testing"

	"github.com/dot-5g/cow-edge/internal/pfcp"
)

func TestGivenNoPFCPAssociationWhenGetPFCPAssociationThenReturnsNil(t *testing.T) {
	upfContext := &pfcp.UPFContext{}
	nodeID := "1.2.3.4"

	pfcpAssociation := upfContext.GetPFCPAssociation(nodeID)

	if pfcpAssociation != nil {
		t.Fatalf("Expected nil PFCP association, got %v", pfcpAssociation)
	}
}

func TestGivenExistingPFCPAssociationWhenGetPFCPAssociationThenReturnsAssociation(t *testing.T) {
	nodeIDValue := "2.3.4.5"
	pfcpAssociation := pfcp.PFCPAssociation{
		NodeID: nodeIDValue,
	}
	upfContext := &pfcp.UPFContext{
		PFCPAssociations: []*pfcp.PFCPAssociation{&pfcpAssociation},
	}

	newPfcpAssociation := upfContext.GetPFCPAssociation(nodeIDValue)

	if newPfcpAssociation == nil {
		t.Fatalf("Expected PFCP association, got nil")
	}
	if newPfcpAssociation.NodeID != nodeIDValue {
		t.Fatalf("Expected PFCP association node ID %v, got %v", nodeIDValue, newPfcpAssociation.NodeID)
	}
}

func TestGivenWhenRemovePFCPAssociationThenAssociationRemoved(t *testing.T) {
	nodeIDValue := "2.3.4.5"
	pfcpAssociation := pfcp.PFCPAssociation{
		NodeID: nodeIDValue,
	}
	upfContext := &pfcp.UPFContext{
		PFCPAssociations: []*pfcp.PFCPAssociation{&pfcpAssociation},
	}

	upfContext.RemovePFCPAssociation(nodeIDValue)

	if len(upfContext.PFCPAssociations) != 0 {
		t.Fatalf("Expected 0 PFCP associations, got %d", len(upfContext.PFCPAssociations))
	}
}
