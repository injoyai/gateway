package v1

import "github.com/injoyai/goutil/other/trunk"

func New() *MQ {
	return &MQ{trunk.NewGroup()}
}

type MQ struct {
	*trunk.Group
}

func (this *MQ) Publish(topic string, data interface{}) error {
	this.Group.Publish(topic, data)
	return nil
}
