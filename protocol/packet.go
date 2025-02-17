package protocol

import (
	"encoding/binary"
	"errors"
)

// Packet structure
type Packet struct {
	Version    uint8
	PacketType uint8
	SeqNum     uint32
	Length     uint16
	Payload    []byte
}

// Marshal converts Packet struct into a byte slice
func (p *Packet) Marshal() ([]byte, error) {
	// Preallocate buffer size: fixed header (8 bytes) + payload
	buf := make([]byte, 8+len(p.Payload))

	// Writing fields directly to buffer
	buf[0] = p.Version
	buf[1] = p.PacketType
	binary.BigEndian.PutUint32(buf[2:6], p.SeqNum)
	binary.BigEndian.PutUint16(buf[6:8], p.Length)

	// Copy payload if present
	if len(p.Payload) > 0 {
		copy(buf[8:], p.Payload)
	}

	return buf, nil
}

// Unmarshal converts a byte slice into a Packet struct
func Unmarshal(data []byte) (*Packet, error) {
	if len(data) < 8 {
		return nil, errors.New("invalid packet size")
	}

	packet := &Packet{
		Version:    data[0],
		PacketType: data[1],
		SeqNum:     binary.BigEndian.Uint32(data[2:6]),
		Length:     binary.BigEndian.Uint16(data[6:8]),
	}

	// Validate Length field
	if int(packet.Length) > len(data)-8 {
		return nil, errors.New("invalid payload length")
	}

	// Allocate payload slice only if necessary
	if packet.Length > 0 {
		packet.Payload = make([]byte, packet.Length)
		copy(packet.Payload, data[8:])
	}

	return packet, nil
}
