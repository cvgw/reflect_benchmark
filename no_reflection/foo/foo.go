package foo

type Foo struct {
	Name string
}

func NewFoo() Foo {
	return Foo{
		Name: "foo",
	}
}

func FooFields() []string {
	return []string{
		"Name",
	}
}


type Foo5 struct {
	Name string
	A string
	B string
	C string
	D string
}

func NewFoo5() Foo5 {
	return Foo5{
		Name: "foo",
		A: "a",
		B: "b",
		C: "c",
		D: "d",
	}
}

func Foo5Fields() []string {
	return []string{
		"Name",
		"A",
		"B",
		"C",
		"D",
	}
}
