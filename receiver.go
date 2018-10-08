package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

func handle(conn net.Conn) {
	reader, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(reader)
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
