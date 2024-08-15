// Program loop-unrolling shows the impact of loop unrolling, see https://sourcegraph.com/blog/slow-to-simd.
package main

import (
	"flag"
	"fmt"
)

func main() {
	v := flag.String("version", "naive", "code version to run (naive, unroll, bound)")
	flag.Parse()

	x := make([]float32, 1024)
	y := make([]float32, 1024)

	var sum float32
	switch *v {
	case "unroll":
		sum = dotUnroll(x, y)
	case "bound":
		sum = dotBoundsCheckingElimination(x, y)
	default:
		sum = dotNaive(x, y)
	}

	fmt.Println(sum)
}

func dotNaive(a, b []float32) float32 {
	var sum float32
	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}

	return sum
}

func dotUnroll(a, b []float32) float32 {
	var sum float32
	for i := 0; i < len(a); i += 4 {
		s0 := a[i] * b[i]
		s1 := a[i+1] * b[i+1]
		s2 := a[i+2] * b[i+2]
		s3 := a[i+3] * b[i+3]

		sum += s0 + s1 + s2 + s3
	}

	return sum
}

func dotBoundsCheckingElimination(a, b []float32) float32 {
	if len(a) != len(b) {
		panic("slices must have equal lengths")
	}

	if len(a)%4 != 0 {
		panic("slice length must be multiple of 4")
	}

	var sum float32
	for i := 0; i < len(a); i += 4 {
		aTmp := a[i : i+4 : i+4]
		bTmp := b[i : i+4 : i+4]

		s0 := aTmp[0] * bTmp[0]
		s1 := aTmp[1] * bTmp[1]
		s2 := aTmp[2] * bTmp[2]
		s3 := aTmp[3] * bTmp[3]

		sum += s0 + s1 + s2 + s3
	}

	return sum
}
