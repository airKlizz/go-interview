package routines

// SumSlices takes two slices of integers, sums them up concurrently using goroutines,
// and returns the total sum.
func SumSlices(slice1, slice2 []int) int {
	panic("unimplemented")
}

func sumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
