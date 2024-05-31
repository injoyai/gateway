package client

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/injoyai/conv"
	"sync"
	"time"
)

var (
	cacheMQTT = make(map[string]*MQTTClient)
	cacheMu   sync.RWMutex
)

func NewMQTTClient(cfg *Config) (*MQTTSubscribe, error) {

	clientKey := cfg.GetKey()
	cacheMu.RLock()
	client, ok := cacheMQTT[clientKey]
	cacheMu.RUnlock()
	m := conv.NewMap(cfg.Param)
	if !ok {
		c, err := newClient(cfg)
		if err != nil {
			return nil, err
		}
		client = &MQTTClient{
			Client:    c,
			key:       clientKey,
			subscribe: make(map[string]*MQTTSubscribe),
		}
		cacheMQTT[cfg.GetKey()] = client
	}

	client.subscribeMu.RLock()
	_, ok = client.subscribe[cfg.Subscribe]
	client.subscribeMu.RUnlock()
	if ok {
		return nil, fmt.Errorf("主题[%s]已被订阅", cfg.Subscribe)
	}

	ms := &MQTTSubscribe{
		MQTTClient:      client,
		publishQos:      m.GetUint8("publishQos"),
		publishRetained: m.GetBool("publishRetained"),
		subscribeTopic:  cfg.Subscribe,
		subscribeQos:    m.GetUint8("subscribeQos"),
		onMessage:       nil,
	}
	if len(cfg.Subscribe) > 0 {
		token := client.Client.Subscribe(cfg.Subscribe, m.GetUint8("subscribeQos"), func(client mqtt.Client, message mqtt.Message) {
			if ms.onMessage != nil && ms.onMessage(ms, message.Payload()) {
				message.Ack()
			}
		})
		token.Wait()
		if token.Error() != nil {
			ms.Close()
			return nil, token.Error()
		}
	}

	client.subscribeMu.Lock()
	client.subscribe[cfg.Subscribe] = ms
	client.subscribeMu.Unlock()
	return ms, nil

}

func newClient(cfg *Config) (mqtt.Client, error) {
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
	return c, nil
}

type MQTTClient struct {
	mqtt.Client
	key         string
	subscribe   map[string]*MQTTSubscribe
	subscribeMu sync.RWMutex
}

type MQTTSubscribe struct {
	*MQTTClient
	publishQos      byte
	publishRetained bool
	subscribeTopic  string
	subscribeQos    byte
	onMessage       MessageFunc
}

func (this *MQTTSubscribe) OnMessage(f MessageFunc) {
	this.onMessage = f
}

func (this *MQTTSubscribe) Publish(topic string, data interface{}) error {
	token := this.Client.Publish(topic, this.publishQos, this.publishRetained, conv.Bytes(data))
	token.Wait()
	return token.Error()
}

func (this *MQTTSubscribe) Close() error {
	delete(this.MQTTClient.subscribe, this.subscribeTopic)
	if len(this.MQTTClient.subscribe) == 0 {
		this.MQTTClient.Client.Disconnect(0)
		cacheMu.Lock()
		delete(cacheMQTT, this.MQTTClient.key)
		cacheMu.Unlock()
	}
	return nil
}
