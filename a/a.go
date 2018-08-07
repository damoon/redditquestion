package a

type P interface {
	Instance(s string) interface{}
}

type I interface {
	F() string
}

type T struct {
	P P
}
