package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {

	// temporary port 80 filter.
	var filter = flag.String("f", "tcp and port 80", "BPF filter for pcap")

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
	if err := handle.SetBPFFilter(*filter); err != nil {
		log.Fatal("error setting BPF filter: ", err)
	}
	gopacket.PacketDataSource()
	// Packet Decoder.
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	defer handle.Close()
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
		switch ipv4.Protocol.String() {
		case "TCP":
			fmt.Println(ipv4.SrcIP.String(), ipv4.DstIP.String(), "TCP", tcp.DstPort.String())
			if app := packet.ApplicationLayer(); app != nil {
				fmt.Println("Payload: " + string(app.Payload()))
			} else {
				fmt.Println("Payload: None")
			}
		case "UDP":
			fmt.Println(ipv4.SrcIP.String(), ipv4.DstIP.String(), "UDP", udp.DstPort.String())
		}
	}
}
