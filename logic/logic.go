package logic

import (
	"github.com/injoyai/conv"
	"github.com/injoyai/gateway/common"
	"github.com/injoyai/gateway/module/distribute"
	"github.com/injoyai/logs"
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

func DealMessage(link *distribute.Link, msg []byte) (ack bool) {
	ack = true

	//3. 协议解析
	exist, data, err := common.Protocol.Decode(link.Protocol, msg)
	if err != nil {
		logs.Err(err)
		return
	}

	//4. 数据整理
	m := &distribute.Message{
		Model:    conv.SelectString(len(data.Model) > 0, data.Model, link.Model),
		ID:       data.ID,
		No:       data.No,
		Protocol: conv.SelectString(exist, link.Protocol, ""), //link.Protocol,
		Data:     data.Nature,
	}

	//5. 数据推送
	err = common.Client.Publish(link.Publish.GetKey(), link.Publish.Publish, m)
	logs.PrintErr(err)

	return
}
