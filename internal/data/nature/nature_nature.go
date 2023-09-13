package nature

import (
	"github.com/injoyai/conv"
	"github.com/injoyai/goutil/g"
	"sync"
)

type Nature interface {
	GetType() string                                      //数据类型,系统,设备,用户
	GetKey() string                                       //数据唯一标识,opc
	GetName() string                                      //数据名称
	GetMemo() string                                      //数据备注
	GetUnit() string                                      //数据单位
	GetValue() interface{}                                //数据值
	GetValueType() string                                 //消息数据类型,STRING,FLOAT,INT,BOOL
	GetValueTime() int64                                  //数据时间
	GetReadable() bool                                    //是否可读
	GetReadValue() (interface{}, error)                   //读取数据
	GetWritable() bool                                    //是否可写
	GetWriteValue(value interface{}, byUser string) error //写入数据
	Copy() Nature                                         //复制属性
}

type NatureExtend interface {
	Nature
	GetValueString() string
	GetValueBool() bool
	GetValueInt() int64
	GetValueFloat() float64
	GetValueVar() *conv.Var
}

type Natures []Nature

func (this Natures) GetByKey(key string) Nature {
	for _, nature := range this {
		if nature.GetKey() == key {
			return nature
		}
	}
	return nil
}

func (this Natures) GMap() g.Map {
	m := g.Map{}
	for _, v := range this {
		m[v.GetKey()] = v.GetValue()
	}
	return m
}

func (this Natures) Conv() conv.Extend {
	c := &_conv{m: make(map[string]interface{})}
	c.Extend = conv.NewExtend(c)
	for _, v := range this {
		c.m[v.GetKey()] = v.GetValue()
	}
	return c
}

type _conv struct {
	conv.Extend
	m  map[string]interface{}
	mu sync.RWMutex
}

func (this *_conv) GetVar(key string) *conv.Var {
	this.mu.RLock()
	defer this.mu.RUnlock()
	return conv.New(this.m[key])
}
