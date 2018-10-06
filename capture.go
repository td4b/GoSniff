package main

import (
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {

	// decoder objects
	var ipv4 layers.IPv4
	var eth layers.Ethernet
	var tcp layers.TCP
	var udp layers.UDP
	var dns layers.DNS

	// Device arg input.
	if len(os.Args) != 2 {
		fmt.Println("Invalid Interface Reference!")
		return
	}
	iface := os.Args[1]
	// Device Handler
	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		panic(err)
	}

	// Packet Decoder.
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	parser := gopacket.NewDecodingLayerParser(
		layers.LayerTypeEthernet,
		&eth,
		&ipv4,
		&tcp,
		&udp,
		&dns,
	)
	decoded := []gopacket.LayerType{}
	for packet := range packetSource.Packets() {
		_ = parser.DecodeLayers(packet.Data(), &decoded)
		// Detects if a packet has flags set for an HTTP/HTTPS message stream.
		//payload := string(ipv4.Payload)
		SrcIP := ipv4.SrcIP.String()
		DstIP := ipv4.DstIP.String()
		switch ipv4.Protocol.String() {
		case "TCP":
			Port := tcp.DstPort.String()
			fmt.Println(SrcIP, DstIP, "TCP", Port)
		case "UDP":
			Port := udp.DstPort.String()
			fmt.Println(SrcIP, DstIP, "UDP", Port)
		}
	}
	defer handle.Close()
}
