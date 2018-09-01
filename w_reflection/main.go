package main

import (
	"fmt"
	"github.com/cvgw/reflect_benchmark/w_reflection/bench"
)

func main() {
	f := bench.Foo{Name: "foo"}
	fields := []string{"Name"}

	fmt.Printf("Name: %s\n", f.Name)

	fMap := make(map[string]interface{})

	bench.MapFoo(&f, fMap, fields)
}
