package protocol

import (
	"github.com/injoyai/gateway/module/protocol/built"
	v1 "github.com/injoyai/gateway/module/protocol/dsl/v1"
	"github.com/injoyai/gateway/module/protocol/internal/common"
	"github.com/injoyai/goutil/g"
	"github.com/injoyai/goutil/oss"
)

type Config struct {
	DslDir string //dsl的配置目录
}

func New(cfg *Config) (*Manager, error) {

	m := &Manager{
		built: built.All, //加载内置协议
		dsl:   map[string]common.Decoder{},
		cfg:   cfg,
	}

	//加载dsl
	err := m.LoadingDSL()

	return m, err
}

type Manager struct {
	built map[string]common.Decoder
	dsl   map[string]common.Decoder
	cfg   *Config
}

// LoadingDSL 加载dsl
func (this *Manager) LoadingDSL() error {
	oss.NewDir(this.cfg.DslDir)
	dsl, err := v1.New(this.cfg.DslDir)
	if err != nil {
		return err
	}
	this.dsl = dsl
	return nil
}

// Register 注册协议解析,代码层面,
// 当数量很大的时候,脚本得速度不能满足时,可以从代码层面进行性能的提升
func (this *Manager) Register(name string, codec common.Decoder) {
	this.built[name] = codec
}

func (this *Manager) Decode(name string, bs []byte) (bool, g.Map, error) {
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
	return false, g.Map{"bytes": bs}, nil
}
