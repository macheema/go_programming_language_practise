package main

import (
	"ch4/section42/exercise"
	"fmt"
)

func main() {
	//Exercise43
	// arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Exercise43(&arr)
	// fmt.Printf("%v", arr)

	//Exercise44
	// slice := []int{0, 1, 2, 3, 4, 5, 6, 7}
	// Exercise44(&slice, 2)
	// fmt.Printf("%v\n", slice[:])

	//Exercise46
	// data := "aAAbCCdEEfGGGhh"
	// data = Exercise45(data)
	// println(data)

	//Exercise47
	data := "АДЖИ"
	bytes := []byte(data)
	exercise.Exercise47(bytes)
	fmt.Println(string(bytes))
}
