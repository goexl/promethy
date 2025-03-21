package promethy

type Gauge interface {
	Counter

	Dec()

	Sub(float64)

	Set(float64)
}
