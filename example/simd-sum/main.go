package main

import "perf/example/simd-sum/sum"

func main() {
	sum.Scalars(make([]int64, 100_000))
}
