package main

import (
	"fmt"

	"github.com/krocos/redditquestion/a"
	"github.com/krocos/redditquestion/b"
)

func main() {
	t := a.T{P: b.NewP()}

	i := t.P.Instance("String to return.")

	fmt.Println((i.(a.I)).F())
}
