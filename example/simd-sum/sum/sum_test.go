package sum

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestAssembly(t *testing.T) {
	if got := Assembly(nil); got != 0 {
		t.Fatalf("expected 0 got %d: nil", got)
	}

	if got := Assembly([]int64{}); got != 0 {
		t.Fatalf("expected 0 got %d: nil", got)
	}

	for n := range 100 {
		t.Run(fmt.Sprintf("sum of %d", n), func(t *testing.T) {
			input := make([]int64, n)
			for i := range n {
				input[i] = int64(i)
			}

			want := Loop(input)
			if got := Assembly(input); got != want {
				t.Fatalf("expected %d got %d: %v", want, got, input)
			}
		})
	}
}

func TestIntrinsics(t *testing.T) {
	if got := Intrinsics(nil); got != 0 {
		t.Fatalf("expected 0 got %d: nil", got)
	}

	if got := Intrinsics([]int64{}); got != 0 {
		t.Fatalf("expected 0 got %d: nil", got)
	}

	for n := range 100 {
		t.Run(fmt.Sprintf("sum of %d", n), func(t *testing.T) {
			input := make([]int64, n)
			for i := range n {
				input[i] = int64(i)
			}

			want := Loop(input)
			if got := Intrinsics(input); got != want {
				t.Fatalf("expected %d got %d: %v", want, got, input)
			}
		})
	}
}

const inputSize = 100_000

func BenchmarkLoop(b *testing.B) {
	input := make([]int64, inputSize)
	for i := range input {
		input[i] = rand.Int63()
	}
	want := Loop(input)

	for b.Loop() {
		if got := Loop(input); got != want {
			b.Fatalf("expected %d got %d: %v", want, got, input)
		}
	}
}

func BenchmarkAssembly(b *testing.B) {
	input := make([]int64, inputSize)
	for i := range input {
		input[i] = rand.Int63()
	}
	want := Loop(input)

	for b.Loop() {
		if got := Assembly(input); got != want {
			b.Fatalf("expected %d got %d: %v", want, got, input)
		}
	}
}

func BenchmarkIntrinsics(b *testing.B) {
	input := make([]int64, inputSize)
	for i := range input {
		input[i] = rand.Int63()
	}
	want := Loop(input)

	for b.Loop() {
		if got := Intrinsics(input); got != want {
			b.Fatalf("expected %d got %d: %v", want, got, input)
		}
	}
}
