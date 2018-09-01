package bench

import "github.com/cvgw/reflect_benchmark/no_reflection/foo"

func MapFoo(f *foo.Foo, fMap map[string]string) {
	fMap["Name"] = f.Name
}

func MapFoo5(f *foo.Foo5, fMap map[string]string) {
	fMap["Name"] = f.Name
	fMap["A"] = f.A
	fMap["B"] = f.B
	fMap["C"] = f.C
	fMap["D"] = f.D
}
