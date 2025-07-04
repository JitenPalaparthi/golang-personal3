package main

import "fmt"

func main() {

	slice1 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// Len - 11 and Cap - 11
	fmt.Println(slice1)

	fmt.Printf("Address of the slice1:%p\n", &slice1)
	fmt.Printf("Pointer of the slice1:%p\n", &slice1[0])

	println()
	MultiplySliceBy(slice1, 2) //m Ptr, Len and Cap
	fmt.Println(slice1)
	println()
	MultiplyAndAppendSliceBy(slice1, 2)
	fmt.Println(slice1)
	println()

	// [40 44 48 52 56 60 64 68 72 76 80 --append 12, 13, 14, 15, 16, 17, 18, 19, 20]
	slice1 = MultiplyAndAppendSliceBy(slice1, 2, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	fmt.Println(slice1)

	// num := 10 // what is the data ? -- 10
	// Incr(num) // 10 is copied to a new variable for func
	// println(num)

}

func Incr(num int) {
	num++
	println("Inside func num", num)
}

func MultiplySliceBy(slice []int, multiplier int) {
	fmt.Printf("Address of the slice: %p\n", &slice)
	fmt.Printf("Pointer of the slice1:%p\n", &slice[0])

	for i, v := range slice {
		slice[i] = v * 2
	}
}

// variadic parameter
// variadic parameter cannot be used as a normal variable, can only be used in functions or methods
// variadic must be the last parameter
func MultiplyAndAppendSliceBy(slice []int, multiplier int, nums ...int) []int {
	fmt.Printf("Address of the slice: %p\n", &slice)
	fmt.Printf("Pointer of the slice1:%p\n", &slice[0])

	// what is the capasity --? 11
	slice = append(slice, nums...)
	// Len-->  11 + 9 --> 20 , does it have the capasity to keep 20 elements? no
	// Cap algorthim kicks and them it reinitialize the slice , so the whole header gets changed
	fmt.Printf("Pointer of the slice1 after append:%p\n", &slice[0])
	for i, v := range slice {
		slice[i] = v * 2
	}
	return slice
}

/* Slice Header */
// What would happen to the slice(header) when ever a slice is appended ?
// Either the length is changed or all 3 of them gets changed
