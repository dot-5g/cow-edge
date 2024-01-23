package pfcp

import (
	"log"
	"time"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
)

func HandlePFCPAssociationSetupRequest(upfContext *UPFContext, pfcpClient client.PfcpClienter, sequenceNumber uint32, msg messages.PFCPAssociationSetupRequest) {
	remoteNodeIDAddress := msg.NodeID.String()
	log.Printf("Received PFCP Association Setup Request from Node %v", remoteNodeIDAddress)
	if upfContext.IsKnownPFCPAssociation(remoteNodeIDAddress) {
		log.Printf("Node ID %v is already known\n", remoteNodeIDAddress)
		return
	}
	newAssociation := PFCPAssociation{
		NodeID: remoteNodeIDAddress,
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
	upfNodeID, err := ie.NewNodeID(upfContext.NodeID)
	if err != nil {
		log.Fatal(err)
		return
	}
	pfcpAssociationSetupResponse := messages.PFCPAssociationSetupResponse{
		NodeID:            upfNodeID,
		Cause:             cause,
		RecoveryTimeStamp: recoveryTimeStamp,
	}
	err = pfcpClient.SendPFCPAssociationSetupResponse(pfcpAssociationSetupResponse, sequenceNumber)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Node ID %v added to known node IDs\n", remoteNodeIDAddress)
}

func HandlePFCPAssociationReleaseRequest(upfContext *UPFContext, pfcpClient client.PfcpClienter, sequenceNumber uint32, msg messages.PFCPAssociationReleaseRequest) {
	//TODO: Delete all the PFCP sessions related to that PFCP association locally;
	remoteNodeIDAddress := msg.NodeID.String()
	log.Printf("Received PFCP Association Release Request from Node %v", remoteNodeIDAddress)

	//Delete the PFCP association and any related information (e.g. Node ID of the CP function)
	if upfContext.IsKnownPFCPAssociation(remoteNodeIDAddress) {
		upfContext.RemovePFCPAssociation(remoteNodeIDAddress)
	}

	//Send a PFCP Association Release Response with a successful cause.
	cause, err := ie.NewCause(ie.RequestAccepted)
	if err != nil {
		log.Fatal(err)
		return
	}
	upfNodeID, err := ie.NewNodeID(upfContext.NodeID)
	if err != nil {
		log.Fatal(err)
		return
	}
	pfcpAssociationReleaseResponse := messages.PFCPAssociationReleaseResponse{
		NodeID: upfNodeID,
		Cause:  cause,
	}
	err = pfcpClient.SendPFCPAssociationReleaseResponse(pfcpAssociationReleaseResponse, sequenceNumber)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Node ID %v removed from known node IDs\n", remoteNodeIDAddress)
}
