package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
	// 1 init map
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Home"] = Vertex{40.1230, -12.092}
	fmt.Println("Home Coords: ", m["Home"])

	// 2 literals
	var map2 = map[string]Vertex{
		"Santa Workshop": Vertex{-100.129, -100.123},
		"Atlantis":       Vertex{80.89, 70.102},
	}
	fmt.Println("Map2: ", map2)

	// 3 Mutating
	map3 := make(map[string]int)

	// 3.1 Insert / Update
	map3["Answer"] = 42
	fmt.Println("Answer: ", map3["Answer"])

	// 3.2 Delete
	delete(map3, "Answer")
	fmt.Println("Deleted Answer....: ", map3["Answer"])

	// 3.3 Check Key Is Present
	v, ok := map3["Answer"]
	fmt.Println("Value: ", v, " Present? ", ok)

}

// EXERCISE
func WordCount(s string) map[string]int {
	fmt.Println("s: ", s)

	var m map[string]int
	m = make(map[string]int)

	var words []string = strings.Fields(s)
	for _, word := range words {
		fmt.Println("word: ", word)
		m[word] += 1
	}

	return m
}
