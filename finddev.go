package main

import (
	"fmt"
	//"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func getinfo(i pcap.Interface) {
	for j := 0; j < len(i.Addresses); j++ {
		fmt.Println(i.Name, i.Addresses[j].IP)
	}
}

func main() {
	value, _ := pcap.FindAllDevs()
	for i := 0; i < len(value); i++ {
		getinfo(value[i])
	}
}
