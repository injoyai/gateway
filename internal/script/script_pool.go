package script

import (
	"github.com/injoyai/conv"
	"github.com/injoyai/goutil/script/js"
	"sync"
)

func NewPool(op ...Option) *Pool {
	p := &Pool{
		options: op,
	}
	p.pool = sync.Pool{
		New: p.new,
	}
	return p
}

type Option func(c *js.Client)

type Pool struct {
	options []Option
	pool    sync.Pool
	count   int
}

func (this *Pool) new() any {
	this.count++
	c := js.New()
	for _, v := range this.options {
		v(c)
	}
	return c
}

func (this *Pool) Exec(text string) (*conv.Var, error) {
	a := this.pool.Get()
	v, err := a.(*js.Client).Exec(text)
	this.pool.Put(a)
	return v, err
}
