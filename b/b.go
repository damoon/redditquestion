package b

import "github.com/krocos/redditquestion/a"

type p struct{}

func (p) Instance(s string) a.I {
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
