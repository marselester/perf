// Package loop-invariant shows the loop invariant code motion technique:
// keeping expressions that don't change outside of the loop.
// See Performance Analysis and Tuning on Modern CPUs by Denis Bakhvalov.
package main

import (
	"flag"
)

func main() {
	v := flag.String("version", "naive", "code version to run (naive, invariant)")
	flag.Parse()

	n := 1024
	x := make([]float32, n)
	y := make([]float32, n)
	z := make([]float32, n)
	for i := 0; i < len(x); i++ {
		x[i] = float32(i)
		y[i] = float32(i)
		z[i] = float32(i)
	}

	switch *v {
	case "invariant":
		multiplyInvariant(x, y, z)
	default:
		multiplyNaive(x, y, z)
	}
}

func multiplyNaive(a, b, c []float32) {
	n := len(a)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[j] = b[j] * c[i]
		}
	}
}

func multiplyInvariant(a, b, c []float32) {
	n := len(a)

	for i := 0; i < n; i++ {
		tmp := c[i]
		for j := 0; j < n; j++ {
			a[j] = b[j] * tmp
		}
	}
}
