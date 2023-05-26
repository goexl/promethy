package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Registry 注册表
type Registry struct {
	params *params
}

func newRegistry(params *params) *Registry {
	return &Registry{
		params: params,
	}
}

func (r *Registry) Register(collectors ...prometheus.Collector) (err error) {
	for _, collector := range collectors {
		err = r.params.registry.Register(collector)
		if nil != err {
			break
		}
	}

	return
}
