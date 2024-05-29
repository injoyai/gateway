package client

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/injoyai/conv"
	"time"
)

func NewMQTTClient(cfg *Config) (*MQTTClient, error) {
	m := conv.NewMap(cfg.Param)
	op := mqtt.NewClientOptions()
	op.AddBroker(cfg.Addr)
	op.SetClientID(m.GetString("clientID"))
	op.SetUsername(m.GetString("username"))
	op.SetPassword(m.GetString("password"))
	op.SetConnectTimeout(time.Duration(m.GetInt64("connectTimeout", 30*1e9)))
	op.SetKeepAlive(time.Duration(m.GetInt64("keepAlive")))
	op.SetCleanSession(m.GetBool("cleanSession"))
	op.SetCleanSession(true)
	op.SetAutoReconnect(true) //自动重连

	c := mqtt.NewClient(op)
	token := c.Connect()
	token.Wait()
	if err := token.Error(); err != nil {
		return nil, err
	}
	ms := &MQTTClient{
		Client:  c,
		options: op,
	}

	if len(cfg.Subscribe) > 0 {
		token = c.Subscribe(cfg.Subscribe, m.GetUint8("qos"), func(client mqtt.Client, message mqtt.Message) {
			if ms.onMessage != nil && ms.onMessage(ms, message.Payload()) {
				message.Ack()
			}
		})
		token.Wait()
		if token.Error() != nil {
			c.Disconnect(0)
			return nil, token.Error()
		}
	}

	return ms, nil
}

type MQTTClient struct {
	mqtt.Client
	options   *mqtt.ClientOptions
	GetQos    func(topic string) (qos byte, retained bool)
	onMessage MessageFunc
}

func (this *MQTTClient) OnMessage(f MessageFunc) {
	this.onMessage = f
}

func (this *MQTTClient) Publish(topic string, data interface{}) error {
	var qos byte
	var retained bool
	if this.GetQos != nil {
		qos, retained = this.GetQos(topic)
	}
	token := this.Client.Publish(topic, qos, retained, conv.Bytes(data))
	token.Wait()
	return token.Error()
}

type MQTTSubscribe struct {
	Topic string
	Qos   uint8
}
