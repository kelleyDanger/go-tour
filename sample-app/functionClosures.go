package main

import "fmt"

// Function Closures
// - closure: function value that references variables from outside its body
// - function may access/assign referenced variables

// Returns a closure
// - each closure is bound to its own 'sum' variable
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			"Pos of ", i, pos(i),
			"Neg of ", i, neg(-i),
		)
	}
}
