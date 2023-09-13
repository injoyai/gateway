package nature

type Model struct {
	Name     string `json:"name"`     //名称
	Memo     string `json:"memo"`     //备注
	Key      string `json:"key"`      //数据标识
	Type     string `json:"type"`     //数据类型
	Unit     string `json:"unit"`     //单位
	Readable bool   `json:"readable"` //是否可写
	Writable bool   `json:"writable"` //是否可读
}

func NewModel(n Nature) *Model {
	return &Model{
		Name:     n.Name(),
		Memo:     n.Memo(),
		Key:      n.Key(),
		Type:     n.Type(),
		Unit:     n.Unit(),
		Readable: n.Readable(),
		Writable: n.Writable(),
	}
}
