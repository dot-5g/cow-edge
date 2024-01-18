package main

import (
	"flag"
	"log"

	"github.com/dot-5g/cow-edge/internal/config"
	"github.com/dot-5g/cow-edge/internal/packet"
	"github.com/dot-5g/cow-edge/internal/pfcp"
	"github.com/dot-5g/pfcp/client"
	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/server"
)

var configFilePath string

func init() {
	flag.StringVar(&configFilePath, "config", "config.yaml", "Path to the config file")
}

func main() {
	flag.Parse()
	config, err := config.ReadConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
	pfcpServer := server.New("localhost:8805")
	upfContext := &pfcp.UPFContext{NodeID: config.NodeID}

	go packet.CapturePackets(config.Interface, upfContext)

	pfcpServer.PFCPAssociationSetupRequest(func(pfcpClient *client.Pfcp, sequenceNumber uint32, msg messages.PFCPAssociationSetupRequest) {
		pfcp.HandlePFCPAssociationSetupRequest(upfContext, pfcpClient, sequenceNumber, msg)
	})
	pfcpServer.PFCPAssociationReleaseRequest(func(pfcpClient *client.Pfcp, sequenceNumber uint32, msg messages.PFCPAssociationReleaseRequest) {
		pfcp.HandlePFCPAssociationReleaseRequest(upfContext, pfcpClient, sequenceNumber, msg)
	})
	pfcpServer.PFCPSessionEstablishmentRequest(func(pfcpClient *client.Pfcp, sequenceNumber uint32, seid uint64, msg messages.PFCPSessionEstablishmentRequest) {
		pfcp.HandlePFCPSessionEstablishmentRequest(upfContext, pfcpClient, sequenceNumber, seid, msg)
	})
	pfcpServer.Run()
	defer pfcpServer.Close()
}
