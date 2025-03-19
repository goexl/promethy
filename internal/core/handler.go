package core

import (
	"net/http"

	"github.com/goexl/promethy/internal/param"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Handler 控制器
type Handler struct {
	params *param.Promethy
}

func NewHandler(params *param.Promethy) *Handler {
	return &Handler{
		params: params,
	}
}

func (h *Handler) Handle() (handler http.Handler, err error) {
	register := prometheus.WrapRegistererWith(h.params.Parse(), h.params.Registry)
	if rpe := register.Register(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{})); nil != rpe {
		err = rpe
	} else if rge := register.Register(collectors.NewGoCollector()); nil != rge {
		err = rge
	} else {
		handler = promhttp.HandlerFor(h.params.Registry, promhttp.HandlerOpts{Registry: register})
	}

	return
}
