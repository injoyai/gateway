package v1

import "github.com/injoyai/goutil/other/trunk"

func New() *MQ {
	return &MQ{trunk.New()}
}

type MQ struct {
	*trunk.Entity
}

func (this *MQ) Publish(topic string, data interface{}) error {
	this.Entity.Publish(topic, data)
	return nil
}
