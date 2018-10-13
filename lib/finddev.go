package finddev

import (
	"errors"
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

//We then iterate through our interface map to get the correct live interface.
func Findnetinerface() (intn string, err error) {
	value, _ := pcap.FindAllDevs()
	m := make(map[string][]pcap.InterfaceAddress)
	for i := 0; i < len(value); i++ {
		m[string(value[i].Name)] = value[i].Addresses
	}
	intip := GetOutboundIP().String()
	for key, val := range m {
		if len(val) != 0 {
			for i := 0; i < len(val); i++ {
				if intip == val[i].IP.String() {
					intn = key
				}
			}
		}
	}
	if intn == "" {
		err := errors.New("No Interface Found.")
		return intn, err
	}
	return intn, err
}
