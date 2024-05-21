package main

import (
	"runtime"
	"testing"
)

var (
	nums    int   = 10000000
	numbers []int = make([]int, nums)
)

// $ GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s
// goos: darwin
// goarch: arm64
// pkg: gmp
// BenchmarkSequential         1204           2885547 ns/op
// BenchmarkConcurrent         1081           3253782 ns/op
// PASS
// ok      gmp     7.756s

// $ GOGC=off go test -cpu 8 -run none -bench . -benchtime 3s
// goos: darwin
// goarch: arm64
// pkg: gmp
// BenchmarkSequential-8               1174           2880741 ns/op
// BenchmarkConcurrent-8               2634           1314148 ns/op
// PASS
// ok      gmp     7.764s

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(numbers)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addConcurrent(runtime.NumCPU(), numbers)
	}
}
