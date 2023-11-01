package listen

import (
	"github.com/injoyai/io"
	"github.com/injoyai/io/dial"
)

func NewUDP(port int, option ...io.OptionServer) (*io.Server, error) {
	return dial.NewUDPServer(port, option...)
}
