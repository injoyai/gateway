package built

import (
	"github.com/injoyai/goutil/g"
	"github.com/injoyai/logs"
	"github.com/injoyai/protocol/dlt645"
)

type Dlt645 struct{}

// Decode 解析协议,测试数据:68AAAAAAAAAAAA68910833333333343333337E16
func (this *Dlt645) Decode(bs []byte) (g.Map, error) {
	p, err := dlt645.Decode(bs)
	if err != nil {
		return nil, err
	}
	logs.Debug(p.Data.HEX())
	f, err := p.Result()
	if err != nil {
		return nil, err
	}
	return g.M{p.Mark.HEX(): f}, nil
}
