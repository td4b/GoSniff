package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {

	// temporary port 80 filter.
	var filter = flag.String("f", "tcp", "BPF filter for pcap")

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

	// Open connection to collector.
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	// Packet Decoder.
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	defer handle.Close()
	for packet := range packetSource.Packets() {

		// send data to socket.
		conn.Write(packet.Data())

	}
}
