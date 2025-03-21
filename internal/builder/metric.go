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

func (m *Metric) Label(required string, optionals ...string) (metric *Metric) {
	m.params.Labels = append(m.params.Labels, required)
	m.params.Labels = append(m.params.Labels, optionals...)
	metric = m

	return
}

func (m *Metric) Labels(labels []string) (metric *Metric) {
	m.params.Labels = append(m.params.Labels, labels...)
	metric = m

	return
}

func (m *Metric) Gauge(gauge *prometheus.GaugeVec) *metric.Gauge {
	return metric.NewGauge(gauge, m.params.Labels...)
}

func (m *Metric) Counter(counter *prometheus.CounterVec) *metric.Counter {
	return metric.NewCounter(counter, m.params.Labels...)
}

func (m *Metric) Histogram(histogram *prometheus.HistogramVec) *metric.Histogram {
	return metric.NewHistogram(histogram, m.params.Labels...)
}
