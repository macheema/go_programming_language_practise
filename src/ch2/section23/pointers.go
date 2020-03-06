package section23

import "fmt"

//Pointers sample to understand concept of pointer.
func Pointers() {
	x1 := new([]int)
	x2 := new([]int)
	x3 := new([0]int)
	x4 := new(struct{})
	fmt.Printf("%p,%p,%p,%p\n", x1, x2, x3, x4)
	fmt.Printf("%p,%p,%p,%p\n", new([]int), new([]int), new([0]int), new(struct{}))
}
