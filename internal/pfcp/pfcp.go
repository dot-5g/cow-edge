package pfcp

import (
	"bytes"
	"log"
	"net"
	"time"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
)

type UPFContext struct {
	NodeID       ie.NodeID
	KnownNodeIDs []ie.NodeID
}

var clients map[string]*client.Pfcp = make(map[string]*client.Pfcp)

func IsKnownNodeID(upfContext UPFContext, nodeID ie.NodeID) bool {
	for _, id := range upfContext.KnownNodeIDs {
		if id.Type == nodeID.Type && bytes.Equal(id.Value, nodeID.Value) {
			return true
		}
	}
	return false
}

func getClientForAddress(addr net.Addr) *client.Pfcp {
	addrStr := addr.String()
	if cl, exists := clients[addrStr]; exists {
		return cl
	}

	cl := client.New(addrStr)
	clients[addrStr] = cl
	return cl
}

func getNodeAddress(nodeID ie.NodeID) string {
	var nodeIDaddress string
	if nodeID.Type == ie.IPv4 || nodeID.Type == ie.IPv6 {
		nodeIDaddress = net.IP(nodeID.Value).String()
	} else {
		nodeIDaddress = string(nodeID.Value)
	}

	return nodeIDaddress
}

func HandlePFCPAssociationSetupRequest(upfContext *UPFContext, addr net.Addr, sequenceNumber uint32, msg messages.PFCPAssociationSetupRequest) {
	var cause ie.Cause
	var recoveryTimeStamp ie.RecoveryTimeStamp
	var err error

	pfcpClient := getClientForAddress(addr)
	remoteNodeIDAddress := getNodeAddress(msg.NodeID)
	log.Printf("Received PFCP Association Setup Request from Node %v", remoteNodeIDAddress)
	remoteNodeID := msg.NodeID
	if IsKnownNodeID(*upfContext, remoteNodeID) {
		log.Printf("Node ID %v is already known\n", remoteNodeIDAddress)
		return
	}
	upfContext.KnownNodeIDs = append(upfContext.KnownNodeIDs, msg.NodeID)
	cause, err = ie.NewCause(ie.RequestAccepted)
	if err != nil {
		log.Fatal(err)
		return
	}
	recoveryTimeStamp, err = ie.NewRecoveryTimeStamp(time.Now())
	if err != nil {
		log.Fatal(err)
		return
	}
	pfcpAssociationSetupResponse := messages.PFCPAssociationSetupResponse{
		NodeID:            upfContext.NodeID,
		Cause:             cause,
		RecoveryTimeStamp: recoveryTimeStamp,
	}
	pfcpClient.SendPFCPAssociationSetupResponse(pfcpAssociationSetupResponse, sequenceNumber)
	log.Printf("Node ID %v added to known node IDs\n", remoteNodeIDAddress)
}
