package pfcp

import (
	"fmt"
	"log"

	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
)

type FailedRuleID uint32

func (sessionContext *SessionContext) ApplyRules() *FailedRuleID {
	failedRuleID := ApplyCreatePDRRule(sessionContext.CreatePDR)
	if failedRuleID != nil {
		return failedRuleID
	}
	return nil
}

func ApplyCreatePDRRule(createPDR ie.CreatePDR) *FailedRuleID {
	// TODO: Apply the rules received in the request
	fmt.Printf("PDR ID: %v\n", createPDR.PDRID)
	fmt.Printf("Precedence: %v\n", createPDR.Precedence)
	fmt.Printf("PDI: %v\n", createPDR.PDI)
	return nil
}

func HandlePFCPSessionEstablishmentRequest(upfContext *UPFContext, pfcpClient client.PfcpClienter, sequenceNumber uint32, seid uint64, msg messages.PFCPSessionEstablishmentRequest) {
	// Store the rules received in the request
	log.Printf("Received PFCP Session Establishment Request")
	sessionContext := SessionContext{
		CPFSEID:   msg.CPFSEID,
		CreatePDR: msg.CreatePDR,
		CreateFAR: msg.CreateFAR,
	}

	// Apply the rules received in the request
	failedRuleID := sessionContext.ApplyRules()
	var causeValue ie.CauseValue
	if failedRuleID == nil {
		causeValue = ie.RequestAccepted
	} else {
		causeValue = ie.RequestRejected
	}

	// Send a PFCP Session Establishment Response with cause "success", if all rules in the PFCP Session Establishment Request are stored and applied
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
