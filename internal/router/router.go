package router

import (
	"github.com/injoyai/gateway/internal/common"
	"github.com/injoyai/gateway/internal/router/api"
)

func Init() {

	apiGroup := common.HTTPServer.Group("/api")

	apiGroup.GET("/listen/list", api.GetListenList)
	apiGroup.POST("/listen/list", api.PostListen)
	apiGroup.DELETE("/listen/list", api.DelListen)

}
