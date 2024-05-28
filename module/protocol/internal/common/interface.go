package common

import "github.com/injoyai/goutil/g"

type Decoder interface {
	Decode(bs []byte) (g.Map, error)
}
