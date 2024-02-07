package main

import (
	"fmt"
	"math"
)

// Functions are Values Too!
// Can be....
// - 1. passed around
// - 2. function arguments
// - 3. return values

// 2. function argument
func compute(
	fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	// 1. passed around
	fmt.Println(compute(hypot))
	// 3. return value
	fmt.Println(compute(math.Pow))
}
