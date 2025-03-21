package promethy

type Histogram interface {
	Observe(float64)
}
