package distribute

import (
	"github.com/injoyai/gateway/module/client"
)

// Link 链路
type Link struct {
	Name      string           `json:"name"`     //
	Model     string           `json:"model"`    //型号,例设备型号
	Subscribe []*client.Config `json:"-"`        //监听
	Protocol  string           `json:"protocol"` //协议解析的标识,可以为空,即不解析
	Publish   []*client.Config `json:"-"`        //推送
}

type Handler interface {
	Do(bs []byte) ([]byte, error)
}
