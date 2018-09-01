package bench

import "testing"

func BenchmarkMapFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := NewFoo()
		fMap := make(map[string]string)
		MapFoo(&f, fMap)
	}
}
