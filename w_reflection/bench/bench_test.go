package bench

import "testing"

func BenchmarkMapFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := NewFoo()
		fields := FooFields()
		fMap := make(map[string]interface{})
		MapFoo(&f, fMap, fields)
	}
}
