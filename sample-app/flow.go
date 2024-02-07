package main

import (
	"fmt"
	"math"
	"runtime"
)

// 1 for loop
func test_for() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("for: ", sum)
}

// 1.5 minimal for loop (While loop)
func test_minimal_for() {
	sum := 1
	// 1st and 3rd are optional
	for sum < 5 {
		sum += sum
	}
	fmt.Println("minimal for/while loop: ", sum)
}

// 2 if
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// 2.5 if/else
// w/ short statement
// note: 'v' is only accessible inside if loop
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// 3 switch
// note: only runs first matching case; no need for 'break'
func test_switch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

// 4 defer
// executes AFTER function returns
func test_defer() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// 4.5 stacking defers
func stacking_defers() {
	fmt.Println("counting...")
	for i := 0; i < 10; i++ {
		defer fmt.Print(" ", i, " ")
		// 9 8 7 ... 1 0
	}
	fmt.Println("Done!")
}

func main() {

	// 1 for loop
	test_for()

	// 1.5 minimal for loop (While loop)
	test_minimal_for()

	// 2 if
	fmt.Println("if: ", sqrt(2), sqrt(-4))

	/// 2.5 if/else w/ short statement
	fmt.Println("if/else w/ short statement: ", pow(3, 2, 10))
	fmt.Println("if/else w/ short statement: ", pow(3, 3, 20))

	// 3 switch
	test_switch()

	// 4 defer
	test_defer()

	// 4.5 stacking defers
	stacking_defers()
}
