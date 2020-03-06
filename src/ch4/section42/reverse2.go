package section42

import "fmt"

//Reverse2 output should be [2,3,4,5,0,1]...
func Reverse2() {
	slice := []int{0, 1, 2, 3, 4, 5}
	Reverse(slice)
	Reverse(slice[:4])
	Reverse(slice[4:])
	fmt.Printf("%v", slice)
}

//Reverse3 output should be [2,3,4,5,0,1]...
func Reverse3() {
	slice := []int{0, 1, 2, 3, 4, 5}
	Reverse(slice[:2]) // [1,0,2,3,4,5]
	Reverse(slice[2:]) // [1,0,5,4,3,2]
	Reverse(slice)     // [2,3,4,5,0,1]
	fmt.Printf("%v", slice)
}

//Reverse4 output should be [rotateBy,4,5,0,1,2] as if rotateBy is 3
func Reverse4(rotateBy uint) {
	slice := []int{}
	for i := uint(0); i < rotateBy*2; i++ {
		slice = append(slice, int(i))
	}
	Reverse(slice[:rotateBy])
	Reverse(slice[rotateBy:])
	Reverse(slice)
	fmt.Printf("%v", slice)
}
