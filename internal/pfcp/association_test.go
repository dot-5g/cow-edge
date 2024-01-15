package pfcp_test

import (
	"testing"

	"github.com/dot-5g/cow-edge/internal/pfcp"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
)

func (m *MockPfcpClient) SendPFCPAssociationSetupResponse(response messages.PFCPAssociationSetupResponse, sequenceNumber uint32) error {
	m.Sent = true
	return nil
}

func (m *MockPfcpClient) SendPFCPAssociationReleaseResponse(response messages.PFCPAssociationReleaseResponse, sequenceNumber uint32) error {
	m.Sent = true
	return nil
}

func TestGivenNodeIDNotKnownWhenHandlePFCPAssociationSetupRequestThenNodeIDAddedToContext(t *testing.T) {
	context := &pfcp.UPFContext{}
	sequenceNumber := uint32(1)
	nodeID, err := ie.NewNodeID("1.2.3.4")
	if err != nil {
		t.Fatal(err)
	}
	message := messages.PFCPAssociationSetupRequest{
		NodeID: nodeID,
	}

	pfcpClient := &MockPfcpClient{}

	pfcp.HandlePFCPAssociationSetupRequest(context, pfcpClient, sequenceNumber, message)

	if len(context.PFCPAssociations) != 1 {
		t.Fatalf("Expected 1 PFCP association, got %d", len(context.PFCPAssociations))
	}

	if !context.IsKnownPFCPAssociation(nodeID) {
		t.Fatalf("Expected node ID %v to be known", nodeID)
	}

}

func TestGivenNodeIDKnownWhenHandlePFCPAssociationSetupRequestThenNodeIDNotReAddedToContext(t *testing.T) {
	knownNodeID, err := ie.NewNodeID("1.2.3.4")
	if err != nil {
		t.Fatal(err)
	}
	pfcpAssociation := pfcp.PFCPAssociation{
		NodeID: knownNodeID,
	}

	context := &pfcp.UPFContext{
		PFCPAssociations: []*pfcp.PFCPAssociation{&pfcpAssociation},
	}
	sequenceNumber := uint32(1)
	nodeID, err := ie.NewNodeID("1.2.3.4")
	if err != nil {
		t.Fatal(err)
	}
	message := messages.PFCPAssociationSetupRequest{
		NodeID: nodeID,
	}
	pfcpClient := &MockPfcpClient{}

	pfcp.HandlePFCPAssociationSetupRequest(context, pfcpClient, sequenceNumber, message)

	if len(context.PFCPAssociations) != 1 {
		t.Fatalf("Expected 1 PFCP association, got %d", len(context.PFCPAssociations))
	}

	if !context.IsKnownPFCPAssociation(knownNodeID) {
		t.Fatalf("Expected node ID %v to be known", knownNodeID)
	}

}

func TestGivenNodeIDNotKnownWhenHandlePFCPAssociationSetupRequestThenResponseIsSent(t *testing.T) {
	context := &pfcp.UPFContext{}

	sequenceNumber := uint32(1)
	nodeID, err := ie.NewNodeID("1.2.3.4")
	if err != nil {
		t.Fatal(err)
	}
	message := messages.PFCPAssociationSetupRequest{
		NodeID: nodeID,
	}

	// Use the mock PFCP client
	pfcpClient := &MockPfcpClient{}

	pfcp.HandlePFCPAssociationSetupRequest(context, pfcpClient, sequenceNumber, message)

	// Check if the response was sent
	if !pfcpClient.Sent {
		t.Fatalf("Expected a response to be sent, but it was not")
	}
}
