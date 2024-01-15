package pfcp

import (
	"log"
	"net"
	"time"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
)

func getNodeAddress(nodeID ie.NodeID) string {
	var nodeIDaddress string
	if nodeID.Type == ie.IPv4 || nodeID.Type == ie.IPv6 {
		nodeIDaddress = net.IP(nodeID.Value).String()
	} else {
		nodeIDaddress = string(nodeID.Value)
	}
	return nodeIDaddress
}

func HandlePFCPAssociationSetupRequest(upfContext *UPFContext, pfcpClient client.PfcpClienter, sequenceNumber uint32, msg messages.PFCPAssociationSetupRequest) {
	remoteNodeIDAddress := getNodeAddress(msg.NodeID)
	log.Printf("Received PFCP Association Setup Request from Node %v", remoteNodeIDAddress)
	remoteNodeID := msg.NodeID
	if upfContext.IsKnownPFCPAssociation(remoteNodeID) {
		log.Printf("Node ID %v is already known\n", remoteNodeIDAddress)
		return
	}
	newAssociation := PFCPAssociation{
		NodeID: remoteNodeID,
	}
	upfContext.AddPFCPAssociation(newAssociation)
	cause, err := ie.NewCause(ie.RequestAccepted)
	if err != nil {
		log.Fatal(err)
		return
	}
	recoveryTimeStamp, err := ie.NewRecoveryTimeStamp(time.Now())
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

func HandlePFCPAssociationReleaseRequest(upfContext *UPFContext, pfcpClient client.PfcpClienter, sequenceNumber uint32, msg messages.PFCPAssociationReleaseRequest) {
	//TODO: Delete all the PFCP sessions related to that PFCP association locally;
	remoteNodeIDAddress := getNodeAddress(msg.NodeID)
	log.Printf("Received PFCP Association Release Request from Node %v", remoteNodeIDAddress)

	//Delete the PFCP association and any related information (e.g. Node ID of the CP function)
	remoteNodeID := msg.NodeID
	if upfContext.IsKnownPFCPAssociation(remoteNodeID) {
		upfContext.RemovePFCPAssociation(remoteNodeID)
	}

	//Send a PFCP Association Release Response with a successful cause.
	cause, err := ie.NewCause(ie.RequestAccepted)
	if err != nil {
		log.Fatal(err)
		return
	}
	pfcpAssociationReleaseResponse := messages.PFCPAssociationReleaseResponse{
		NodeID: upfContext.NodeID,
		Cause:  cause,
	}
	pfcpClient.SendPFCPAssociationReleaseResponse(pfcpAssociationReleaseResponse, sequenceNumber)
	log.Printf("Node ID %v removed from known node IDs\n", remoteNodeIDAddress)
}
