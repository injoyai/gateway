package client

import (
	"fmt"
	"github.com/injoyai/conv"
	"github.com/injoyai/logs"
	"sync"
)

func New() *Manager {
	return &Manager{
		m: make(map[string]Client),
	}
}

type Manager struct {
	m  map[string]Client
	mu sync.RWMutex
}

func (this *Manager) Dial(cfg *Config) (Client, error) {
	key := cfg.GetKey()
	logs.Debug(key)
	this.mu.RLock()
	c, ok := this.m[key]
	this.mu.RUnlock()
	if !ok {
		var err error
		c, err = cfg.Dial()
		if err != nil {
			return nil, err
		}
		this.mu.Lock()
		this.m[key] = c
		this.mu.Unlock()
	}
	return c, nil

}

func (this *Manager) Publish(key string, topic string, data interface{}) (err error) {

	defer func() {
		logs.Writef("[%s.%s] 值: %s, 结果: %s\n", key, topic, conv.String(data), conv.New(err).String("成功"))
	}()

	c, ok := this.m[key]
	if !ok {
		return fmt.Errorf("客户端不存在: %s", key)
	}

	err = c.Publish(topic, data)
	if err != nil {
		return err
	}

	return nil
}
