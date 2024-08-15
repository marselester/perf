// Package loop-tiling shows the loop tiling technique: keeping data in the CPU cache
// by traversing the matrix in 8x8 blocks instead of traversing it linearly.
// See Performance Analysis and Tuning on Modern CPUs by Denis Bakhvalov.
package main

import (
	"flag"
)

func main() {
	v := flag.String("version", "naive", "code version to run (naive, tiling)")
	flag.Parse()

	n := 1024
	x := make([][]float32, n)
	y := make([][]float32, n)
	for i := 0; i < len(x); i++ {
		x[i] = make([]float32, n)
		y[i] = make([]float32, n)
	}

	switch *v {
	case "tiling":
		addTiling(x, y)
	default:
		addNaive(x, y)
	}
}

func addNaive(a, b [][]float32) {
	n := len(a)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] += b[j][i]
		}
	}
}

func addTiling(a, b [][]float32) {
	n := len(a)

	for ib := 0; ib < n; ib += 8 {
		for jb := 0; jb < n; jb += 8 {
			for i := ib; i < ib+8; i++ {
				for j := jb; j < jb+8; j++ {
					a[i][j] += b[j][i]
				}
			}
		}
	}
}
