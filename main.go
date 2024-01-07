package main

import (
	"fmt"

	"github.com/dot-5g/pfcp/messages"
	"github.com/dot-5g/pfcp/server"
)

func HandleHeartbeatRequest(sequenceNumber uint32, msg messages.HeartbeatRequest) {
	fmt.Printf("Received Heartbeat Request - Recovery TimeStamp: %v", msg.RecoveryTimeStamp)
}

func main() {
	pfcpServer := server.New("localhost:8805")
	pfcpServer.HeartbeatRequest(HandleHeartbeatRequest)
	pfcpServer.Run()
	defer pfcpServer.Close()
}
