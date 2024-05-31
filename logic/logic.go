package logic

import (
	"github.com/injoyai/conv"
	"github.com/injoyai/gateway/common"
	"github.com/injoyai/gateway/module/distribute"
	"github.com/injoyai/gateway/module/publish"
	"github.com/injoyai/logs"
	"time"
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

	for _, subscribe := range link.Subscribe {
		c, err := common.Client.Dial(subscribe)
		if err != nil {
			return err
		}
		c.OnMessage(func(c interface{}, bs []byte) (ack bool) {
			return DealMessage(link, bs)
		})
	}

	for _, publish := range link.Publish {
		c, err := common.Client.Dial(publish)
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
	m := &publish.Message{
		Model:    conv.SelectString(len(data.Model) > 0, data.Model, link.Model),
		ID:       data.ID,
		No:       data.No,
		Protocol: conv.SelectString(exist, link.Protocol, ""), //link.Protocol,
		Data:     data.Nature,
		Time:     time.Now().UnixMilli(),
	}

	//5. 数据推送
	for _, publish := range link.Publish {
		err = common.Client.Publish(publish.GetKey(), publish.Publish, m)
		logs.PrintErr(err)
	}

	return
}
