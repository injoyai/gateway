package boot

import "github.com/injoyai/gateway/internal/common"

func Init() {
	common.Init() //优先级2

	initScript()
}
