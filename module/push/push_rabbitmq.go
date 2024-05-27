package push

import v1 "github.com/injoyai/gateway/module/push/internal/mq/v1"

func NewRabbitMQ() *v1.MQ {
	return v1.New()
}
