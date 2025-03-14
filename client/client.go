package client

import (
	"flag"
	"fmt"
	//	"fmt"
	"log"
	"net"

	//	"os"

	"github.com/suyash200/GoFastUPD/protocol"
)

func StartClient() {
	var protocolVar int
	var packetData string
	var seqNum int64
	//	var ipAddr net.UDPAddr

	flag.IntVar(&protocolVar, "version", 0, "identifying version")
	flag.StringVar(&packetData, "data", "dummy", "packet data")
	flag.Int64Var(&seqNum, "seqNum", 0, "sequence number")

	addr, err := net.ResolveUDPAddr("udp", "localhost:9000")
	if err != nil {
		log.Fatal("Error resolving address:", err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Error dialing UDP server:", err)
	}
	defer conn.Close()

	packet := &protocol.Packet{
		Version:    uint8(protocolVar),
		PacketType: 0, // Data
		SeqNum:     1,
		Length:     uint16(len([]byte(packetData))),
		Payload:    []byte(packetData),
	}

	fmt.Print(packet)
	protocol.SendPacket(conn, packet)
}
