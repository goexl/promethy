package core

import (
	"github.com/goexl/promethy/internal/param"
)

type Promethy struct {
	params *param.Promethy
}

func NewPromethy(params *param.Promethy) *Promethy {
	return &Promethy{
		params: params,
	}
}

func (p *Promethy) Register() *Registry {
	return NewRegistry(p.params)
}

func (p *Promethy) Handler() *Handler {
	return NewHandler(p.params)
}
