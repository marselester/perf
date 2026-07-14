package sum

//go:noinline
func Loop(input []int64) int64 {
	var sum int64
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	return sum
}
