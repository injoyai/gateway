package client

import (
	"github.com/injoyai/conv"
	"github.com/injoyai/io"
	"github.com/injoyai/io/buf"
	"github.com/injoyai/io/listen"
)

type IOClient struct {
	*io.Client
}

func (this *IOClient) Publish(topic string, data interface{}) error {
	_, err := this.Client.WriteAny(data)
	return err
}

func (this *IOClient) OnMessage(f MessageFunc) {
	this.Client.SetDealFunc(func(c *io.Client, msg io.Message) {
		f(c, msg.Bytes())
	})
}

/*



 */

func NewIOServer(cfg *Config, frame *buf.Frame) (*IOServer, error) {
	listenFunc := listen.WithTCP(conv.Int(cfg.Subscribe))
	switch cfg.Type {
	case "tcp":
		listenFunc = listen.WithTCP(conv.Int(cfg.Subscribe))
	case "udp":
		listenFunc = listen.WithUDP(conv.Int(cfg.Subscribe))
	case "websocket":
		listenFunc = listen.WithWebsocket(conv.Int(cfg.Subscribe))
	case "memory":
		listenFunc = listen.WithMemory(cfg.Subscribe)
	}
	s, err := io.NewServer(listenFunc, func(s *io.Server) {
		//2. 数据分包
		s.SetReadFunc(frame.ReadMessage)
	})
	return &IOServer{Server: s}, err
}

type IOServer struct {
	*io.Server
}

func (this *IOServer) Publish(topic string, data interface{}) error {
	_, err := this.Server.WriteClient(topic, conv.Bytes(data))
	return err
}

func (this *IOServer) OnMessage(f MessageFunc) {
	this.Server.SetDealFunc(func(c *io.Client, msg io.Message) {
		f(c, msg.Bytes())
	})
}
