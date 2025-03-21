package promethy

type Counter interface {
	Inc()

	Add(float64)
}
