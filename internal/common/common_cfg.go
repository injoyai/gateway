package common

import (
	"github.com/injoyai/conv/cfg"
)

var Cfg *cfg.Entity

func initCfg() {
	Cfg = cfg.New(DefaultConfigPath)
}
