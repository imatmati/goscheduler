package main

import (
	"fmt"
)

type Itf interface {
	Foo(obj Itf)
}

type A struct {
	i int
}

func (a A) Foo(obj Itf) {
	fmt.Println(obj.(A).i)
}

type B struct {
	b string
	A
}

func main() {

	a := A{15}
	a.Foo(a)
	b := B{b: "ok", A: A{i: 10}}
	a = b.A
}
