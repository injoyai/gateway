package listen

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewMQTTServer(sub *MQTTSubscribe, options ...func(*mqtt.ClientOptions)) (*MQTTServer, error) {
	op := mqtt.NewClientOptions()
	for _, f := range options {
		f(op)
	}
	c := mqtt.NewClient(op)
	token := c.Connect()
	token.Wait()
	if err := token.Error(); err != nil {
		return nil, err
	}
	ms := &MQTTServer{
		Client:  c,
		options: op,
		sub:     sub,
		message: make(chan mqtt.Message),
	}

	token = c.Subscribe(sub.Topic, sub.Qos, func(client mqtt.Client, message mqtt.Message) {
		ms.message <- message
	})
	token.Wait()
	if token.Error() != nil {
		c.Disconnect(0)
		return nil, token.Error()
	}

	return ms, nil
}

type MQTTServer struct {
	mqtt.Client
	options *mqtt.ClientOptions
	sub     *MQTTSubscribe
	message chan mqtt.Message
	GetQos  func(topic string) (qos byte, retained bool)
}

func (this *MQTTServer) Publish(topic string, data interface{}) error {
	var qos byte
	var retained bool
	if this.GetQos != nil {
		qos, retained = this.GetQos(topic)
	}
	token := this.Client.Publish(topic, qos, retained, data)
	token.Wait()
	return token.Error()
}

type MQTTSubscribe struct {
	Topic string
	Qos   uint8
}
