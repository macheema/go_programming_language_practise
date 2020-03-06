package exercise

// PopCountRightShift returns the population count (number of set bits) of x.
func PopCountRightShift(x uint64) int {
	var sum uint64 = 0
	for x > 0 {
		sum += x & 1
		x = x >> 1
	}
	return int(sum)
}
