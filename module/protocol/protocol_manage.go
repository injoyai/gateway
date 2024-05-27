package protocol

import (
	"github.com/injoyai/gateway/module/protocol/built"
	v1 "github.com/injoyai/gateway/module/protocol/dsl/v1"
	"github.com/injoyai/gateway/module/protocol/internal/common"
)

type Config struct {
	DslDir string //dsl的配置目录
}

func New(cfg *Config) (*Manager, error) {

	m := &Manager{
		built: map[string]common.Decoder{},
		dsl:   map[string]common.Decoder{},
	}

	//加载内置协议
	for k, v := range built.All {
		m.Register(k, v)
	}

	//加载dsl
	var err error
	m.dsl, err = v1.New(cfg.DslDir)

	return m, err
}

type Manager struct {
	built map[string]common.Decoder
	dsl   map[string]common.Decoder
}

// Register 注册协议解析,代码层面,
// 当数量很大的时候,脚本得速度不能满足时,可以从代码层面进行性能的提升
func (this *Manager) Register(name string, codec common.Decoder) {
	this.built[name] = codec
}

func (this *Manager) Decode(name string, bs []byte) (bool, []byte, error) {
	//尝试在内置协议协议中查找
	if c, ok := this.built[name]; ok {
		bs, err := c.Decode(bs)
		return true, bs, err
	}

	//尝试在dsl中查找
	if c, ok := this.dsl[name]; ok {
		bs, err := c.Decode(bs)
		return true, bs, err
	}

	//不解析
	return false, bs, nil
}
