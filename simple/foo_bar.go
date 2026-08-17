package simple

type Foo struct {
}

type Bar struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

func NewBar() *Bar {
	return &Bar{}
}

type FooBar struct {
	Foo *Foo
	Bar *Bar
}