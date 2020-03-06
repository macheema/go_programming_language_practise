package section26

import "fmt"

//Celsius represents unit of temperature
type Celsius float64

//Fahrenheit represents unit of temperature
type Fahrenheit float64

const (
	//AbsoluteZeroC is absolut point in Celsius where electrons of matter become still
	AbsoluteZeroC Celsius = -273.15
	//FreezingC is Freezing point in Celsius
	FreezingC Celsius = 0
	//BoilingC is boiling point of water in Celsius
	BoilingC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

//CToF Converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//FToC Converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
