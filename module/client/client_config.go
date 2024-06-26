package client

import (
	"fmt"
	"github.com/injoyai/conv"
	"github.com/injoyai/goutil/g"
	"github.com/injoyai/io"
	"github.com/injoyai/io/buf"
)

type Config struct {
	Type      string                 //监听/推送类型
	Addr      string                 //监听/推送地址
	Subscribe string                 //监听/订阅 端口/主题
	Publish   string                 //发布/推送 topic
	Param     map[string]interface{} //监听/推送参数
	Frame     *buf.Frame             //分包配置
}

func (this *Config) GetListenPort() int {
	return conv.Int(this.Subscribe)
}

func (this *Config) init() {
	if this.Param == nil {
		this.Param = make(map[string]interface{})
	}
	if this.Param["clientID"] == nil {
		this.Param["clientID"] = g.RandString(8)
	}
}

func (this *Config) GetKey() string {
	switch this.Type {
	case io.MQTT:
		this.init()
		return fmt.Sprintf("%s#%s#%s", this.Type, this.Addr, this.Param["clientID"])
	case io.Serial:
		return fmt.Sprintf("%s#%s", this.Type, this.Subscribe)
	case io.TCP, io.UDP, io.HTTP, io.Websocket:
		return fmt.Sprintf("net#%s", this.Subscribe)
	default:
		return fmt.Sprintf("%s#%s#%s", this.Type, this.Addr, this.Subscribe)
	}

}

func (this *Config) Dial() (Client, error) {
	if this.Frame == nil {
		this.Frame = &buf.Frame{}
	}

	switch this.Type {
	//case "rabbitmq":
	//	return NewRabbitMQ()
	//case "rocketmq":
	//	return NewRocketMQ()
	//case "kafka":
	//	return NewKafka()
	//case "http":
	//	return NewHTTP()
	case io.MQTT:
		this.init()
		return NewMQTTClient(this)

	case io.TCP + "client":
		return NewIOClient(this, this.Frame)

	default:
		return NewIOServer(this, this.Frame)
	}
}
