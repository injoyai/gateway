package common

import (
	"fmt"
	"github.com/injoyai/conv/cfg"
	"github.com/injoyai/gateway/module/client"
	"github.com/injoyai/gateway/module/listen"
	"github.com/injoyai/gateway/module/protocol"
	"github.com/injoyai/gateway/module/register"
	"github.com/injoyai/goutil/database/mysql"
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
	Client      = client.New()
	Distributor interface{}
)

func Init() {

	var err error

	//初始化数据库
	DB = mysql.NewXorm(&xorms.Option{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
			cfg.GetString("database.username"),
			cfg.GetString("database.password"),
			cfg.GetString("database.host"),
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

	//初始化客户端管理
	//_, err = Client.Dial(&client.Config{
	//	Type:  cfg.GetString("push.type"),
	//	Addr:  cfg.GetString("push.addr"),
	//	Port:  cfg.GetString("push.port"),
	//	Param: cfg.GetMap("push"),
	//})
	g.PanicErr(err)

}
