package promethy

import (
	"github.com/goexl/simaqian"
	"github.com/goexl/valuer"
	"github.com/prometheus/client_golang/prometheus"
)

type params struct {
	logger   simaqian.Logger
	registry *prometheus.Registry
	parser   *valuer.Parser
	labels   map[string]string
}

func newParams() *params {
	return &params{
		logger:   simaqian.Default(),
		registry: prometheus.NewRegistry(),
		parser:   valuer.New().Build(),
		labels:   make(map[string]string),
	}
}
