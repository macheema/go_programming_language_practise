package section23

import "fmt"

//Tuples sample to understand concept of pointer.
func Tuples() {

	var x int
	x = 10
	ok := false
	// val, ok := x.(int) // invalid line as Type assertion can be check on interface
	var y assertCheck
	var xObject = assertCheck2{result: 1}
	var xI test = xObject
	xI = nil
	// return underneath value if it is of that type but return default new object with default values
	//if object is not of that type
	val, ok := xI.(assertCheck2)
	if !ok {
		fmt.Printf("%v,%p,%v,%p,%v,%p,%v,%v\n", val, &val, xI, &xI, xObject, &xObject, y, x)
	}
}

type assertCheck struct {
	result int
}

type assertCheck2 struct {
	result int
}

// func (*assertCheck) test() {
// }
func (assertCheck) test() {
}
func (assertCheck2) test() {
}

type test interface {
	test()
}
