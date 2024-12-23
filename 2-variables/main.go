package main

import "fmt"

func main() {

	var num1 int = 100
	var ok1 bool = true
	var str1 string = "Hello World"

	var num2 = 123       // int
	var ok2 = true       // bool
	var str2 = "GO Lang" // string

	var (
		num3       = 123.45 // float64
		age  uint8 = 39     // when dont give type automatically assigns
	)

	age = 69 // mutate or reassign a value to the variable

	// var char1 rune = 'A'
	// var char2 int32 = 'B'
	// var char3 uint8 = 'C'
	// var char4 byte = 'D'
	// var char5 rune = '你'
	// var char6 = 'A' + 'B'
	// var char7 uint64 = '你'

	var (
		char1 rune   = 'A'
		char2 int32  = 'B'
		char3 uint8  = 'C'
		char4 byte   = 'D'
		char5 rune   = '你'
		char6        = 'A' + 'B'
		char7 uint64 = '你'
	)
	char6 = char6 + 'C' // mutating a variable with a new value

	//var char8 bool = 'A'

	println(num1, num2, num3, ok1, ok2, str1, str2, age)
	fmt.Println(num1, num2, num3, ok1, ok2, str1, str2, age)
	fmt.Println(char1, char2, char3, char4, char5, char6, char7)
	num4 := 1231232131 // shorthand declaration
	println(num4)

	const PI float32 = 3.14

	const (
		MAX = 9999 // shorthand delcartion for constants
		MIN = 1111
	)
	//PI = 3.143

	fmt.Println(PI, MAX, MIN)

	str3 := "Hello World" // short hand notation ..
	fmt.Println(str3)

	a, b := 100.123, 200 // a is float64 and b is int

	fmt.Println(a, b)

	a1, b1 := 10, 20
	// swapping no need of temp variable in Go
	t1 := a1
	a1 = b1
	b1 = t1

	a1, b1 = b1, a1 // easy swapping
	fmt.Println("After swapping twice,", a1, b1)
	a2, b2, c2 := 10, 20, 30
	fmt.Println("Before swapping a1,b1,c1", a2, b2, c2)
	a2, b2, c2 = b2, c2, a2
	fmt.Println("After swapping a1,b1,c1", a2, b2, c2)

}

// defined/primitive/builtin types
// numbers
// 	int,uint,int8,uint8,int16,uint16,
//  int,uint32,int64,uint64,float32,float64
// 	rune, byte --> these two are not a new types but just alias of int32 and uint8
// boolean -> bool
// strings -> string

// 1. Zero value
// 2. Type inference
// 3. Shorthand notation of creating a variable
// all the above done by the compiler so it is at compile time

// fmt is a pacakge which is a standard package
