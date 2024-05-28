package built

import "github.com/injoyai/gateway/module/protocol/internal/common"

var All = map[string]common.Decoder{
	"dlt645":     &Dlt645{},
	"modbus tcp": &ModbusTCP{},
}
