package nature

import (
	"errors"
	"github.com/injoyai/conv"
)

type ValueType string

func (this ValueType) Value(value interface{}) interface{} {
	switch this {
	case Float:
		return conv.Float64(value)
	case Int:
		return conv.Int64(value)
	case Bool:
		return conv.Bool(value)
	case String:
		return conv.String(value)
	}
	return value
}

const (
	String ValueType = "STRING"
	Bool   ValueType = "BOOL"
	Int    ValueType = "INT"
	Float  ValueType = "FLOAT"
)

var (
	ErrNatureReadInvalid  = errors.New("属性不可读")
	ErrNatureWriteInvalid = errors.New("属性不可写")
)

func GetValue(valueType ValueType, value interface{}) interface{} {
	return valueType.Value(value)
}
