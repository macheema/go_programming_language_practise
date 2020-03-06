package section27

import "fmt"

//Scope ...
func Scope() {
	RedeclareCheck()
}

//Redeclare -> In different level of scope var declarations shadow the previous declarations
func Redeclare() {
	x := 10
	y := 10
	{
		x, y := 20, 20
		fmt.Printf("x-> %p\n", &x)
		fmt.Printf("y-> %p\n", &y)
	}
	fmt.Printf("x-> %p\n", &x)
	fmt.Printf("y-> %p\n", &y)
}

//RedeclareCheck -> In same level of scope var declarations resuse the previous declarations but syntactally legal if
//one of all variable is not already declare
func RedeclareCheck() {
	x := 10
	// y := 10
	x, y := 20, 20
	fmt.Printf("x-> %p\n", &x)
	fmt.Printf("y-> %p\n", &y)
	{
		x, y := 20, 20
		fmt.Printf("x-> %p\n", &x)
		fmt.Printf("y-> %p\n", &y)
	}
	fmt.Printf("x-> %p\n", &x)
	fmt.Printf("y-> %p\n", &y)
}
