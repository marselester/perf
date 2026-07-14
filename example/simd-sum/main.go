package main

import "perf/example/simd-sum/sum"

func main() {
	sum.Loop(make([]int64, 100_000))
}
