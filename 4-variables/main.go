package main

import (
	"fmt"
	"reflect"
)

func main() {
	var any1 any
	var str1 string
	var num1 int
	fmt.Println(any1, str1, num1)
	any1 = 100 // int
	fmt.Println(any1)
	any1 = true // bool
	fmt.Println(any1)
	str1 = "Hello World"
	any1 = str1 // string
	fmt.Println(any1)

	c1 := complex(12.45, 1.4) // complex is a builtin function
	any1 = c1
	var r1, imag1 float32 = 12.45, 1.4
	c2 := complex(r1, imag1)
	any1 = c2 // complex64

	// can print the type of a variable

	fmt.Println("Value of num1:", num1, "type of num1:", reflect.TypeOf(num1))
	fmt.Println("Value of any1:", any1, "type of any1:", reflect.TypeOf(any1))

	fmt.Printf("Value of any1:%v type of any1:%T\n", any1, any1)
	fmt.Printf("Value of num1:%d type of num1:%T\n", num1, num1)
	fmt.Printf("Value of str1:%s type of str1:%T\n", str1, str1)
	num1 = 213134
	fmt.Printf("Value of num1:%b type of num1:%T\n", num1, num1)
	fmt.Printf("Value of num1:0X%X type of num1:%T\n", num1, num1)

	num2 := 1231231.1231232
	fmt.Printf("Value of num2:%.3f type of num2:%T\n", num2, num2)

}

// any is a special type of a variable
// any or interface{}

// what is teh zero value of any1, str1 and num1
// num1 zero value is 0
// str1 zero value is ""
// any zero value is nil
