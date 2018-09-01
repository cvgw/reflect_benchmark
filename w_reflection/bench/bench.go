package bench

import "reflect"

type Foo struct {
	Name string
}

func NewFoo() Foo {
	return Foo{Name: "foo"}
}

func FooFields() []string {
	return []string{"Name"}
}

func MapFoo(f interface{}, fMap map[string]interface{}, fields []string) {
	val := reflect.ValueOf(f)

	e := val.Elem()

	for _, field := range fields {
		sField := e.FieldByName(field)
		fVal := sField.String()
		fMap[field] = fVal
	}
}
