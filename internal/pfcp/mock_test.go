package pfcp_test

import "github.com/dot-5g/pfcp/client"

type MockPfcpClient struct {
	client.PFCP
	Sent bool // Indicates whether a response was sent
}
