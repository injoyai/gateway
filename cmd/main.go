package main

import (
	"github.com/injoyai/gateway/internal/boot"
	"github.com/injoyai/gateway/internal/common"
)

func main() {

	//初始化
	boot.Init()

	//运行http服务
	common.RunHTTPServer()
}
