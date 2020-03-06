package exercise

// PopCountClearRightShift returns the population count (number of set bits) of x.
func PopCountClearRightShift(x uint64) int {
	var sum uint64 = 0
	for x > 0 {
		x = x & (x - 1)
		if (x & 1) == 0 {
			sum++
		}
		x = x >> 1
	}
	return int(sum)
}
