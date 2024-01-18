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
		CPFSEID: FSEID{
			V4:   msg.CPFSEID.V4,
			V6:   msg.CPFSEID.V6,
			SEID: msg.CPFSEID.SEID,
			IPv4: msg.CPFSEID.IPv4,
			IPv6: msg.CPFSEID.IPv6,
		},
		PDR: PDR{
			PDRID: PDRID{
				RuleID: msg.CreatePDR.PDRID.RuleID,
			},
			Precedence: Precedence{
				Value: msg.CreatePDR.Precedence.Value,
			},
			PDI: PDI{
				SourceInterface: SourceInterface{
					Value: msg.CreatePDR.PDI.SourceInterface.Value,
				},
			},
		},
		FAR: FAR{
			FARID: FARID{
				Value: msg.CreateFAR.FARID.Value,
			},
			ApplyAction: ApplyAction{
				DFRT: msg.CreateFAR.ApplyAction.DFRT,
				IPMD: msg.CreateFAR.ApplyAction.IPMD,
				IPMA: msg.CreateFAR.ApplyAction.IPMA,
				DUPL: msg.CreateFAR.ApplyAction.DUPL,
				NOCP: msg.CreateFAR.ApplyAction.NOCP,
				BUFF: msg.CreateFAR.ApplyAction.BUFF,
				FORW: msg.CreateFAR.ApplyAction.FORW,
				DROP: msg.CreateFAR.ApplyAction.DROP,
				DDPN: msg.CreateFAR.ApplyAction.DDPN,
				BDPN: msg.CreateFAR.ApplyAction.BDPN,
				EDRT: msg.CreateFAR.ApplyAction.EDRT,
			},
		},
	}

	// Apply the rules received in the request
	cpNodeID := NodeIDToString(msg.NodeID)
	pfcpAssociation := upfContext.GetPFCPAssociation(cpNodeID)
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
	upfNodeID, err := ie.NewNodeID(upfContext.NodeID)
	if err != nil {
		log.Fatal(err)
		return
	}
	pfcpSessionEstablishmentResponse := messages.PFCPSessionEstablishmentResponse{
		NodeID: upfNodeID,
		Cause:  cause,
	}
	pfcpClient.SendPFCPSessionEstablishmentResponse(pfcpSessionEstablishmentResponse, seid, sequenceNumber)
}
