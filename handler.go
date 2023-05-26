package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Handler 控制器
type Handler struct {
	params *params
}

func newHandler(params *params) *Handler {
	return &Handler{
		params: params,
	}
}

func (h *Handler) Handle() (handler http.Handler, err error) {
	register := prometheus.WrapRegistererWith(h.params.labels, h.params.registry)
	register.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	register.MustRegister(collectors.NewGoCollector())
	handler = promhttp.HandlerFor(h.params.registry, promhttp.HandlerOpts{Registry: register})

	return
}
