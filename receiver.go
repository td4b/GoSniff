package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	_ "github.com/mattn/go-sqlite3"
)

func update(id int, sourceip string, destip string, protocol string, port string) {
	db, _ := sql.Open("sqlite3", "./data.db")
	defer db.Close()
	query, _ := db.Prepare(`
	insert into ipflows (id, sourceip, destip, protocol, port) values (?, ?, ?, ?, ?)
	`)
	defer query.Close()
	query.Exec(id, sourceip, destip, protocol, port)

}

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
	buf := make([]byte, 256)
	count := 0
	for {
		_, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		_ = parser.DecodeLayers(buf, &decoded)

		// This catches HTTP data for parsing. =)
		// Will print out Netflow information for non-HTTP flows.

		// The data at this point should be handled appropriatley.
		// We should pass the data to buffered I/O so it can be handled/stored in a local Database system.

		// for now throttling data to only HTTP for DB storage.
		if tcp.DstPort.String() == "80(http)" {
			update(count, ipv4.SrcIP.String(), ipv4.DstIP.String(), ipv4.Protocol.String(), tcp.DstPort.String())
			count++
		} else {
			// for now discard non HTTP flows.
			continue
		}
	}
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
