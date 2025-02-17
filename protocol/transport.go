package protocol

import (
	"net"
)

// SendPacket sends a packet over a UDP connection
func SendPacket(conn *net.UDPConn, packet *Packet) error {
	data, err := packet.Marshal()
	if err != nil {
		return err
	}
	_, err = conn.Write(data) // No need to store bytes written
	return err
}

// ReceivePacket receives a packet over a UDP connection
func ReceivePacket(conn *net.UDPConn) (*Packet, error) {
	buffer := make([]byte, 1500) // Use MTU size (optimized for UDP)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		return nil, err
	}
	return Unmarshal(buffer[:n])
}

// SetupUDP sets up a UDP connection (server or client)
func SetupUDP(address string, isServer bool) (*net.UDPConn, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	if isServer {
		return net.ListenUDP("udp", udpAddr) // Bind to address for receiving
	}
	return net.DialUDP("udp", nil, udpAddr) // Client mode
}
