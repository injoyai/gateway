package register

import (
	"github.com/injoyai/base/maps"
	"github.com/injoyai/conv"
	"time"
)

func New() *Manager {
	return &Manager{maps.NewSafe()}
}

type Manager struct {
	m *maps.Safe
}

func (this *Manager) Set(group string, key string, value interface{}) {
	m, _ := this.m.GetOrSetByHandler(group, func() (interface{}, error) {
		return maps.NewSafe(), nil
	})
	v, _ := m.(*maps.Safe).GetOrSetByHandler(key, func() (interface{}, error) {
		return &Register{Group: group, Key: key}, nil
	})
	v.(*Register).Value = value
	v.(*Register).Time = time.Now()
}

func (this *Manager) Get(group string, key string) *Register {
	m, ok := this.m.Get(group)
	if !ok {
		return nil
	}
	i := m.(*maps.Safe).MustGet(key)
	if i == nil {
		return nil
	}
	return i.(*Register)
}

type Register struct {
	Group string      //分组
	Key   string      //唯一标识
	Type  string      //数据类型,如int,float,string,bool
	Value interface{} //实时值
	Time  time.Time   //数据时间
}

func (this *Register) GetValue() interface{} {
	switch this.Type {
	case "string":
		return conv.String(this.Value)
	case "int":
		return conv.Int64(this.Value)
	case "float":
		return conv.Float64(this.Value)
	case "bool":
		return conv.Bool(this.Value)
	default:
		return this.Value
	}
}
