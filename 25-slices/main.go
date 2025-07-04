package main

import "fmt"

func main() {

	slice1 := make([]int, 3, 5)
	slice1[0], slice1[1], slice1[2] = 10, 11, 12

	// fmt.Printf("Address of the slice1: %p\n", &slice1)

	PrintSliceHeader(slice1, "slice1")
	slice2 := slice1 // what exactly happen

	PrintSliceHeader(slice2, "slice2")
	// a := 10
	// b := a // the value of a is copied to b

	slice2[0] = 111111 // what is changed, slice2[0]
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)

	slice2 = append(slice2, 13) // :0x1400012c000, len 4 ,cap 5
	slice1[1] = 22222
	fmt.Println("slice1:", slice1) // :0x1400012c000, len 3 ,cap 5
	fmt.Println("slice2:", slice2)

	slice2 = append(slice2, 14, 15) // beyond the original capasity -> it is reallocated,so slice1 and slice2 are divergent
	slice2[2] = 33333

	slice1 = slice2 // what happen? the headers of slice2 are copied to slice1

	fmt.Println("slice1:", slice1) // :0x1400012c000, len 3 ,cap 5
	fmt.Println("slice2:", slice2)
	PrintSliceHeader(slice1, "slice1")
	PrintSliceHeader(slice2, "slice2")

}

func PrintSliceHeader(slice []int, name string) {
	fmt.Printf("Address of the slice: %s: %p\n", name, &slice)
	fmt.Println("Slice", name, ":", slice)
	if len(slice) > 0 {
		fmt.Printf("Pointer of the slice %s :%p\n", name, &slice[0])
	}
	fmt.Printf("Len of the slice %s :%d Capasity of slice %s: %d\n", name, len(slice), name, cap(slice))
	println()
}
