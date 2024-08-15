package main

import "testing"

func BenchmarkSumNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumNaive(100_000)
	}
}

func BenchmarkSumPadded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumPadding(100_000)
	}
}
