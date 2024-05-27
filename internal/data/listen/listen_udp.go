package listen

import (
	"github.com/injoyai/io"
	"github.com/injoyai/io/listen"
)

func NewUDP(port int, option ...io.OptionServer) (*io.Server, error) {
	return listen.NewUDPServer(port, option...)
}
