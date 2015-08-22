package packets

import (
	"encoding/binary"
	"io"
)

const (
	CONNECT     = 1
	CONNACK     = 2
	PUBLISH     = 3
	PUBACK      = 4
	PUBREC      = 5
	PUBREL      = 6
	PUBCOMP     = 7
	SUBSCRIBE   = 8
	SUBACK      = 9
	UNSUBSCRIBE = 10
	UNSUBACK    = 11
	PINGREQ     = 12
	PINGRESP    = 13
	DISCONNECT  = 14
)

func ReadPacket(r io.Reader) {
	//var h = FixedHeader
}

func boolToByte(b bool) byte {
	switch b {
	case true:
		return 1
	default:
		return 0
	}
}

func decodeByte(b io.Reader) byte {
	num := make([]byte, 1)
	b.Read(num)
	return num[0]
}

func decodeUint16(b io.Reader) uint16 {
	num := make([]byte, 2)
	b.Read(num)
	return binary.BigEndian.Uint16(num)
}

func encodeUint16(num uint16) []byte {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, num)
	return bytes
}

func encodeString(field string) []byte {
	fieldLength := make([]byte, 2)
	binary.BigEndian.PutUint16(fieldLength, uint16(len(field)))
	return append(fieldLength, []byte(field)...)
}

func decodeString(b io.Reader) string {
	fieldLength := decodeUint16(b)
	field := make([]byte, fieldLength)
	b.Read(field)
	return string(field)
}

func decodeBytes(b io.Reader) []byte {
	fieldLength := decodeUint16(b)
	field := make([]byte, fieldLength)
	b.Read(field)
	return field
}

func encodeBytes(field []byte) []byte {
	fieldLength := make([]byte, 2)
	binary.BigEndian.PutUint16(fieldLength, uint16(len(field)))
	return append(fieldLength, field...)
}

func encodeLength(length int) []byte {
	var encLength []byte
	for {
		digit := byte(length % 128)
		length /= 128
		if length > 0 {
			digit |= 0x80
		}
		encLength = append(encLength, digit)
		if length == 0 {
			break
		}
	}
	return encLength
}

func decodeLength(r io.Reader) int {
	var rLength uint32
	var multiplier uint32 = 0
	b := make([]byte, 1)
	for {
		io.ReadFull(r, b)
		digit := b[0]
		rLength |= uint32(digit&127) << multiplier
		if (digit & 128) == 0 {
			break
		}
		multiplier += 7
	}
	return int(rLength)
}
