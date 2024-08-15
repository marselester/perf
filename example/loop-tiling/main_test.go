package main

import "testing"

func BenchmarkAddNaive(b *testing.B) {
	x := make([][]float32, 1024)
	y := make([][]float32, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = make([]float32, 1024)
		y[i] = make([]float32, 1024)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		addNaive(x, y)
	}
}

func BenchmarkAddTiling(b *testing.B) {
	x := make([][]float32, 1024)
	y := make([][]float32, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = make([]float32, 1024)
		y[i] = make([]float32, 1024)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		addTiling(x, y)
	}
}
