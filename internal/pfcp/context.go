// Package pfcp provides the PFCP context and the logic to handle PFCP messages.
package pfcp

type SourceInterface struct {
	Value int
}

type PDI struct {
	SourceInterface SourceInterface
}

type PDRID struct {
	RuleID uint16
}

type Precedence struct {
	Value uint32
}

type FSEID struct {
	V4   bool
	V6   bool
	SEID uint64
	IPv4 []byte
	IPv6 []byte
}

type PDR struct {
	PDRID      PDRID
	Precedence Precedence
	PDI        PDI
}

type FARID struct {
	Value uint32
}

type ApplyAction struct {
	DFRT bool
	IPMD bool
	IPMA bool
	DUPL bool
	NOCP bool
	BUFF bool
	FORW bool
	DROP bool
	DDPN bool
	BDPN bool
	EDRT bool
}

type FAR struct {
	FARID       FARID
	ApplyAction ApplyAction
}

type SessionContext struct {
	CPFSEID FSEID
	PDR     PDR
	FAR     FAR
}

type PFCPAssociation struct {
	NodeID   string
	Sessions []SessionContext
}

type UPFContext struct {
	NodeID           string
	PFCPAssociations []*PFCPAssociation
}

func (pfcpAssociation *PFCPAssociation) AddPFCPSession(session SessionContext) {
	pfcpAssociation.Sessions = append(pfcpAssociation.Sessions, session)
}

func (upfContext *UPFContext) AddPFCPAssociation(association PFCPAssociation) {
	upfContext.PFCPAssociations = append(upfContext.PFCPAssociations, &association)
}

func (upfContext *UPFContext) GetPFCPAssociation(nodeID string) *PFCPAssociation {
	for _, id := range upfContext.PFCPAssociations {
		if id.NodeID == nodeID {
			return id
		}
	}
	return nil
}

func (upfContext *UPFContext) GetPFCPSession() *SessionContext {
	// TODO: Implement logic to get the PFCP session from the UPF context
	return &upfContext.PFCPAssociations[0].Sessions[0]
}

func (session *SessionContext) GetPDRWithHighestPrecedence() *PDR {
	// Implement logic to get the PDR with the highest precedence from the session context
	return nil
}

func (upfContext *UPFContext) RemovePFCPAssociation(nodeID string) {
	for i, id := range upfContext.PFCPAssociations {
		if id.NodeID == nodeID {
			upfContext.PFCPAssociations = append(upfContext.PFCPAssociations[:i], upfContext.PFCPAssociations[i+1:]...)
		}
	}
}

func (upfContext *UPFContext) IsKnownPFCPAssociation(nodeID string) bool {
	for _, id := range upfContext.PFCPAssociations {
		if id.NodeID == nodeID {
			return true
		}
	}
	return false
}
