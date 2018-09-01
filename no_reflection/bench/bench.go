package bench

type Foo struct {
	Name string
}

func NewFoo() Foo {
	return Foo{Name: "foo"}
}

func MapFoo(f *Foo, fMap map[string]string) {
	fMap["Name"] = f.Name
}
