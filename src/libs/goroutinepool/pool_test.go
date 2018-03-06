package goroutinepool

import (
	"testing"
)

//go test -bench="Benchmark_GoPool"
//goos: linux
//goarch: amd64
//pkg: libs/goroutinepool
//Benchmark_GoPool-8   	30000000	        64.3 ns/op
//PASS
//ok  	libs/goroutinepool	2.033s
func Benchmark_GoPool(b *testing.B) {
	var pool = New(8, b.N)
	pool.Start()
	for i := 0; i < b.N; i++ {
		var _ = i
		pool.AddJob(func() {
			for j := 0; j < 100000; j++ {
			}
		})
	}
}

//go test -bench="Benchmark_NoPool"
//goos: linux
//goarch: amd64
//pkg: libs/goroutinepool
//Benchmark_NoPool-8   	 1000000	      6322 ns/op
//PASS
//ok  	libs/goroutinepool	6.590s
func Benchmark_NoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = i
		go func() {
			for j := 0; j < 100000; j++ {

			}
		}()
	}
}
