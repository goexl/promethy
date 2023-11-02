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
	labels := make(map[string]string)
	for key, value := range h.params.Labels {
		labels[h.params.Parser.Parse(key)] = h.params.Parser.Parse(value)
	}
	register := prometheus.WrapRegistererWith(labels, h.params.Registry)
	if re := register.Register(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{})); nil != re {
		err = re
	} else if ge := register.Register(collectors.NewGoCollector()); nil != ge {
		err = ge
	} else {
		handler = promhttp.HandlerFor(h.params.Registry, promhttp.HandlerOpts{Registry: register})
	}

	return
}
