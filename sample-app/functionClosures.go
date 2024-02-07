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

// EXERCISE: fibonacci closure
// fibonacci is a func that returns a func that returns an int
func fibonacci() func() int {
	// 0 1 1 2 3 5 8 13 21 ...
	// first = second
	// second = first + second
	first, second := 0, 1
	return func() int {
		saveFirst := first
		first, second = second, first+second
		return saveFirst
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

	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
