package pfcp

import (
	"log"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
)

func HandlePFCPSessionEstablishmentRequest(upfContext *UPFContext, pfcpClient client.PfcpClienter, sequenceNumber uint32, seid uint64, msg messages.PFCPSessionEstablishmentRequest) {
	// Store the rules received in the request
	log.Printf("Received PFCP Session Establishment Request")
	sessionContext := SessionContext{
		CPFSEID:   msg.CPFSEID,
		CreatePDR: msg.CreatePDR,
		CreateFAR: msg.CreateFAR,
	}

	// Apply the rules received in the request
	pfcpAssociation := upfContext.GetPFCPAssociation(msg.NodeID)
	if pfcpAssociation == nil {
		log.Printf("Node ID %v is not known\n", msg.NodeID)
		return
	}
	pfcpAssociation.AddPFCPSession(sessionContext)

	// Send a PFCP Session Establishment Response with cause "success", if all rules in the PFCP Session Establishment Request are stored and applied
	causeValue := ie.RequestAccepted
	cause, err := ie.NewCause(causeValue)
	if err != nil {
		log.Fatal(err)
		return
	}
	pfcpSessionEstablishmentResponse := messages.PFCPSessionEstablishmentResponse{
		NodeID: upfContext.NodeID,
		Cause:  cause,
	}
	pfcpClient.SendPFCPSessionEstablishmentResponse(pfcpSessionEstablishmentResponse, seid, sequenceNumber)
}
