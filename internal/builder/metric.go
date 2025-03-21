package builder

import (
	"github.com/goexl/promethy/internal/internal/metric"
	"github.com/goexl/promethy/internal/param"
	"github.com/prometheus/client_golang/prometheus"
)

type Metric struct {
	params *param.Metric
}

func NewMetric() *Metric {
	return &Metric{
		params: param.NewMetric(),
	}
}

func (m *Metric) Label(label string) (metric *Metric) {
	m.params.Labels = append(m.params.Labels, label)
	metric = m

	return
}

func (m *Metric) Gauge(gauge *prometheus.GaugeVec) *metric.Gauge {
	return metric.NewGauge(gauge, m.params.Labels...)
}

func (m *Metric) Counter(counter *prometheus.CounterVec) *metric.Counter {
	return metric.NewCounter(counter, m.params.Labels...)
}
