package status

type Interface interface {
	// Running 运行情况
	Status() (*Status, error)

	// Enable 启用禁用
	Enable(b ...bool) error

	// Disable 禁用
	Disable() error
}

type Status struct {
	Running   bool   `json:"running"`   //是否在运行(在线/离线)
	Error     string `json:"error"`     //错误信息
	StartTime int64  `json:"startTime"` //开始时间
}
