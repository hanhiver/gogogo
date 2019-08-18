package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1968, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "jacqueline Bisset"}},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n\n", data)

	data, err = json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n\n", data)

	var titles []struct{ Title string }
	if err = json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)

	var actors []struct{ Actors []string }
	if err = json.Unmarshal(data, &actors); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(actors)
}
