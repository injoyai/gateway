package main

import (
	"github.com/injoyai/conv"
	"github.com/injoyai/conv/cfg"
	"github.com/injoyai/conv/codec"
	"github.com/injoyai/gateway/api"
	"github.com/injoyai/gateway/common"
	"github.com/injoyai/gateway/logic"
	"github.com/injoyai/gateway/module/client"
	"github.com/injoyai/gateway/module/distribute"
	"github.com/injoyai/goutil/g"
	"github.com/injoyai/goutil/oss"
	"github.com/injoyai/logs"
)

func main() {
	logs.SetFormatterWithTime()
	//读取配置文件
	cfg.Default = cfg.New("./config/config.yaml", codec.Yaml)

	//初始化实例
	common.Init()

	links := []*distribute.Link(nil)
	for _, v := range cfg.Default.GetInterfaces("link") {
		m := conv.NewMap(v)
		logs.Debug(v)
		links = append(links, &distribute.Link{
			Model: m.GetString("model"),
			Subscribe: &client.Config{
				Type:      m.GetString("subscribe.type"),
				Addr:      m.GetString("subscribe.addr"),
				Subscribe: m.GetString("subscribe.subscribe"),
				Param:     m.GetGMap("subscribe"),
			},
			//Frame:    v.(*distribute.Link).Frame,
			Protocol: m.GetString("protocol"),
			Publish: &client.Config{
				Type:      m.GetString("publish.type"),
				Addr:      m.GetString("publish.addr"),
				Subscribe: m.GetString("publish.subscribe"),
				Param:     m.GetGMap("publish"),
			},
		})
	}
	err := logic.Init(links)
	g.PanicErr(err)

	api.Init()
	oss.Wait()
}
