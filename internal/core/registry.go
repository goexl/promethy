package core

import (
	"github.com/goexl/promethy/internal/param"
	"github.com/prometheus/client_golang/prometheus"
)

// Registry 注册表
type Registry struct {
	params *param.Promethy
}

func NewRegistry(params *param.Promethy) *Registry {
	return &Registry{
		params: params,
	}
}

func (r *Registry) Register(required prometheus.Collector, optionals ...prometheus.Collector) (err error) {
	register := prometheus.WrapRegistererWith(r.params.Parse(), r.params.Registry)
	for _, collector := range append([]prometheus.Collector{required}, optionals...) {
		err = register.Register(collector)
		if nil != err {
			break
		}
	}

	return
}
