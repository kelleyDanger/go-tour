package main

import "fmt"

// Struct are basically like Python Classes

type Vertex struct {
	X int
	Y int
}

// 4 Struct Literals
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1} // default Y:0
	v3 = Vertex{}     // default X:0 and Y:0
	p  = &Vertex{1, 2}
)

func main() {
	// 1 Instantiate
	v := Vertex{1, 2}

	// 2 Access Fields
	fmt.Println("v.X Before: ", v.X)
	fmt.Println("v.Y Before: ", v.Y)
	fmt.Println("Setting v.X to 4....")
	v.X = 4
	fmt.Println("v.X After: ", v.X)

	// 3 Pointers to Structs
	p := &v
	fmt.Println("*p: ", *p)
	fmt.Println("Pointer v Before: ", v)
	fmt.Println("Setting v.Y to 100...")
	p.Y = 100
	fmt.Println("Pointer v After: ", v)

	// 4 Struct Literals
	fmt.Println("Struct Literals")
	fmt.Println(v1, v2, v3, p)

}
