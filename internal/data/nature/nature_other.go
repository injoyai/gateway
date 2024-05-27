package nature

/*
type Device struct {
	ID   int64  // ID
	Name string // 名称
	Memo string // 备注
}

// Codec 编解码
type Codec struct {
	//怎么处理字节
	//字节序

}

func (this *Codec) Decode() interface{} {
	return nil
}

type Transport struct {
	Type           string                 `json:"type"`           //类型,串口,tcp,udp,http...
	Address        string                 `json:"address"`        //通讯地址
	ConnectTimeout int64                  `json:"connectTimeout"` //连接超时
	Param          map[string]interface{} `json:"param"`          //通讯参数

	ProtocolClass  string `json:"protocolClass"`  //协议分类 标准协议
	ProtocolVendor string `json:"ProtocolVendor"` //协议厂商 Modbus
	Protocol       string `json:"protocol"`       //协议类型 Modbus TCP

	Interval int64 //采集间隔

}

// Gather 数据采集配置
type Gather struct {
	Slave       string
	Address     string
	AddressNum  int
	AddressType string
}

type Value struct {
	ID         int64  // ID
	DeviceID   int64  //设备ID
	NatureID   int64  //属性ID
	Value      string //实时值
	ValueDate  int64  //时间
	Value1     string //上1次值
	ValueDate1 int64  //上1次时间
	Value2     string //上2次值
	ValueDate2 int64  //上2次时间
	Value3     string //上3次值
	ValueDate3 int64  //上3次时间
	Value4     string //上4次值
	ValueDate4 int64  //上4次时间
	Value5     string //上5次值
	ValueDate5 int64  //上5次时间
}

func (this *Value) Chart() *Chart {
	return &Chart{
		Name:  "",
		Unit:  "",
		Date:  []int64{this.ValueDate, this.ValueDate1, this.ValueDate2, this.ValueDate3, this.ValueDate4, this.ValueDate5},
		Value: []string{this.Value, this.Value1, this.Value2, this.Value3, this.Value4, this.Value5},
	}
}



//// Model 型号
//type Mode1l struct {
//	ID   int64  // ID
//	Name string // 名称 QL-GW21
//	Memo string // 备注 边缘网关
//}
//
//// Nature 一个型号有多个属性
//type Nature1 struct {
//	ID      int64 // ID
//	ModelID int64 // 型号ID
//	InDate  int64 // 入库时间
//
//	Key          string // 键
//	Name         string // 名称
//	Memo         string // 备注
//	ValueType    string // 值类型 string int float bool script
//	DefaultValue string // 默认值
//	Readable     bool   // 是否可读
//	Writable     bool   // 是否可写
//	Unit         string // 单位 ℃
//}

// Method 一个型号有多个方法,例如数据偏移,放大缩小
type Method struct {
	NatureID int64  // 属性ID
	Script   string // 脚本
}

// Event 一个属性有多个事件
type Event struct {
	Message string
}

// Alarm 一个属性有多个告警,属于事件
type Alarm struct {
}

type Chart struct {
	Name  string
	Unit  string
	Date  []int64
	Value []string
}

*/
