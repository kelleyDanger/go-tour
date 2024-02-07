package main

import "fmt"

func main() {
	// 1 pointer
	i, j := 1, 6

	// create pointer to i
	p := &i
	// read i through the pointer
	fmt.Println("i through pointer *p: ", *p)
	fmt.Println("pointer memory address: ", p)

	// set i through the pointer
	*p = 21
	fmt.Println("new i: ", i) // new value of i

	// create pointer to j
	p = &j

	// set j through the pointer
	*p = *p / 2
	fmt.Println("new j: ", j) // new value of j
}
