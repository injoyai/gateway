package listen

import (
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/injoyai/base/maps"
	"github.com/injoyai/conv"
	"github.com/injoyai/io"
	"github.com/injoyai/io/listen"
)

func New() *Manager {
	return &Manager{m: maps.NewSafe()}
}

type Manager struct {
	m *maps.Safe
}

func (this *Manager) Get(Type string, port string) *io.Server {
	if m := this.m.MustGet(Type); m != nil {
		s := m.(*maps.Safe).MustGet(port)
		if s != nil {
			return s.(*io.Server)
		}
	}
	return nil
}

func (this *Manager) ListenMQTT(sub *MQTTSubscribe, deal Deal, options ...func(*mqtt.ClientOptions)) (*MQTTClient, error) {
	return NewMQTTClient(sub, deal, options...)
}

func (this *Manager) Listen(Type string, port string, options ...io.OptionServer) (*io.Server, error) {

	var listenFunc io.ListenFunc
	switch Type {
	case "tcp":
		listenFunc = listen.WithTCP(conv.Int(port))
	case "udp":
		listenFunc = listen.WithUDP(conv.Int(port))
	case "websocket":
		listenFunc = listen.WithWebsocket(conv.Int(port))
	default:
		return nil, errors.New("未知监听类型:" + string(Type))
	}

	key := fmt.Sprintf("%s:%s", Type, port)
	s, err := io.NewServer(listenFunc, func(s *io.Server) {
		s.SetKey(key)
		s.SetOptions(options...)
	})
	if err != nil {
		return nil, err
	}

	m, _ := this.m.GetOrSetByHandler(Type, func() (interface{}, error) {
		return maps.NewSafe(), nil
	})

	m.(*maps.Safe).Set(key, s)
	go s.Run()

	return s, nil
}
