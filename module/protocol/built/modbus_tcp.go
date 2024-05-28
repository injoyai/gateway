package built

import (
	"encoding/hex"
	"errors"
	"github.com/injoyai/conv"
	"github.com/injoyai/gateway/module/protocol/built/internal/modbus"
	"github.com/injoyai/goutil/g"
)

type ModbusTCP struct{}

func (this *ModbusTCP) Decode(bs []byte) (g.Map, error) {
	length := len(bs)
	if length < 9 {
		return nil, errors.New("数据长度异常(小于9):" + hex.EncodeToString(bs))
	}

	f := &modbus.TCPFrame{
		Order:    [2]byte{bs[0], bs[1]},
		Protocol: [2]byte{bs[2], bs[3]},
		Length:   [2]byte{bs[4], bs[5]},
		Slave:    bs[6],
		Control:  modbus.Control(bs[7]),
		Data:     bs[8:],
	}
	if f.Control.Byte() > 0x80 {
		return nil, f.Control
	}

	dataLen := conv.Int(f.Length)
	if dataLen != len(f.Data)+2 {
		return nil, errors.New("数据长度错误:" + f.HEX())
	}

	m := g.Map{}
	for i := range g.Range(dataLen, 2) {
		m[conv.String(i)] = f.Data[i : i+2]
	}

	return g.Map{
		"id":   conv.String(conv.Int(f.Order)),
		"no":   string(f.Slave),
		"data": m,
	}, nil

}
