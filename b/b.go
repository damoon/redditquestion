package b

type p struct{}

func (p) Instance(s string) interface{} {
	return &i{s: s}
}

type i struct {
	s string
}

func (i i) F() string {
	return i.s
}

func NewP() *p {
	return &p{}
}
