# Reflect Benchmarking
## Description

A simple set of code to demonstrate the performance impact of using the reflect
package in goland.

## Execute Benchmarks
Navigate to the `bench` package in the directory of the code you wish to benchmark.

Execute the benchmark with `go test -bench=.`

## Results
`Foo` benchmarks a struct with a single field

`Foo5` benchmarks a struct with five fields

### With Reflection
````
goos: darwin
goarch: amd64
pkg: github.com/cvgw/reflect_benchmark/w_reflection/bench
BenchmarkMapFoo-8    	10000000	       222 ns/op
BenchmarkMapFoo5-8   	 1000000	      1120 ns/op
PASS
ok  	github.com/cvgw/reflect_benchmark/w_reflection/bench	7.460s
````

### Without Reflection
````
goos: darwin
goarch: amd64
pkg: github.com/cvgw/reflect_benchmark/no_reflection/bench
BenchmarkMapFoo-8    	30000000	        41.1 ns/op
BenchmarkMapFoo5-8   	10000000	       169 ns/op
PASS
ok  	github.com/cvgw/reflect_benchmark/no_reflection/bench	7.026s
````
