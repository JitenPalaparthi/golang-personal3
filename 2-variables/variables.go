package main

import (
	"fmt"
	"time"
)

func main() { // main is a function

	var num1 int = 1231312

	var num2 uint8 = 255

	var float1 float32 = 1231.2312

	var ok1 bool = true

	var str1 string = "Hello World!" // collection of bytes

	var c1 complex64 = 123.2 + 4.0i // very very rare use

	var char1 rune = 'A'

	var char2 rune = 22000

	// Group of variables
	var (
		num3 uint32 = 2123
		num4 uint   = 1212312
		ok2  bool   = false
		str2 string = "Hello Universe!"
	)

	var (
		num5 int
		ok3  bool
		str3 string
	)

	fmt.Println(num1, num2, num3, num4, num5, ok1, ok2, ok3, float1, str1, str2, str3, char1, char2, c1)
	fmt.Println("time right now-->", time.Now())
	fmt.Println("Global variables")
	fmt.Println(g)

	println("constants")

	println(PI, MAX, MIN)

}

// every binary project must have a file that has pacakge main and also must have a func called main

// Constants --> Value cannot be changed, immutable --> once assigned that is same throughout the program
// Local Variables --> Variables which are local for the func
// Global Variables --> Variables which are global , can be used globally also called them as package level variables

// types of the variables, defined variable types

// Numbers
// int,int8,int16,int32,int64
// uint,uint8,uint16,uint32,uint64
// rune, byte (not a new type but a different name that is alias of int32, uint8 respectively)
// float32, float64
// complex64, complex128 --> real number + imag number
// uintptr --> int type specially used to store addresses

// Boolean --> bool
// String --> string
// Any --> interface{} or any

// unicode chars utf8

// zero value --> a default value is given if no value is assigned to the variable
// all number types the zero value is 0
// bool the zero value is false
// string the zero value is empty string ""
