package listen

import "time"

type Message struct {
	ListenType string    `json:"listenType"` //监听类型
	Port       string    `json:"port"`       //端口
	DataNo     int64     `json:"dataNo"`     //数据编号
	Time       time.Time `json:"time"`       //时间
	Payload    []byte    `json:"payload"`    //数据内容
}
