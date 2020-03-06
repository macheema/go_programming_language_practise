package exercise

import "ch2/section26"

// PopCountWithLoop returns the population count (number of set bits) of x.
func PopCountWithLoop(x uint64) int {
	var sum int = 0
	for i := uint(0); i < 8; i++ {
		sum = int(section26.Pc[byte(x>>(i*8))])
	}
	return sum
}
