package log

import "github.com/injoyai/io"

type Level int

const (
	LevelInfo Level = 1 + iota
	LevelDebug
	LevelError
)

type Log interface {
	SetLevel(Level)
	SetWriter(Level, ...io.Writer)
	Info(...interface{})
	Debug(...interface{})
	Err(...interface{})
}
