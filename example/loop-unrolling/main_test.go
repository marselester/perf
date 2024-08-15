package main

import "testing"

var Sum float32

func BenchmarkDotNaive(b *testing.B) {
	x := make([]float32, 1024)
	y := make([]float32, 1024)
	b.ResetTimer()

	var sum float32
	for i := 0; i < b.N; i++ {
		sum = dotNaive(x, y)
	}
	Sum = sum
}

func BenchmarkDotUnroll(b *testing.B) {
	x := make([]float32, 1024)
	y := make([]float32, 1024)
	b.ResetTimer()

	var sum float32
	for i := 0; i < b.N; i++ {
		sum = dotUnroll(x, y)
	}
	Sum = sum
}

func BenchmarkDotBoundsCheckingElimination(b *testing.B) {
	x := make([]float32, 1024)
	y := make([]float32, 1024)
	b.ResetTimer()

	var sum float32
	for i := 0; i < b.N; i++ {
		sum = dotBoundsCheckingElimination(x, y)
	}
	Sum = sum
}
