package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func handlepackets(p *gopacket.PacketSource, conn *tls.Conn) {
	for packet := range p.Packets() {
		conn.Write([]byte(packet.Data()))
	}
}

func main() {

	rootPEM := `-----BEGIN CERTIFICATE-----
MIICvjCCAaagAwIBAgIJAM4fjuiCdhDyMA0GCSqGSIb3DQEBCwUAMBQxEjAQBgNV
BAMTCTEyNy4wLjAuMTAeFw0xODEwMDgwMTI0MDdaFw0yODEwMDUwMTI0MDdaMBQx
EjAQBgNVBAMTCTEyNy4wLjAuMTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC
ggEBALtl0lKQT3S5ROIAi24YD1JbgNEsBqYEBQUX5PHFRlE1Bxb6xhaV+uuPv5WQ
rrq9VITLvqYQp7QvdJyjTlADZlNlKyxeZma2tVfb9NcTg8s1pUr7PVnNP8vQt23b
bWnz17MMHE61lOItOUm4vcv/5+1vGr/tLVjwFu4ass3HCxWGFipAhRlZ4TLW9PVJ
UnzMQlaRHgfUFfRBk3dhFwpaGQ1dGH7vBPkaZxiBhpftZyrivZlop9IE36OdaAmR
H5G41eEaLtbGQgTKMOFhqLwIBLTipQ2vvrXbGVVFWuDVmcmoGSunAk1HMdFsHJVG
a6j7frjzrWdKk/rz5no5n1lyvrsCAwEAAaMTMBEwDwYDVR0RBAgwBocEfwAAATAN
BgkqhkiG9w0BAQsFAAOCAQEAesrFXdty+rGMBDK22rhiqv/cDDmyzLGP+dDtcu7N
6Cy7lFX4LQ/NRaePcO7EdEEDw4IkFNLFjHp2/l6cJDHt11dh58DLOJF/DmCbrjYl
bQniL0WsHaG0bfChJVtSGr/PhyLOZdrGrhNmf1W0Qul/rApCuyTHAb7rHjLFp4JU
3i7t32zXCgRixM9pH9QjXEikSFWyVJkRGDPF7zOWJTxnpqAi3nkMROpsZjYz7ioG
d/TiryifZbUseQIJyVN2sf/+xym7Nf+Pr0N3MLDPpSwZ2klHEWFOQqLAvwdbnaAk
X+DareRG2QiUII3RtVhESZtVGQeiy8rqNFr/jYGNa/DUYQ==
-----END CERTIFICATE-----
`

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}
	tlsCfg := &tls.Config{RootCAs: roots}

	conn, err := tls.Dial("tcp", "127.0.0.1:8081", tlsCfg)
	if err != nil {
		log.Fatal(err)
	}

	// Set up listener on interface.
	// Device arg input.
	//if len(os.Args) != 2 {
	//	fmt.Println("Invalid Interface Reference!")
	//	return
	// }
	iface := "en0"
	// Device Handler
	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		panic(err)
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	handlepackets(packetSource, conn)
}
