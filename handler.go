package promethy

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
	labels := make(map[string]string)
	for key, value := range h.params.labels {
		labels[key] = h.params.parser.Parse(value)
	}
	register := prometheus.WrapRegistererWith(labels, h.params.registry)
	if pe := register.Register(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{})); nil != pe {
		err = pe
	} else if ge := register.Register(collectors.NewGoCollector()); nil != ge {
		err = ge
	} else {
		handler = promhttp.HandlerFor(h.params.registry, promhttp.HandlerOpts{Registry: register})
	}

	return
}
