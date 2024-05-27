package push

import v1 "github.com/injoyai/gateway/module/push/internal/mq/v1"

type Interface interface {
	Publish(topic string, data interface{}) error
}

type Config struct {
	Type  string
	Host  string
	Port  int
	Param map[string]interface{}
}

func New(cfg *Config) Interface {
	switch cfg.Type {
	case "rabbitmq":
		return NewRabbitMQ()
	case "rocketmq":
		return NewRocketMQ()
	case "kafka":
		return NewKafka()
	case "http":
		return NewHTTP()
	default:
		return v1.New()
	}
}
