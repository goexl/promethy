package param

import (
	"github.com/goexl/env"
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
		Labels: map[string]string{
			"namespace": env.Get("NAMESPACE"), // 命名空间
			"pod":       env.Get("POD"),       // 容器
			"node":      env.Get("NODE"),      // 节点
			"ip":        env.Get("IP"),        // 地址
		},
	}
}
