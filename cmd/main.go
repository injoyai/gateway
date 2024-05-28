package main

import (
	"github.com/injoyai/gateway/api"
	"github.com/injoyai/gateway/common"
	"github.com/injoyai/gateway/logic"
)

func main() {
	common.Init()
	logic.Init()
	api.Init()
}
