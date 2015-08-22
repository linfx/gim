package packets

import (
	"bytes"
	"io"
)

type ConnAckMessage struct {
	FixedHeader
	ReturnCode byte
}

func (p *ConnAckMessage) WriteTo(w io.Writer) error {
	var body bytes.Buffer

	body.WriteByte(p.ReturnCode)
	packet := p.FixedHeader.Pack()
	packet.Write(body.Bytes())

	_, err := packet.WriteTo(w)
	return err
}

func (p *ConnAckMessage) ReadForm(r io.Reader) {
	p.ReturnCode = decodeByte(r)
}
