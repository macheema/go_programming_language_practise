package section45

import (
	"encoding/json"
	"fmt"
)

//GetMovies ...
func GetMovies() []Movie {
	return []Movie{{Title: "Movie ABC", Year: 2020, Color: false, Actors: []string{"ABC", "XYZ"}},
		{Title: "Movie XYZ", Year: 2020, Color: false, Actors: []string{"ABC", "XYZ"}}}
}

//PrintMoviesJSON ...
func PrintMoviesJSON() string {
	var movies = GetMovies()
	fmt.Println(movies)
	var data, err = json.Marshal(movies)
	if err != nil {
		fmt.Printf("JSON marshalling failed : %s\n", err)
	}
	fmt.Printf("%s\n", data)
	return string(data)
}

// PrintMoviesJSONIndented ...
func PrintMoviesJSONIndented() string {
	var movies = GetMovies()
	fmt.Println(movies)
	var data, err = json.MarshalIndent(movies, "#", " ")
	if err != nil {
		fmt.Printf("JSON marshalling failed : %s\n", err)
	}
	fmt.Printf("%s\n", data)
	return string(data)
}

//PrintMoviesFromJSON ...
func PrintMoviesFromJSON(moviesJSONStr string) {
	var movies = []Movie{}

	var err = json.Unmarshal([]byte(moviesJSONStr), &movies)
	if err != nil {
		fmt.Printf("JSON unmarshalling failed : %s\n", err)
	}
	fmt.Println(movies)
}
