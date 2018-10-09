package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func handle(conn net.Conn) {

	// decoder objects
	var ipv4 layers.IPv4
	var eth layers.Ethernet
	var tcp layers.TCP
	var udp layers.UDP
	var dns layers.DNS

	parser := gopacket.NewDecodingLayerParser(
		layers.LayerTypeEthernet,
		&eth,
		&ipv4,
		&tcp,
		&udp,
		&dns,
	)
	decoded := []gopacket.LayerType{}

	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}

		// printing out buf (bytes allocated to our buffer) prints out payload.
		// there still seems to be an issue decoding the packet however using parser.DecodeLayers()

		fmt.Println("Length of Byte Received: " + string(n))
		buf = append(buf, tmp[:n]...)

		_ = parser.DecodeLayers(buf, &decoded)
		switch ipv4.Protocol.String() {
		case "TCP":
			fmt.Println(ipv4.SrcIP.String(), ipv4.DstIP.String(), "TCP", tcp.DstPort.String())
		case "UDP":
			fmt.Println(ipv4.SrcIP.String(), ipv4.DstIP.String(), "UDP", udp.DstPort.String())
		}
	}
	// Need to somehow convert []bytes received into packet decoder interface.

}

func main() {

	cert, err := tls.LoadX509KeyPair("server.crt", "server.pem")
	if err != nil {
		log.Fatal("Error loading certificate. ", err)
	}

	tlsCfg := &tls.Config{Certificates: []tls.Certificate{cert}}

	listener, err := tls.Listen("tcp4", "127.0.0.1:8081", tlsCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		log.Println("Waiting for clients")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}
