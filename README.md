# Reflect Benchmarking
## Description

A simple set of code to demonstrate the performance impact of using the reflect
package in goland.

## Execute Benchmarks
Navigate to the `bench` package in the directory of the code you wish to benchmark.

Execute the benchmark with `go test -bench=.`

## Results

### With Reflection
````
goos: darwin
goarch: amd64
pkg: github.com/cvgw/reflect_benchmark/w_reflection/bench
BenchmarkMapFoo-8   	10000000	       220 ns/op
PASS
ok  	github.com/cvgw/reflect_benchmark/w_reflection/bench	2.442s
````

### Without Reflection
````
goos: darwin
goarch: amd64
pkg: github.com/cvgw/reflect_benchmark/no_reflection/bench
BenchmarkMapFoo-8   	30000000	        44.9 ns/op
PASS
ok  	github.com/cvgw/reflect_benchmark/no_reflection/bench	5.197s
````
