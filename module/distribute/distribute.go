package distribute

import (
	"github.com/injoyai/goutil/g"
	"github.com/injoyai/io/buf"
	json "github.com/json-iterator/go"
)

// Link 链路
type Link struct {
	ListenAddr string     `json:"listenAddr"` //
	ListenType string     `json:"listenType"` //监听类型
	ListenPort string     `json:"listenPort"` //监听端口
	Frame      *buf.Frame `json:"-" yaml:"-"` //分包配置
	Protocol   string     //协议解析的标识,可以为空,即不解析
	Topic      string     //推送的topic //需要可变,根据解析的数据进行变化
	Model      string     `json:"model"` //型号,例设备型号
}

// Message 消息
type Message struct {
	ID string `json:"id"` //消息id,uuid,全局唯一,防止重复

	Link     []*Link `json:"link"`     //链路信息
	Protocol string  `json:"protocol"` //协议解析的标识
	Data     g.Map   `json:"data"`     //数据

	//Key   string `json:"key"`   //消息标识,这批设备的编号,例某个型号的某个版本
	Model string `json:"model"` //型号,例设备型号
	//No    string `json:"no"`    //编号,例设备唯一编号
}

func (this *Message) Bytes() []byte {
	bs, _ := json.Marshal(this)
	return bs
}
