package listen

import (
	"github.com/injoyai/io"
	"github.com/injoyai/io/listen"
)

func NewTCP(port int, option ...io.OptionServer) (*io.Server, error) {
	return listen.NewTCPServer(port, option...)
}
