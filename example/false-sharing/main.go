// Program false-sharing shows the impact of false sharing,
// see Performance Analysis and Tuning on Modern CPUs by Denis Bakhvalov.
package main

import (
	"flag"
	"sync"
)

func main() {
	n := flag.Int("n", 100_000, "sum up to n")
	v := flag.String("version", "naive", "code version to run (naive, padding)")
	flag.Parse()

	switch *v {
	case "padding":
		sumPadding(*n)
	default:
		sumNaive(*n)
	}
}

func sumNaive(n int) {
	var (
		s  Sum
		wg sync.WaitGroup
	)

	wg.Add(2)
	go func() {
		s.CalcA(n)
		wg.Done()
	}()
	go func() {
		s.CalcB(n)
		wg.Done()
	}()

	wg.Wait()
}

func sumPadding(n int) {
	var (
		s  SumPadded
		wg sync.WaitGroup
	)

	wg.Add(2)
	go func() {
		s.CalcA(n)
		wg.Done()
	}()
	go func() {
		s.CalcB(n)
		wg.Done()
	}()

	wg.Wait()
}

type Sum struct {
	A int
	B int
}

func (s *Sum) CalcA(n int) {
	for i := 0; i < n; i++ {
		s.A += i
	}
}

func (s *Sum) CalcB(n int) {
	for i := 0; i < n; i++ {
		s.B += i
	}
}

type SumPadded struct {
	A int
	_ [64]byte // Padding to avoid false sharing.
	B int
}

func (s *SumPadded) CalcA(n int) {
	for i := 0; i < n; i++ {
		s.A += i
	}
}

func (s *SumPadded) CalcB(n int) {
	for i := 0; i < n; i++ {
		s.B += i
	}
}
