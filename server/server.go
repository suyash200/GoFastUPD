package server

import (
	"fmt"
	"github.com/suyash200/GoFastUPD/protocol"
	"log"
)

// StartServer: Initializes the UDP server
func StartServer() {
	addr := ":9000"
	conn, err := protocol.SetupUDP(addr, true)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	defer conn.Close()

	fmt.Println("UDP Server is listening on", addr)

	for {
		packet, err := protocol.ReceivePacket(conn)
		if err != nil {
			log.Println("Error receiving packet:", err)
			continue
		}
		fmt.Printf("Received packet: %+v\n", packet)
	}
}
