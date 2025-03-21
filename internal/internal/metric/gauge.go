package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Gauge struct {
	gauge  *prometheus.GaugeVec
	labels []string
}

func NewGauge(gauge *prometheus.GaugeVec, labels ...string) *Gauge {
	return &Gauge{
		gauge:  gauge,
		labels: labels,
	}
}

func (g *Gauge) Inc() {
	g.gauge.WithLabelValues(g.labels...).Inc()
}

func (g *Gauge) Add(value float64) {
	g.gauge.WithLabelValues(g.labels...).Add(value)
}

func (g *Gauge) Dec() {
	g.gauge.WithLabelValues(g.labels...).Dec()
}

func (g *Gauge) Sub(value float64) {
	g.gauge.WithLabelValues(g.labels...).Sub(value)
}

func (g *Gauge) Set(value float64) {
	g.gauge.WithLabelValues(g.labels...).Set(value)
}
