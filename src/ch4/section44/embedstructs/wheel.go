package embedstructs

//Point ...
type point struct {
	X, Y int
}

//Circle ...
type circle struct {
	point
	Radius int
}

//Wheel ...
type Wheel struct {
	circle
	Spokes int
}
