package built

import (
	"github.com/injoyai/gateway/module/protocol/internal/common"
	"github.com/injoyai/goutil/g"
	"github.com/injoyai/protocol/dlt645"
)

type Dlt645 struct{}

// Decode 解析协议,测试数据:68AAAAAAAAAAAA68910833333333343333337E16
func (this *Dlt645) Decode(bs []byte) (*common.Message, error) {
	p, err := dlt645.Decode(bs)
	if err != nil {
		return nil, err
	}
	p.Mark = p.Mark.Sub0x33()
	f, err := p.Result()
	if err != nil {
		return nil, err
	}
	return &common.Message{
		No:     p.No,
		Nature: g.M{p.Mark.HEX(): f},
	}, nil
}
