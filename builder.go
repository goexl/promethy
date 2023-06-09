package promethy

import (
	"github.com/goexl/simaqian"
	"github.com/goexl/valuer"
)

type builder struct {
	params *params
}

func newBuilder() *builder {
	return &builder{
		params: newParams(),
	}
}

func (b *builder) Logger(logger simaqian.Logger) (builder *builder) {
	b.params.logger = logger
	builder = b

	return
}

func (b *builder) Parser(parser *valuer.Parser) (builder *builder) {
	b.params.parser = parser
	builder = b

	return
}

func (b *builder) Label(key string, value string) (builder *builder) {
	b.params.labels[key] = value
	builder = b

	return
}

func (b *builder) Labels(labels map[string]string) (builder *builder) {
	for key, value := range labels {
		b.params.labels[key] = value
	}
	builder = b

	return
}

func (b *builder) Build() *prom {
	return newProm(b.params)
}
