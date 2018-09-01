package bench

import (
	"testing"
	"github.com/cvgw/reflect_benchmark/no_reflection/foo"
)

func BenchmarkMapFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := foo.NewFoo()
		fMap := make(map[string]string)
		MapFoo(&f, fMap)
	}
}

func BenchmarkMapFoo5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := foo.NewFoo5()
		fMap := make(map[string]string)
		MapFoo5(&f, fMap)
	}
}
