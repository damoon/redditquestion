package a

type P interface {
	Instance(s string) I
}

type I interface {
	F() string
}

type T struct {
	P P
}
