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

func (r *Registry) Register(collectors ...prometheus.Collector) (err error) {
	for _, collector := range collectors {
		err = r.params.Registry.Register(collector)
		if nil != err {
			break
		}
	}

	return
}
