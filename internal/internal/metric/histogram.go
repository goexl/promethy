package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Histogram struct {
	histogram *prometheus.HistogramVec
	labels    []string
}

func NewHistogram(histogram *prometheus.HistogramVec, labels ...string) *Histogram {
	return &Histogram{
		histogram: histogram,
		labels:    labels,
	}
}

func (h *Histogram) Observe(value float64) {
	h.histogram.WithLabelValues(h.labels...).Observe(value)
}
