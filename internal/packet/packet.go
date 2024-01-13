package packet

import (
	"log"
	"net"

	"github.com/dot-5g/cow-edge/internal/pfcp"
	"github.com/dot-5g/pfcp/ie"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func CapturePackets(interfaceName string, upfContext *pfcp.UPFContext) {
	handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		go processPacket(packet, upfContext) // Pass the UPF context to processPacket
	}
}

func processPacket(packet gopacket.Packet, upfContext *pfcp.UPFContext) {
	// Extract packet details
	srcIP, dstIP := getIPAddresses(packet)
	srcPort, dstPort := getPortNumbers(packet)
	protocol := getProtocol(packet)

	// Iterate over sessions in the UPF context
	for _, session := range upfContext.Sessions {
		pdr := session.CreatePDR // Assuming each session has one CreatePDR
		if matchesPDR(pdr, srcIP, dstIP, srcPort, dstPort, protocol) {
			applyPDRActions(pdr, packet)
			break // Stop after finding the first matching PDR
		}
	}
}

func matchesPDR(pdr ie.CreatePDR, srcIP, dstIP net.IP, srcPort, dstPort int, protocol string) bool {
	// Implement logic to check if the packet matches the PDR criteria
	// Compare packet details with PDR fields like source/destination IP, ports, etc.
	return true // Placeholder, implement actual matching logic
}

func applyPDRActions(pdr ie.CreatePDR, packet gopacket.Packet) {
	// Implement the actions as per the PDR
	// For example, forwarding the packet, modifying it, or dropping it
}

func getIPAddresses(packet gopacket.Packet) (net.IP, net.IP) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		return ip.SrcIP, ip.DstIP
	}
	// Handle IPv6 similarly if needed
	return nil, nil
}

func getPortNumbers(packet gopacket.Packet) (int, int) {
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		return int(tcp.SrcPort), int(tcp.DstPort)
	}
	// Handle other protocols (UDP, etc.) similarly if needed
	return 0, 0
}

func getProtocol(packet gopacket.Packet) string {
	// Implement logic to return the protocol (e.g., "TCP", "UDP")
	// This could be based on which layer is present in the packet
	return ""
}
