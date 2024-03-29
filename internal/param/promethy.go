package param

import (
	"github.com/goexl/log"
	"github.com/goexl/valuer"
	"github.com/prometheus/client_golang/prometheus"
)

type Promethy struct {
	Logger   log.Logger
	Registry *prometheus.Registry
	Parser   *valuer.Parser
	Labels   map[string]string
}

func NewParams() *Promethy {
	return &Promethy{
		Logger:   log.New().Apply(),
		Registry: prometheus.NewRegistry(),
		Parser:   valuer.New().Build(),
		Labels:   make(map[string]string),
	}
}
