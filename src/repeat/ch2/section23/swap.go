package section23

import "fmt"

func Swap1() {
	x, y := 10, 20
	x, y = y, x //pre,post increment and decrement are not allowed
	fmt.Printf("x=%d,y=%d", x, y)
}
