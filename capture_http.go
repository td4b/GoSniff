package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// function handles cleartext HTTP data only.
func http_methods(data string) bool {
	val := false
	methods := []string{"GET", "PUT", "POST"}
	for _, httpmsg := range methods {
		if strings.Contains(strings.Split(string(data), "\n")[0], httpmsg) == true {
			val = true
		}
	}
	return val
}

func main() {

	// decoder objects
	var ipv4 layers.IPv4
	var eth layers.Ethernet
	var tcp layers.TCP

	// Device arg input.
	if len(os.Args) != 2 {
		fmt.Println("Invalid Interface Reference!")
		return
	}
	iface := os.Args[1]

	// Device Handler
	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	defer handle.Close()
	if err != nil {
		panic(err)
	}

	// Packet Decoder.
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &ipv4, &tcp)
	decoded := []gopacket.LayerType{}
	for packet := range packetSource.Packets() {
		_ = parser.DecodeLayers(packet.Data(), &decoded)
		// Detects if a packet has flags set for an HTTP/HTTPS message stream.
		if tcp.PSH == true && tcp.ACK == true {
			if http_methods(string(ipv4.Payload)) == true {
				payload := string(ipv4.Payload)[20:]
				fmt.Println(payload)
			} else {
				// catches encrypted data.
				fmt.Println("### Encrypted Alert ###", string(ipv4.Payload))
		}
		}
	}
}
