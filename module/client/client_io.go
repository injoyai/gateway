package client

import (
	"github.com/injoyai/conv"
	"github.com/injoyai/goutil/net/http"
	"github.com/injoyai/io"
	"github.com/injoyai/io/buf"
	"github.com/injoyai/io/dial"
	"github.com/injoyai/io/listen"
)

func NewIOClient(cfg *Config, frame *buf.Frame) (*IOClient, error) {
	m := conv.NewMap(cfg.Param)
	dialFunc := dial.WithTCP(cfg.Subscribe)
	switch cfg.Type {
	case "tcp":
		dialFunc = dial.WithTCP(cfg.Subscribe)
	case "udp":
		dialFunc = dial.WithUDP(cfg.Subscribe)
	case "websocket":
		dialFunc = dial.WithWebsocket(cfg.Subscribe, func() http.Header {
			header := http.Header{}
			for k, v := range m.GetGMap("header") {
				header[k] = conv.Strings(v)
			}
			return header
		}())
	case "memory":
		dialFunc = dial.WithMemory(cfg.Subscribe)
	}
	s, err := io.NewDial(dialFunc, func(s *io.Client) {
		//2. 数据分包
		s.SetReadWithFrame(frame)
	})
	return &IOClient{Client: s}, err
}

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
	case io.TCP:
		listenFunc = listen.WithTCP(conv.Int(cfg.Subscribe))
	case io.UDP:
		listenFunc = listen.WithUDP(conv.Int(cfg.Subscribe))
	case io.Websocket:
		listenFunc = listen.WithWebsocket(conv.Int(cfg.Subscribe))
	case io.Memory:
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

func (this *IOServer) Close() error {
	return this.Server.Close()
}
