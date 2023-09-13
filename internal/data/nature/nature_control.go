package nature

type Control struct {
	Type  ControlType            `json:"type"`  //控制类型(可读可写,根据后端数据,部分长按可放大) 开关,枚举,滑条,旋转按钮,仪表盘,窗口...
	Param map[string]interface{} `json:"param"` //类型对应参数
	Style string                 `json:"style"` //样式 ""默认样式 "xxx"xxx样式
}

type ControlType string

const (
	ControlSwitch    ControlType = "switch"    //开关
	ControlEnum      ControlType = "enum"      //枚举
	ControlSlider    ControlType = "slider"    //滑条
	ControlRotating  ControlType = "rotating"  //旋转按钮
	ControlDashboard ControlType = "dashboard" //仪表盘
	ControlWindow    ControlType = "window"    //窗口
	ControlPlayer    ControlType = "player"    //播放器
	ControlCustom    ControlType = "custom"    //自定义
)
