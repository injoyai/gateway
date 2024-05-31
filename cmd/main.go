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
	"path/filepath"
)

var Debug = "true"

func init() {
	//if !conv.Bool(Debug) {
	logs.SetFormatterWithTime()
	//}
	logCfg := cfg.New("./config/log.json")
	level := logs.LevelAll
	switch logCfg.GetString("level") {
	case "all":
		level = logs.LevelAll
	case "trace":
		level = logs.LevelTrace
	case "write":
		level = logs.LevelWrite
	case "read":
		level = logs.LevelRead
	case "info":
		level = logs.LevelInfo
	case "debug":
		level = logs.LevelDebug
	case "warn":
		level = logs.LevelWarn
	case "error":
		level = logs.LevelError
	case "none":
		level = logs.LevelNone
	default:
		logs.SetLevel(logs.LevelAll)
	}
	logs.SetLevel(level)
	logs.SetShowColor(logCfg.GetBool("color", true))
	logs.SetSaveTime(logCfg.GetDuration("saveTime"))
	logs.DefaultErr.SetWriter(logs.Stdout)
	for k, v := range logCfg.GetGMap("output") {
		e := logs.DefaultTrace
		dir, filename := filepath.Split(conv.String(v))
		switch k {
		case "trace":
			e = logs.DefaultTrace
		case "write":
			e = logs.DefaultTrace
		case "read":
			e = logs.DefaultTrace
		case "info":
			e = logs.DefaultTrace
		case "debug":
			e = logs.DefaultTrace
		case "warn":
			e = logs.DefaultTrace
		case "error":
			e = logs.DefaultTrace
		default:
			continue
		}
		e.WriteToFile(dir, filename)
	}
}

func main() {

	//读取配置文件
	cfg.Default = cfg.New("./config/config.yaml", codec.Yaml)

	//初始化实例
	common.Init()

	//加载默认link
	links := []*distribute.Link(nil)
	for _, v := range cfg.Default.GetInterfaces("link") {
		m := conv.NewMap(v)
		links = append(links, &distribute.Link{
			Name:  m.GetString("name"),
			Model: m.GetString("model"),
			Subscribe: func() []*client.Config {
				list := []*client.Config(nil)
				for _, subscribe := range m.GetInterfaces("subscribe") {
					m2 := conv.NewMap(subscribe)
					list = append(list, &client.Config{
						Type:      m2.GetString("type"),
						Addr:      m2.GetString("addr"),
						Subscribe: m2.GetString("subscribe"),
						Param:     m2.GetGMap(""),
					})
				}
				return list
			}(),
			Protocol: m.GetString("protocol"),
			Publish: func() []*client.Config {
				list := []*client.Config(nil)
				for _, publish := range m.GetInterfaces("publish") {
					m2 := conv.NewMap(publish)
					list = append(list, &client.Config{
						Type:    m2.GetString("type"),
						Addr:    m2.GetString("addr"),
						Publish: m2.GetString("publish"),
						Param:   m2.GetGMap(""),
					})
				}
				return list
			}(),
		})
	}
	err := logic.Init(links)
	g.PanicErr(err)

	api.Init()
	oss.Wait()
}
