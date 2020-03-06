package main

import (
	"ch2/section26/exercise"
	"fmt"
)

func main() {
	kelvin := exercise.Kelvin(10)
	fmt.Printf("Kelive(%v) to Celsius(%v)", kelvin, exercise.KToC(kelvin))
}
