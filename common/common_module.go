package common

import (
	"fmt"
	"github.com/injoyai/conv/cfg"
	"github.com/injoyai/conv/codec"
	"github.com/injoyai/gateway/module/listen"
	"github.com/injoyai/gateway/module/protocol"
	"github.com/injoyai/gateway/module/push"
	"github.com/injoyai/gateway/module/register"
	"github.com/injoyai/goutil/database/redis"
	"github.com/injoyai/goutil/database/xorms"
	"github.com/injoyai/goutil/g"
)

var (
	DB       *xorms.Engine
	Redis    *redis.Client
	Register = register.New()
)

var (
	Listener    = listen.New()
	Protocol    *protocol.Manager
	Pusher      push.Interface
	Distributor interface{}
)

func Init() {

	//读取配置文件
	cfg.Default = cfg.New("./config/config.yaml", codec.Yaml)
	var err error

	//初始化数据库
	DB = xorms.New(&xorms.Config{
		Type: cfg.GetString("database.type"),
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			cfg.GetString("database.username"),
			cfg.GetString("database.password"),
			cfg.GetString("database.host", "localhost"),
			cfg.GetInt("database.port", 3306),
			cfg.GetString("database.database"),
		),
		FieldSync:   cfg.GetBool("database.fieldSync"),
		TablePrefix: cfg.GetString("database.tablePrefix"),
	})
	g.PanicErr(DB.Err())

	//初始化redis
	Redis = redis.New(
		cfg.GetString("addr"),
		cfg.GetString("password"),
		cfg.GetInt("db"),
	)
	g.PanicErr(Redis.Ping())

	//初始化协议
	Protocol, err = protocol.New(&protocol.Config{
		DslDir: cfg.GetString("protocol.dslDir"),
	})
	g.PanicErr(err)

	//初始化推送
	Pusher = push.New(&push.Config{
		Type:  cfg.GetString("push.type"),
		Host:  cfg.GetString("push.host"),
		Port:  cfg.GetInt("push.port"),
		Param: cfg.GetMap("push"),
	})

}
