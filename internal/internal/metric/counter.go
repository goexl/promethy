package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Counter struct {
	counter *prometheus.CounterVec
	labels  []string
}

func NewCounter(counter *prometheus.CounterVec, labels ...string) *Counter {
	return &Counter{
		counter: counter,
		labels:  labels,
	}
}

func (c *Counter) Inc() {
	c.counter.WithLabelValues(c.labels...).Inc()
}

func (c *Counter) Add(value float64) {
	c.counter.WithLabelValues(c.labels...).Add(value)
}
