package model

import "github.com/injoyai/gateway/internal/data/nature"

type Model struct {
	Natures nature.Natures `json:"natures"`
}
