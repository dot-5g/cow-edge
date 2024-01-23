package packet

import (
	"fmt"
	"log"
	"os"

	"github.com/dot-5g/cow-edge/internal/pfcp"
	"github.com/google/gopacket"
	"github.com/google/gopacket/afpacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const bufferSize = 8

func CapturePackets(interfaceName string, upfContext *pfcp.UPFContext) {
	tpacket, err := setupPacketCapture(interfaceName)
	if err != nil {
		log.Fatalf("setupPacketCapture: %v", err)
	}
	defer tpacket.Close()

	packetSource := gopacket.NewPacketSource(tpacket, layers.LinkTypeEthernet)
	for packet := range packetSource.Packets() {
		go processPacket(packet, upfContext)
	}
}

func setupPacketCapture(interfaceName string) (*afpacket.TPacket, error) {
	szFrame, szBlock, numBlocks, err := computeAfpacketSize(bufferSize)
	if err != nil {
		return nil, fmt.Errorf("computeAfpacketSize: %w", err)
	}

	tpacket, err := afpacket.NewTPacket(
		afpacket.OptInterface(interfaceName),
		afpacket.OptFrameSize(szFrame),
		afpacket.OptBlockSize(szBlock),
		afpacket.OptNumBlocks(numBlocks),
		afpacket.OptPollTimeout(pcap.BlockForever),
	)
	if err != nil {
		return nil, fmt.Errorf("could not read packets on interface %s: %w", interfaceName, err)
	}
	return tpacket, nil
}

func computeAfpacketSize(targetSizeMb int) (int, int, int, error) {
	frameSize := os.Getpagesize()
	blockSize := frameSize * 128
	numBlocks := (targetSizeMb * 1024 * 1024) / blockSize

	if numBlocks == 0 {
		return 0, 0, 0, fmt.Errorf("interface buffer size is too small")
	}

	return frameSize, blockSize, numBlocks, nil
}

func processPacket(packet gopacket.Packet, upfContext *pfcp.UPFContext) {
	pfcpSession := upfContext.GetPFCPSession()
	pfcpSessionPDR := pfcpSession.GetPDRWithHighestPrecedence()
	if pfcpSessionPDR == nil {
		return
	}
	applyPDRInstructions(*pfcpSessionPDR, packet)
}

// Apply Instructions set in the PDR
func applyPDRInstructions(pdr pfcp.PDR, packet gopacket.Packet) {
	// Implement the actions as per the PDR
	// For example, forwarding the packet, modifying it, or dropping it

	fmt.Printf("PDR ID: %v\n", pdr.PDRID)
	fmt.Printf("Precedence: %v\n", pdr.Precedence)
	fmt.Printf("PDI: %v\n", pdr.PDI)

}
