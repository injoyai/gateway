package listen

import (
	"github.com/injoyai/io"
	"github.com/injoyai/io/dial"
)

func NewSerial(cfg *dial.SerialConfig, option ...io.OptionClient) (*io.Client, error) {
	return dial.NewSerial(cfg, option...)
}
