package listen

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Deal interface {
	DealMessage(c interface{}, msg []byte) (ack bool)
}

func NewMQTTClient(sub *MQTTSubscribe, deal Deal, options ...func(*mqtt.ClientOptions)) (*MQTTClient, error) {
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
	ms := &MQTTClient{
		Client:  c,
		options: op,
		sub:     sub,
	}

	token = c.Subscribe(sub.Topic, sub.Qos, func(client mqtt.Client, message mqtt.Message) {
		if deal.DealMessage(ms, message.Payload()) {
			message.Ack()
		}
	})
	token.Wait()
	if token.Error() != nil {
		c.Disconnect(0)
		return nil, token.Error()
	}

	return ms, nil
}

type MQTTClient struct {
	mqtt.Client
	options *mqtt.ClientOptions
	sub     *MQTTSubscribe
	GetQos  func(topic string) (qos byte, retained bool)
}

func (this *MQTTClient) Publish(topic string, data interface{}) error {
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
