package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

var G float32 = 3.14 // data segment

var R int = rand.IntN(232132) // evaluated even before main

var G2 float32 = G * G // This is an expression

func init() { // init is a special function
	fmt.Println("The random number-->", R)
}

func init() { // init is a special function
	fmt.Println("Value of G", G)
}

func init() { // init is a special function
	fmt.Println("Calling init again-3.Value of G2-->", G2)
}
func init() { // init is a special function
	fmt.Println("Calling init again-4")
}

func main() {
	println("start of main")
	Greet()
	str1 := "Hello World"             // 11
	str2 := "It is ☔, I am drining ☕" // 27
	str3 := "你好，世界"                   // 15

	// byte1 := byte(65)
	// byte2 := uint8(65)

	println("Length of str1:", len(str1), str1)
	println("Length of str2:", len(str2), str2)
	println("Length of str3:", len(str3), str3)

	str3 = "你好，世界-> You know this is Chinees Hello World"
	str1 = "hello world" // still the legnth is same, it reallocate the underlining array

	str4 := strings.Clone(str1) // This clones the entire string to a new variable
	str5 := str1                // This does not clone .. the headers of str1 are copied to str5
	println(str4, "\n", str5)

	var any1 any
	if any1 == nil {
		println("Yes any1 is nil")
	}

	var str6 string // even no byte array is there to the ptr of string header still
	// a dummy ptr is assigned

	// if str6 == nil { // cannot check nil on string
	// 	println("yes str6 is nil")
	// }
	if str6 == "" {
		println("yes str6 is empty")
	}
	//println(str5)
}

func Greet() {
	println("Hello World")
}

// String is not considered as array of chars. String is array of bytes
