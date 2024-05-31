package client

import "github.com/injoyai/io"

type Client interface {
	Publish(topic string, data interface{}) error
	OnMessage(f MessageFunc)
	io.Closer
}

type MessageFunc func(c interface{}, bs []byte) (ack bool)
