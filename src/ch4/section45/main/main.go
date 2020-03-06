package main

import "ch4/section45"

func main() {
	var json = section45.PrintMoviesJSON()
	section45.PrintMoviesJSONIndented()
	section45.PrintMoviesFromJSON(json)
}
