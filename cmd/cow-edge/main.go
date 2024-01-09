package main

import (
	"log"
	"net"

	"github.com/dot-5g/cow-edge/internal/config"
	"github.com/dot-5g/cow-edge/internal/pfcp"
	"github.com/dot-5g/pfcp/ie"
	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/server"
)

var nodeID ie.NodeID

func main() {
	var err error
	config, err := config.ReadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
	pfcpServer := server.New("localhost:8805")
	nodeID, err = ie.NewNodeID(config.UPF.NodeID)
	if err != nil {
		log.Fatal(err)
		return
	}
	pfcpContext := &pfcp.UPFContext{NodeID: nodeID}

	pfcpServer.PFCPAssociationSetupRequest(func(addr net.Addr, sequenceNumber uint32, msg messages.PFCPAssociationSetupRequest) {
		pfcp.HandlePFCPAssociationSetupRequest(pfcpContext, addr, sequenceNumber, msg)
	})
	pfcpServer.Run()
	defer pfcpServer.Close()
}
