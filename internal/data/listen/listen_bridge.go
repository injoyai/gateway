package listen

import (
	"github.com/injoyai/base/maps"
	"github.com/injoyai/io"
	"github.com/injoyai/io/listen"
	json "github.com/json-iterator/go"
)

func NewBridge(port int, option ...io.OptionServer) (*BridgeServer, error) {
	s, err := listen.NewTCPServer(port, func(s *io.Server) {
		s.SetOptions(option...)
		s.SetDealFunc(func(c *io.Client, msg io.Message) {

			req := new(Request)
			json.Unmarshal(msg.Bytes(), req)
			switch req.Type {
			case RequestTypeBridge:
				data := new(BridgeRequest)
				json.Unmarshal(msg.Bytes(), data)

			}

		})
	})
	return &BridgeServer{
		Server: s,
		bridge: maps.NewSafe(),
	}, err
}

type BridgeServer struct {
	*io.Server
	bridge *maps.Safe
}

type Request struct {
	Type string      `json:"type"`
	UUID string      `json:"uuid"`
	Data interface{} `json:"data"`
}

type Response struct {
	Code int    `json:"code"`
	UUID string `json:"uuid"`
	Msg  string `json:"msg"`
}

type BridgeRequest struct {
	Type string `json:"type"`
	Data struct {
		Bridge string `json:"bridge"` //桥接对象
	} `json:"data"`
}

const (
	RequestTypeBridge = "bridge"
	RequestTypeWrite  = "write"
)
