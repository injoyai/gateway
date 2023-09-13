package listen

import (
	"github.com/injoyai/io"
	"github.com/injoyai/io/dial"
)

func NewTCP(port int, option ...io.OptionServer) (*io.Server, error) {
	return dial.NewTCPServer(port, option...)
}
