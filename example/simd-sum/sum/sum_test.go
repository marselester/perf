package sum

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestVectors(t *testing.T) {
	if got := Vectors(nil); got != 0 {
		t.Fatalf("expected 0 got %d: nil", got)
	}

	if got := Vectors([]int64{}); got != 0 {
		t.Fatalf("expected 0 got %d: nil", got)
	}

	for n := range 100 {
		t.Run(fmt.Sprintf("sum of %d", n), func(t *testing.T) {
			input := make([]int64, n)
			for i := range n {
				input[i] = int64(i)
			}

			want := Scalars(input)
			if got := Vectors(input); got != want {
				t.Fatalf("expected %d got %d: %v", want, got, input)
			}
		})
	}
}

const inputSize = 100_000

func BenchmarkScalars(b *testing.B) {
	input := make([]int64, inputSize)
	for i := range input {
		input[i] = rand.Int63()
	}
	want := Scalars(input)

	for b.Loop() {
		if got := Scalars(input); got != want {
			b.Fatalf("expected %d got %d: %v", want, got, input)
		}
	}
}

func BenchmarkVectors(b *testing.B) {
	input := make([]int64, inputSize)
	for i := range input {
		input[i] = rand.Int63()
	}
	want := Scalars(input)

	for b.Loop() {
		if got := Vectors(input); got != want {
			b.Fatalf("expected %d got %d: %v", want, got, input)
		}
	}
}
