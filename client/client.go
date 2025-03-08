package client

import (
	"github.com/suyash200/GoFastUPD/protocol"
	"log"
	"net"
)

func startClient() {
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
		Version:    1,
		PacketType: 0, // Data
		SeqNum:     1,
		Length:     uint16(len([]byte("Hello, UDP!"))),
		Payload:    []byte("Hello, UDP!"),
	}

	protocol.SendPacket(conn, packet)
}
