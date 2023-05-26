package prometheus

type prom struct {
	params *params
}

func newProm(params *params) *prom {
	return &prom{
		params: params,
	}
}

func (p *prom) Register() *Registry {
	return newRegistry(p.params)
}

func (p *prom) Handler() *Handler {
	return newHandler(p.params)
}
