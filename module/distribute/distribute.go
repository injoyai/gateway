package distribute

import (
	"github.com/injoyai/gateway/module/client"
	"github.com/injoyai/goutil/g"
	"github.com/injoyai/io/buf"
	json "github.com/json-iterator/go"
)

// Link 链路
type Link struct {
	Model     string         `json:"-"`          //型号,例设备型号
	Subscribe *client.Config `json:"-"`          //监听
	Frame     *buf.Frame     `json:"-" yaml:"-"` //分包配置
	Protocol  string         `json:"protocol"`   //协议解析的标识,可以为空,即不解析
	Publish   *client.Config `json:"-"`          //推送
}

// Message 消息
type Message struct {
	Model string `json:"model"` //型号,例设备型号
	ID    string `json:"id"`    //消息id,uuid,全局唯一,防止重复

	Link     []*Link `json:"-"`        //链路信息
	Protocol string  `json:"protocol"` //协议解析的标识
	Data     g.Map   `json:"data"`     //数据
}

func (this *Message) Bytes() []byte {
	bs, _ := json.Marshal(this)
	return bs
}
