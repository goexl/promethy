package promethy

import (
	"github.com/goexl/promethy/internal/builder"
)

func Metric() *builder.Metric {
	return builder.NewMetric()
}
