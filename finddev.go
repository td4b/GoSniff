package main

import (
	"fmt"
	"log"
	"net"
	//"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Function determines primary network ip by resolving DNS.
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

//We then iterate through our interface list to return the network adapter.
// The below methods work but do not seem ideal will need to tweak.

func getint(i pcap.Interface) string {
	intip := GetOutboundIP().String()
	name := ""
	for j := 0; j < len(i.Addresses); j++ {
		if i.Addresses[j].IP.String() == intip {
			name = i.Name
		}
	}
	return name
}

func findnetinerface() string {
	value, _ := pcap.FindAllDevs()
	result := ""
	for i := 0; i < len(value); i++ {
		if getint(value[i]) != "" {
			result = getint(value[i])
		}
	}
	return result
}

func main() {
	inter := findnetinerface()
	fmt.Println(inter)
}
