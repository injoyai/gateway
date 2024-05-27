package register

import (
	"github.com/injoyai/conv"
	"time"
)

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
