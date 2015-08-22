package packets

import (
	"bytes"
	"io"
)

type ConnectMessage struct {
	FixedHeader

	ProtocolName    string
	ProtocolVersion byte
	CleanSession    bool
	WillFlag        bool
	WillQos         byte
	WillRetain      bool
	UsernameFlag    bool
	PasswordFlag    bool
	KeepaliveTimer  uint16

	ClientIdentifier string
	WillTopic        string
	WillMessage      []byte
	Username         string
	Password         []byte
}

func (p *ConnectMessage) WriteTo(w io.Writer) error {
	var body bytes.Buffer
	body.Write(encodeString(p.ProtocolName))
	body.WriteByte(p.ProtocolVersion)
	body.WriteByte(boolToByte(p.CleanSession)<<1 | boolToByte(p.WillFlag)<<2 | p.WillQos<<3 | boolToByte(p.WillRetain)<<5 | boolToByte(p.PasswordFlag)<<6 | boolToByte(p.UsernameFlag)<<7)
	body.Write(encodeUint16(p.KeepaliveTimer))
	body.Write(encodeString(p.ClientIdentifier))
	if p.WillFlag {
		body.Write(encodeString(p.WillTopic))
		body.Write(encodeBytes(p.WillMessage))
	}
	if p.UsernameFlag {
		body.Write(encodeString(p.Username))
	}
	if p.PasswordFlag {
		body.Write(encodeBytes(p.Password))
	}

	p.FixedHeader.RemainingLength = body.Len()
	packet := p.FixedHeader.Pack()
	packet.Write(body.Bytes())

	_, err := packet.WriteTo(w)
	return err
}

func (p *ConnectMessage) ReadForm(r io.Reader) {
	p.ProtocolName = decodeString(r)
	p.ProtocolVersion = decodeByte(r)
	options := decodeByte(r)
	p.CleanSession = 1&(options>>1) > 0
	p.WillFlag = 1&(options>>2) > 0
	p.WillQos = 3 & (options >> 3)
	p.WillRetain = 1&(options>>5) > 0
	p.PasswordFlag = 1&(options>>6) > 0
	p.UsernameFlag = 1&(options>>7) > 0
	p.KeepaliveTimer = decodeUint16(r)
	p.ClientIdentifier = decodeString(r)
	if p.WillFlag {
		p.WillTopic = decodeString(r)
		p.WillMessage = decodeBytes(r)
	}
	if p.UsernameFlag {
		p.Username = decodeString(r)
	}
	if p.PasswordFlag {
		p.Password = decodeBytes(r)
	}
}
