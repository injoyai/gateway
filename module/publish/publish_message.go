package publish

import (
	"github.com/injoyai/goutil/g"
	json "github.com/json-iterator/go"
)

// Message 消息
type Message struct {
	Model    string `json:"model"`    //型号,例设备型号
	ID       string `json:"id"`       //消息id,uuid,全局唯一,防止重复
	No       string `json:"no"`       //设备唯一标识
	Protocol string `json:"protocol"` //协议解析的标识
	Data     g.Map  `json:"data"`     //数据
	Time     int64  `json:"time"`     //时间,毫秒
}

func (this *Message) Bytes() []byte {
	bs, _ := json.Marshal(this)
	return bs
}
