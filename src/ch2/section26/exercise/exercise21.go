package exercise

import (
	"ch2/section26"
	"fmt"
)

//Kelvin represents unit of temperature
type Kelvin float64

func (f Kelvin) String() string { return fmt.Sprintf("%g°K", f) }

const (
	//AbsoluteZeroK is absolut point in Kelvin where electrons of matter become still
	AbsoluteZeroK Kelvin = Kelvin(section26.AbsoluteZeroC - section26.AbsoluteZeroC)
	//FreezingK is Freezing point in Kelvin
	FreezingK Kelvin = Kelvin(section26.FreezingC - section26.AbsoluteZeroC)
	//BoilingK is boiling point of water in Kelvin
	BoilingK Kelvin = Kelvin(section26.BoilingC - section26.AbsoluteZeroC)
)

//CToK Converts Celsius to Kelvin
func CToK(c section26.Celsius) Kelvin { return Kelvin(c + 273.15) }

//KToC Converts Kelvin to Celsius
func KToC(k Kelvin) section26.Celsius { return section26.Celsius(k - 273.15) }

//FToK Converts Fahrenheit to Kelvin
func FToK(f section26.Fahrenheit) Kelvin { return CToK(section26.FToC(f)) }

//KToF Converts Kelvin to Fahrenheit
func KToF(k Kelvin) section26.Fahrenheit { return section26.CToF(KToC(k)) }
