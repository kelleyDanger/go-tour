package main

import (
	"fmt"
	"strings"
)

// Slices is dynamically-sized.
// More command than arrays.
// var slice [n]T = array[low : high]
// includes low, excludes high

func main() {
	// 1 init
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var slice []int = primes[1:4] // elem 1-3
	fmt.Println(slice)

	// 2 like references; change underlying array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("Names before: ", names)

	a := names[0:2] // John, Paul
	b := names[1:3] // Paul, George
	fmt.Println("a, b", a, b)

	// updating slice updates array
	b[0] = "XXX" // Paul -> "XXX"
	fmt.Println("a, b (Paul -> 'XXX')", a, b)
	fmt.Println("Names (Paul -> 'XXX')", names)

	// 3 slice literal
	// - like an array w/out length limitation!

	// 3.1 array literal
	q := [3]bool{true, true, false}
	// 3.2 slice
	r := []bool{true, true, false}
	fmt.Println("q, r", q, r)

	// 3.3 slice of type struct
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
	}
	fmt.Println("slice of type struct: ", s)

	// 4 defaults
	// - low bound: 0
	// - high bound: length of slice
	t := []int{2, 3, 5, 7, 11, 13}

	default_low := t[:2]
	fmt.Println("default low: ", default_low)

	default_high := t[1:]
	fmt.Println("default high: ", default_high)

	// 5 length + capacity
	// - length, len(s): number of elems
	// - capacity, cap(s): number of elems in underlying array, starting from first elem in slice
	fmt.Println("Original Slice before PrintSlice()")
	printSlice("t", t)

	// zero length
	t = t[:0]
	fmt.Println("Zero Length...")
	printSlice("t", t)

	// extend length
	t = t[:4]
	fmt.Println("Extend length...")
	printSlice("t", t)

	// drop first 2 values
	t = t[2:]
	fmt.Println("Drop first 2 values...")
	printSlice("t", t)

	// 6 Nil slice
	// - length: 0
	// - capacity: 0
	var nilSlice []int
	printSlice("nil slice", nilSlice)
	// fmt.Printf(
	// 	"Nil Slice: len=%d cap=%d \n",
	// 	len(nil_slice),
	// 	cap(nil_slice),
	// )

	// 7 Init with make
	// use make() for dynamically sized arrays
	u := make([]int, 5) //len=5,cap=5,[0 0 0 0 0]
	printSlice("u", u)

	// specify capacity by passing 3rd argument
	v := make([]int, 0, 5) //len=0,cap=5, []
	printSlice("v", v)

	// v.cap=5, v[:2] [0 1 x2 x3 x4 x5 ]
	// len=2, cap=5
	w := v[:2]
	printSlice("w", w)

	// w.cap=5, w[2:5] [x0 x1 2 3 4 x5]
	//len=3, cap=3
	x := w[2:5] //
	printSlice("x", x)

	// 8 Slices of Slice
	ticTacToeBoard := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// players take turns
	ticTacToeBoard[0][0] = "X"
	ticTacToeBoard[2][2] = "O"
	ticTacToeBoard[1][2] = "X"
	ticTacToeBoard[1][0] = "O"
	ticTacToeBoard[0][2] = "X"

	for i := 0; i < len(ticTacToeBoard); i++ {
		fmt.Printf(
			"%s \n",
			strings.Join(ticTacToeBoard[i], " "))
	}

	// 9 Append to Slice
	// - append(slice, value)
	// - returns slice w/ old+new values
	var appendSlice []int
	printSlice("Empty slice", appendSlice)

	// append 0
	appendSlice = append(appendSlice, 0)
	printSlice("Append 0 to slice", appendSlice)

	// slice grows as needed
	appendSlice = append(appendSlice, 1)
	printSlice("Growing slice", appendSlice)

	// append multiple elems
	appendSlice = append(appendSlice, 2, 3, 4)
	printSlice("Append multiple elems", appendSlice)

	// 10 Range For Loop
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// i=index, v=elem copy @ index
	for i, v := range pow {
		fmt.Printf(
			"Range Loop %d: 2**%d = %d \n",
			v, i, v)
	}

	// 10.1 Skip Index
	// - for _, value := range pow
	// 10.2 Skip Value
	// - for i, _ := range pow
	// - for i := range pow

}

func printSlice(c string, s []int) {
	fmt.Printf(
		"%s len=%d cap=%d %v \n",
		c, len(s), cap(s), s,
	)
	fmt.Println("slice: ", s)
}
