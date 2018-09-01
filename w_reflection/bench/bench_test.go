package bench

import (
	"testing"
	"github.com/cvgw/reflect_benchmark/no_reflection/foo"
)

func BenchmarkMapFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := foo.NewFoo()
		fields := foo.FooFields()
		fMap := make(map[string]interface{})
		MapFoo(&f, fMap, fields)
	}
}

func BenchmarkMapFoo5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := foo.NewFoo5()
		fields := foo.Foo5Fields()
		fMap := make(map[string]interface{})
		MapFoo(&f, fMap, fields)
	}
}
