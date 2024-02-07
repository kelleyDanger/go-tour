package main

import "fmt"

// Arrays have fixed size.

func main() {
	// 1 initialize array
	// var a[n]T -> array 'a' of 'n' values of type 'T'
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 2 quick init and assign
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
