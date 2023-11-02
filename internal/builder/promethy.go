package builder

import (
	"github.com/goexl/log"
	"github.com/goexl/promethy/internal/core"
	"github.com/goexl/promethy/internal/param"
	"github.com/goexl/valuer"
)

type Promethy struct {
	params *param.Promethy
}

func NewPromethy() *Promethy {
	return &Promethy{
		params: param.NewParams(),
	}
}

func (p *Promethy) Logger(logger log.Logger) (builder *Promethy) {
	p.params.Logger = logger
	builder = p

	return
}

func (p *Promethy) Parser(parser *valuer.Parser) (builder *Promethy) {
	p.params.Parser = parser
	builder = p

	return
}

func (p *Promethy) Label(key string, value string) (builder *Promethy) {
	p.params.Labels[key] = value
	builder = p

	return
}

func (p *Promethy) Labels(labels map[string]string) (builder *Promethy) {
	for key, value := range labels {
		p.params.Labels[key] = value
	}
	builder = p

	return
}

func (p *Promethy) Build() *core.Promethy {
	return core.NewPromethy(p.params)
}
