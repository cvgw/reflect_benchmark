package main

import (
	"fmt"
	"github.com/cvgw/reflect_benchmark/no_reflection/bench"
)

func main() {
	f := bench.Foo{Name: "foo"}

	fmt.Printf("Name: %s\n", f.Name)

	fMap := make(map[string]string)

	bench.MapFoo(f, fMap)
}
