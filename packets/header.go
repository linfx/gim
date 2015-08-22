package packets

import (
	"bytes"
)

type FixedHeader struct {
	MessageType     byte
	Dup             bool
	Qos             byte
	Retain          bool
	RemainingLength int
}

func (h *FixedHeader) Pack() bytes.Buffer {
	var header bytes.Buffer
	header.WriteByte(h.MessageType<<4 | boolToByte(h.Dup)<<3 | h.Qos<<1 | boolToByte(h.Retain))
	header.Write(encodeLength(h.RemainingLength))
	return header
}

func (h *FixedHeader) UnPack() {
}
