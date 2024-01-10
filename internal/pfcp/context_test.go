package pfcp_test

import (
	"testing"

	"github.com/dot-5g/cow-edge/internal/pfcp"
	"github.com/dot-5g/pfcp/ie"
)

func TestGivenInitialContextWhenGetKnownNodeIDsThenReturnsNil(t *testing.T) {
	context := &pfcp.UPFContext{}

	knownNodeIDs := context.GetKnownNodeIDs()

	if knownNodeIDs != nil {
		t.Fatalf("Expected nil, got %v", knownNodeIDs)
	}
}

func TestGivenContextWithKnownNodeIDsWhenGetKnownNodeIDsThenReturnsKnownNodeIDs(t *testing.T) {
	knownNodeID, err := ie.NewNodeID("1.2.3.4")
	if err != nil {
		t.Fatal(err)
	}

	context := &pfcp.UPFContext{
		KnownNodeIDs: []ie.NodeID{
			knownNodeID,
		},
	}

	knownNodeIDs := context.GetKnownNodeIDs()

	if len(knownNodeIDs) != 1 {
		t.Fatalf("Expected 1 known node ID, got %d", len(knownNodeIDs))
	}

}
