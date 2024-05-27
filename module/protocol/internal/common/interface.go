package common

type Decoder interface {
	Decode(bs []byte) ([]byte, error)
}
