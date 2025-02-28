package types

// Packet structure

type Packet struct {
	Version    uint8
	PacketType uint8
	SeqNum     uint32
	Length     uint16
	Payload    []byte
}
