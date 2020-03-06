package section42

import "fmt"

//Reverse1 ...
func Reverse1() {
	slice := []int{0, 1, 2, 3, 4, 5}
	Reverse(slice)
	fmt.Printf("%v", slice)
}

//Reverse ...
func Reverse(param []int) {
	for i, j := 0, len(param)-1; i <= j; i, j = i+1, j-1 {
		param[i], param[j] = param[j], param[i]
	}
}
