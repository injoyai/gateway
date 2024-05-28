package logic

import (
	"github.com/injoyai/conv"
	"github.com/injoyai/gateway/common"
	"github.com/injoyai/gateway/module/distribute"
	"github.com/injoyai/goutil/g"
	"github.com/injoyai/io"
	"github.com/injoyai/io/buf"
	"github.com/injoyai/logs"
	json "github.com/json-iterator/go"
)

func Init() {
	links := []*distribute.Link{
		{

			ListenType: "tcp",
			ListenPort: "8080",
			Frame:      &buf.Frame{},
			Protocol:   "dlt645",
			Topic:      "api",
			Model:      "TEST",
		},
		{
			ListenType: "mqtt",
			ListenPort: "topic",
			Frame:      &buf.Frame{},
			Protocol:   "dlt645",
			Topic:      "api",
			Model:      "TEST",
		},
	}
	for _, v := range links {
		switch v.ListenType {
		case "mqtt":
			common.Listener.ListenMQTT()
		default:
			if err := Listen(v); err != nil {
				logs.Err(err)
			}
		}

	}

}

func Listen(link *distribute.Link) error {

	// 1. 数据监听
	_, err := common.Listener.Listen(link.ListenType, link.ListenPort, func(s *io.Server) {

		//2. 数据分包
		s.SetReadFunc(link.Frame.ReadMessage)

		s.SetDealFunc(func(c *io.Client, msg io.Message) {

			//尝试解析成Message
			m := &distribute.Message{
				ID:   g.UUID(),
				Link: []*distribute.Link{link},
			}
			if json.Unmarshal(msg.Bytes(), m) == nil {
				m.Link = append(m.Link, link)
				msg = conv.Bytes(m.Data["bytes"])
			}

			//3. 协议解析
			exist, data, err := common.Protocol.Decode(link.Protocol, msg)
			if err != nil {
				logs.Err(err)
				return
			}

			//4. 数据推送
			if exist {
				m.Protocol = link.Protocol
			}
			m.Model = link.Model
			m.Data = data
			logs.PrintErr(common.Pusher.Publish(link.Topic, m))

		})

		//

	})

	return err
}
