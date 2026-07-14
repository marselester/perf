package sum

import "simd/archsimd"

//go:noinline
func IntrinsicsBoundsChecks(input []int64) (sum int64) {
	i := 0
	loopEnd := len(input) - len(input)%4
	if loopEnd >= 8 {
		y0 := archsimd.LoadInt64x4Slice(input) // y0 := input[0:4]
		for i += 4; i < loopEnd; i += 4 {
			y1 := archsimd.LoadInt64x4Slice(input[i : i+4]) // y1 := input[i:i+4]
			y0 = y0.Add(y1)                                 // y0 += y1
		}

		// Horizontal reduction.
		x0, x1 := y0.GetLo(), y0.GetHi()    // x0, x1 := y0[0:2], y0[2:4]
		x0 = x0.Add(x1)                     // x0 += x1
		sum = x0.GetElem(0) + x0.GetElem(1) // sum = x0[0] + x0[1]
	}

	for ; i < len(input); i++ {
		sum += input[i]
	}

	return sum
}
