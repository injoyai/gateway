package push

import v1 "github.com/injoyai/gateway/module/push/internal/mq/v1"

/*
DefaultInternalMQ
内部消息队列
推送数据到内存
*/
var DefaultInternalMQ = v1.New()
