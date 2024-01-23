package pfcp_test

import (
	"testing"

	"github.com/dot-5g/cow-edge/internal/pfcp"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
)

func (m *MockPfcpClient) SendPFCPSessionEstablishmentResponse(response messages.PFCPSessionEstablishmentResponse, seid uint64, sequenceNumber uint32) error {
	m.Sent = true
	return nil
}

func TestGivenUnknownSessionWhenHandlePFCPSessionEstablishmentRequestThenSessionAddedToContext(t *testing.T) {
	nodeIDValue := "1.2.3.4"
	nodeID, err := ie.NewNodeID(nodeIDValue)
	if err != nil {
		t.Fatal(err)
	}
	pfcpAssociation := pfcp.PFCPAssociation{
		NodeID: nodeIDValue,
	}
	upfContext := &pfcp.UPFContext{
		PFCPAssociations: []*pfcp.PFCPAssociation{&pfcpAssociation},
	}
	pfcpClient := &MockPfcpClient{}
	sequenceNumber := uint32(1)
	seid := uint64(1)

	msg := messages.PFCPSessionEstablishmentRequest{
		NodeID: nodeID,
	}

	pfcp.HandlePFCPSessionEstablishmentRequest(upfContext, pfcpClient, sequenceNumber, seid, msg)

	if len(upfContext.PFCPAssociations) != 1 {
		t.Fatalf("Expected 1 PFCP association, got %d", len(upfContext.PFCPAssociations))
	}

	if !upfContext.IsKnownPFCPAssociation(nodeIDValue) {
		t.Fatalf("Expected node ID %v to be known", nodeID)
	}

	if len(upfContext.PFCPAssociations[0].Sessions) != 1 {
		t.Fatalf("Expected 1 PFCP session, got %d", len(upfContext.PFCPAssociations[0].Sessions))
	}
}

func TestGivenUnknownSessionWhenHandlePFCPSessionEstablishmentRequestThenPFCPResponseIsSent(t *testing.T) {
	nodeIDValue := "1.2.3.4"
	nodeID, err := ie.NewNodeID(nodeIDValue)
	if err != nil {
		t.Fatal(err)
	}
	pfcpAssociation := pfcp.PFCPAssociation{
		NodeID: nodeIDValue,
	}
	upfContext := &pfcp.UPFContext{
		PFCPAssociations: []*pfcp.PFCPAssociation{&pfcpAssociation},
	}
	pfcpClient := &MockPfcpClient{}
	sequenceNumber := uint32(1)
	seid := uint64(1)

	msg := messages.PFCPSessionEstablishmentRequest{
		NodeID: nodeID,
	}

	pfcp.HandlePFCPSessionEstablishmentRequest(upfContext, pfcpClient, sequenceNumber, seid, msg)

	if !pfcpClient.Sent {
		t.Errorf("Expected PFCP response to be sent")
	}
}
