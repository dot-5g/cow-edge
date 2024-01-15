package pfcp_test

import (
	"bytes"
	"testing"

	"github.com/dot-5g/cow-edge/internal/pfcp"
	"github.com/dot-5g/pfcp/ie"
)

func TestGivenNoPFCPAssociationWhenGetPFCPAssociationThenReturnsNil(t *testing.T) {
	upfContext := &pfcp.UPFContext{}
	nodeID, err := ie.NewNodeID("1.2.3.4")
	if err != nil {
		t.Fatal(err)
	}

	pfcpAssociation := upfContext.GetPFCPAssociation(nodeID)

	if pfcpAssociation != nil {
		t.Fatalf("Expected nil PFCP association, got %v", pfcpAssociation)
	}
}

func TestGivenExistingPFCPAssociationWhenGetPFCPAssociationThenReturnsAssociation(t *testing.T) {
	nodeID, err := ie.NewNodeID("2.3.4.5")
	if err != nil {
		t.Fatal(err)
	}
	pfcpAssociation := pfcp.PFCPAssociation{
		NodeID: nodeID,
	}
	upfContext := &pfcp.UPFContext{
		PFCPAssociations: []*pfcp.PFCPAssociation{&pfcpAssociation},
	}

	newPfcpAssociation := upfContext.GetPFCPAssociation(nodeID)

	if newPfcpAssociation == nil {
		t.Fatalf("Expected PFCP association, got nil")
	}

	if !bytes.Equal(newPfcpAssociation.NodeID.Value, nodeID.Value) {
		t.Fatalf("Expected PFCP association node ID %v, got %v", nodeID, newPfcpAssociation.NodeID)
	}

}

func TestGivenWhenRemovePFCPAssociationThenAssociationRemoved(t *testing.T) {
	nodeID, err := ie.NewNodeID("1.2.3.4")
	if err != nil {
		t.Fatal(err)
	}
	pfcpAssociation := pfcp.PFCPAssociation{
		NodeID: nodeID,
	}

	upfContext := &pfcp.UPFContext{
		PFCPAssociations: []*pfcp.PFCPAssociation{&pfcpAssociation},
	}

	upfContext.RemovePFCPAssociation(nodeID)

	if len(upfContext.PFCPAssociations) != 0 {
		t.Fatalf("Expected 0 PFCP associations, got %d", len(upfContext.PFCPAssociations))
	}
}
