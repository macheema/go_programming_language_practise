package exercise

import (
	"ch4/section41"
	"fmt"
)

//PC is population count of all numbers which can be reperensted in a single byte
var PC [256]byte

func init() {
	for i := range PC {
		PC[i] = PC[i/2] + byte(i&1)
	}
}

//BinaryCount ...
func BinaryCount(hash [32]uint8) byte {

	sum := byte(0)
	for i, val := range hash {
		println(i, val)
		sum += PC[val]
	}
	return sum
}

//Exercise41 ...
func Exercise41() {
	hash1, hash2 := section41.SHA256("y", "Y")
	fmt.Printf("hash1 : %b\n hash2 : %b\n", hash1, hash2)
	fmt.Printf("hash1 : PC : %v", BinaryCount(hash1))
	fmt.Printf("hash2 : PC : %v", BinaryCount(hash2))
}
