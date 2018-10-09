package main

import (
  "fmt"
  //"github.com/google/gopacket"
  "github.com/google/gopacket/pcap"
)

func main() {
    value, _ := pcap.FindAllDevs()
    fmt.Println(value)
}
