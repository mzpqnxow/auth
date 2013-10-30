package martini

type Martini struct {
	handlers []interface{}
}

func New() *Martini {
	return &Martini{}
}

func (m *Martini) Use(handler interface{}) {
	m.handlers = append(m.handlers, handler)
}