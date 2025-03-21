package param

type Metric struct {
	Labels []string
}

func NewMetric() *Metric {
	return &Metric{
		Labels: make([]string, 0),
	}
}
