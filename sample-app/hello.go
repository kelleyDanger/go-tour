package main

// Import packages
// Convention: package name is same as last element in import path name
import (
	"fmt"
	"math/cmplx"
	"math/rand"

	// 1
	"math"
)

// 2 function 'add' takes 2 params of type int
// -  type comes after var name
// -  when 2+ consecutite named function params share a type, you can combine type to last param
//   - x int, y int -> x, y int
func add(x, y int) int {
	return x + y
}

// 3 func can return multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// 4 return values can be named, defined at top of function
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// "naked return": returns named return values
	return
}

// 5 var declares list of variables
// - can be at package or function level
var cute, hot, attractive bool

func list_vars() {
	var i int
	fmt.Println(i, cute, hot, attractive)
}

// 6 initialize vars
var i, j int = 1, 2

// 7 short variable declarations
// inside a func, use ':=' to short assign var w/implicit type
func short_assign() {
	k := 3
	fmt.Println(k)
}

// 8 basic types
var (
	ToBe     bool       = false
	MaxInt   uint64     = 1<<64 - 1
	z        complex128 = cmplx.Sqrt(-5 + 12i)
	NickName string     = "Danger"

	// sized integers
	// 32-bits wide on 32-bit-systems and 64-bits on 64-bit-systems
	// int8  int16  int32  int64
	//uint uint8 uint16 uint32 uint64 uintptr

	// byte (alias for unit8)
	// rune (alias for int32)
	// float32 float64
	// complex64 complex128
)

// 9 Default, Zero Values
// numeric: 0
// boolean: False
// string: "" (empty)
var d int
var e float64
var f bool
var g string

// 10 conversion
var i_1 int = 42
var i_2 float64 = float64(i_1)
var i_3 uint = uint(i_2)

// 11 type inference
func type_interface() {
	type_interface_int := 42
	type_interface_float64 := 3.142
	type_interface_complex128 := 0.867 + 0.5i

	fmt.Println(
		type_interface_int,
		type_interface_float64, type_interface_complex128)
}

// 12 constants
const Pi = 3.14

func test_constants() {
	const Truth = true
	if Truth {
		fmt.Println("Pi Value: ", Pi)
	}
}

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("My favorite number is ", rand.Intn(10))

	// 1 In Go, a name is exported if it begins with a capital letter. (example: math.Pi)
	fmt.Println(math.Pi)

	// 2 call custom func
	fmt.Println(add(1, 2))

	// 3 multiple results
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// 4 naked return, named return values
	x, y := split(17)
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println(split(17))

	// 5 variables, package + function
	list_vars()

	// 6 initialize vars
	fmt.Println(i, j)

	// 7 short assign ':='
	short_assign()

	// 8 basic types
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", NickName, NickName)

	// 9 default, zero values
	fmt.Printf("%v %v %v %q\n", d, e, f, g)

	// 10 convert types
	fmt.Println(i_1, i_2, i_3)

	// 11 type interface
	type_interface()

	// 12 constants
	test_constants()
}
