package main

import "testing"

func BenchmarkMultiplyNaive(b *testing.B) {
	n := 1024
	x := make([]float32, n)
	y := make([]float32, n)
	z := make([]float32, n)
	for i := 0; i < len(x); i++ {
		x[i] = float32(i)
		y[i] = float32(i)
		z[i] = float32(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		multiplyNaive(x, y, z)
	}
}

func BenchmarkMultiplyInvariant(b *testing.B) {
	n := 1024
	x := make([]float32, n)
	y := make([]float32, n)
	z := make([]float32, n)
	for i := 0; i < len(x); i++ {
		x[i] = float32(i)
		y[i] = float32(i)
		z[i] = float32(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		multiplyInvariant(x, y, z)
	}
}
