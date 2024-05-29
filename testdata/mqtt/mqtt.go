package main

import (
	"github.com/injoyai/io"
	"github.com/injoyai/io/dial"
	"time"
)

func main() {
	<-dial.RedialMQTT(&dial.MQTTIOConfig{
		Subscribe: []dial.MQTTSubscribe{{Topic: "client1"}},
		Publish:   []dial.MQTTPublish{{Topic: "server"}},
	}, dial.WithMQTTBase(&dial.MQTTBaseConfig{
		BrokerURL: "tcp://192.168.10.23:1883",
		ClientID:  "client1",
		Timeout:   time.Second * 10,
	}), func(c *io.Client) {
		c.SetPrintWithHEX()
		c.GoTimerWriter(time.Second*10, func(w *io.IWriter) error {
			_, err := w.WriteHEX("68AAAAAAAAAAAA68910833333333343333337E16")
			return err
		})
	}).DoneAll()
}
