package pfcp_test

import (
	"testing"

	"github.com/dot-5g/cow-edge/internal/pfcp"
	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
)

type MockPfcpClient struct {
	client.Pfcp
	Sent bool // Indicates whether a response was sent
}

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

	if len(context.KnownNodeIDs) != 1 {
		t.Fatalf("Expected 1 known node ID, got %d", len(context.KnownNodeIDs))
	}

	if !context.IsKnownNodeID(nodeID) {
		t.Fatalf("Expected node ID %v to be known", nodeID)
	}

}

func TestGivenNodeIDKnownWhenHandlePFCPAssociationSetupRequestThenNodeIDNotReAddedToContext(t *testing.T) {
	knownNodeID, err := ie.NewNodeID("1.2.3.4")
	if err != nil {
		t.Fatal(err)
	}
	context := &pfcp.UPFContext{
		KnownNodeIDs: []ie.NodeID{
			knownNodeID,
		},
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

	if len(context.KnownNodeIDs) != 1 {
		t.Fatalf("Expected 1 known node ID, got %d", len(context.KnownNodeIDs))
	}

	if !context.IsKnownNodeID(knownNodeID) {
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
