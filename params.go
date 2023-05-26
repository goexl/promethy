package prometheus

import (
	"github.com/goexl/simaqian"
	"github.com/prometheus/client_golang/prometheus"
)

type params struct {
	logger   simaqian.Logger
	registry *prometheus.Registry
	labels   map[string]string
}

func newParams() *params {
	return &params{
		logger:   simaqian.Default(),
		registry: prometheus.NewRegistry(),
		labels:   make(map[string]string),
	}
}
