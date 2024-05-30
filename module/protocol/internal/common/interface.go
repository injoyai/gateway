package common

import "github.com/injoyai/goutil/g"

type Decoder interface {
	Decode(bs []byte) (*Message, error)
}

type Message struct {
	Model  string `json:"model"`  //型号,可选,从协议中解析出协议是啥型号的设备
	No     string `json:"no"`     //设备唯一标识,必须,确定数据属于谁
	ID     string `json:"id"`     //消息ID,可选,主动上报没有消息id
	Nature g.Map  `json:"nature"` //属性
}

func (this *Message) Deal() {
	if this == nil {
		return
	}
	if len(this.ID) == 0 {
		//this.ID = g.RandString(16)
	}
}
