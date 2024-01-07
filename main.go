package main

import (
	"bytes"
	"log"
	"net"

	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/server"
)

type PFCPAssociation struct {
	NodeIDs []ie.NodeID
}

var pfcpAssociation PFCPAssociation

func IsKnownNodeID(nodeID ie.NodeID) bool {
	for _, id := range pfcpAssociation.NodeIDs {
		if id.NodeIDType == nodeID.NodeIDType && bytes.Equal(id.NodeIDValue, nodeID.NodeIDValue) {
			return true
		}
	}
	return false
}

func HandlePFCPAssociationSetupRequest(sequenceNumber uint32, msg messages.PFCPAssociationSetupRequest) {
	var nodeIDAddress string
	nodeIDType := msg.NodeID.NodeIDType
	if nodeIDType == ie.IPv4 {
		nodeIDAddress = net.IP(msg.NodeID.NodeIDValue).String()
	}
	log.Printf("Received PFCP Association Setup Request from Node %v", nodeIDAddress)
	nodeID := msg.NodeID
	if IsKnownNodeID(nodeID) {
		log.Printf("Node ID %v is already known\n", nodeID)
		return
	}

	pfcpAssociation.NodeIDs = append(pfcpAssociation.NodeIDs, msg.NodeID)
	log.Printf("Node ID %v added to known node IDs\n", nodeIDAddress)
}

func main() {
	pfcpServer := server.New("localhost:8805")
	pfcpServer.PFCPAssociationSetupRequest(HandlePFCPAssociationSetupRequest)
	pfcpServer.Run()
	defer pfcpServer.Close()
}
