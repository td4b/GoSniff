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
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

//We then iterate through our interface list to return the network adapter.
// The below methods work but do not seem ideal will need to tweak.

func getint(i pcap.Interface) (inter string) {
	intip := GetOutboundIP().String()
	for j := 0; j < len(i.Addresses); j++ {
		if i.Addresses[j].IP.String() == intip {
			inter = i.Name
		}
	}
	return inter
}

func findnetinerface() (result string) {
	value, _ := pcap.FindAllDevs()
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
