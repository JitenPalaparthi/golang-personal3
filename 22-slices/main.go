package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	slice1 := make([]int, 10)
	if slice1 == nil {
		print("yes slice is nil")
	}
	fmt.Println(slice1)

	AssignSlice(slice1) // passing it as an argument
	fmt.Println(slice1)
	fmt.Printf("Address of slice1:%p\n", &slice1)
	fmt.Printf("Ptr of slice1:    %p\n", &slice1[0])

	var slice2 []int // the declaration but not instantiation
	// Ptr --> nil
	// Len --> 0
	// Cap --> 0

	if slice2 == nil {
		println("yes slice2 is nil")
	}

	slice3 := []int{} // even there is no length, hence no capasity still the slice is not nil
	// Ptr --> a dummy pointer
	// Len --> 0
	// Cap --> 0

	if slice3 == nil {
		println("yes slice3 is nil")
	}
	fmt.Println("--->", slice3)

	//fmt.Printf("Ptr of slice3:    %p\n", &slice3[0])

	if AssignSlice == nil {
		println("Yes Assign_Slice is nil")
	}
	//fmt.Printf("Ptr of slice3:%p\n", &(AssignSlice))

	var slice4 []int // nil
	slice4 = append(slice4, 10)
	fmt.Println(slice4)

	str := "Hello World 🍎 🧪 中"
	println("Length", len(str))

	// String Header
	// Ptr --> Some address
	// Len --> 25

	var str2 string

	println(str2)

	var any1 any = 10

	// Header of any
	// DataPtr --> nil
	// TypePtr --> nil
	any1 = str

	if any1 == nil {
		println("Yes any1 is nil")
	}

	// arr := [3]int{10, 11, 12} // What is the header of this array

	// Type of the array -> [3]int
	// if arr == nil {

	// }
}

func AssignSlice(slice []int) {
	for i := range slice {
		slice[i] = rand.IntN(1000)
	}
}

/*
Slice Header
------------
Ptr --> 0x123abc123
Len --> 10
Cap --> 10
*/
// zero values of the slice
// nil is very curious in Go

/* String

 */
