package logic

import (
	"github.com/injoyai/conv"
	"github.com/injoyai/gateway/common"
	"github.com/injoyai/gateway/module/distribute"
	"github.com/injoyai/goutil/g"
	"github.com/injoyai/logs"
	json "github.com/json-iterator/go"
)

func Init(links []*distribute.Link) error {
	for _, v := range links {
		logs.Tracef("默认配置: %#v\n", v)
		if err := Listen(v); err != nil {
			return err
		}
	}
	return nil
}

func Listen(link *distribute.Link) error {

	{
		c, err := common.Client.Dial(link.Subscribe)
		if err != nil {
			return err
		}
		c.OnMessage(func(c interface{}, bs []byte) (ack bool) {
			return DealMessage(link, bs)
		})
	}

	{
		c, err := common.Client.Dial(link.Publish)
		if err != nil {
			return err
		}
		c.OnMessage(func(c interface{}, bs []byte) (ack bool) {
			return DealMessage(link, bs)
		})
	}

	return nil
}

func Dial(link *distribute.Link) error {

	c, err := common.Client.Dial(link.Publish)
	if err != nil {
		return err
	}
	c.OnMessage(func(c interface{}, bs []byte) (ack bool) {
		return DealMessage(link, bs)
	})
	return nil
}

func DealMessage(link *distribute.Link, msg []byte) (ack bool) {
	//3. 协议解析
	//尝试解析成Message
	m := &distribute.Message{
		ID:   g.UUID(),
		Link: []*distribute.Link{link},
	}
	logs.Debug(string(msg))
	if json.Unmarshal(msg, m) == nil {
		m.Link = append(m.Link, link)
		msg = conv.Bytes(m.Data["bytes"])
	}

	switch {
	case len(link.Protocol) == 0:
		//不解析,直接推送

	default:
		//解析协议
		exist, data, err := common.Protocol.Decode(link.Protocol, msg)
		if err != nil {
			logs.Err(err)
			return
		}
		if exist {
			//fmt.Sprintf(" ->%s(%s)", link.Protocol, conv.SelectString(exist, "成功", "协议不存在"))
			m.Protocol += link.Protocol
		}
		m.Data = data

	}

	//4. 数据推送
	m.Model = link.Model
	logs.PrintErr(common.Client.Publish(link.Publish.GetKey(), link.Publish.Publish, m))
	return
}
