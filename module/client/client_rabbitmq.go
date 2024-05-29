package client

import v1 "github.com/injoyai/gateway/module/client/internal/mq/v1"

func NewRabbitMQ() (*v1.MQ, error) {
	return v1.New(), nil
}

func NewRocketMQ() (*v1.MQ, error) {
	return v1.New(), nil
}

func NewKafka() (*v1.MQ, error) {
	return v1.New(), nil
}
