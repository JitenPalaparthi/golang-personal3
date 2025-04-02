package main

import (
	"fmt"
	"reflect"
)

var (
	garr1 [5]int               // BSS(uninitialized) which is a part of the data segment
	garr2 = [3]int{10, 11, 12} // datasegment which is a part of the data segment
)

const (
	PI = 3.14
)

func main() {

	var arr1 [3]int
	// what is the type of this array? [3]int
	// what are the zero value of this array? [0 0 0]
	fmt.Println("arr1:", arr1, "Type of arr1:", reflect.TypeOf(arr1))

	for i := 0; i < len(arr1); i++ {
		print(arr1[i], " ")
	}
	println()
	arr1[0], arr1[1], arr1[2] = 100, 200, 300
	for _, v := range arr1 {
		print(v, " ")
	}
	println()

	arr2 := [5]int{1, 2, 3, 4, 5} // shorthand declaraion
	arr3 := [...]int{10, 20, 30}  // the length is evaluated at compile time
	fmt.Println(arr2, arr3)

	// rintln(garr1)
	// fmt.Println(garr2)
	fmt.Println(PI)
	println(&garr1[0], 100198780)
	println(&garr2[0], 100168530)
	var num int = 100
	fmt.Printf("%x\n", &num)

	s1 := SumOf(arr1)
	//s2 := SumOf([3]int(arr2))
	fmt.Println("Sum of arr1:", s1)
	fmt.Println("address of arr1:", &arr1[0])

	var arr4 [999999]int
	println("address of arr4:", &arr4[0])

	arr2d := [2][2]int{{1, 2}, {3, 4}}
	arr3d := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	println("arr2d")

	for _, a1 := range arr2d {
		for _, v := range a1 {
			print(v, " ")
		}
		println()
	}

	println("arr3d")

	for _, a1 := range arr3d {
		for _, a2 := range a1 {
			for _, v := range a2 {
				print(v, " ")
			}
		}
		println()
	}

}

//100198780 B main.garr1
//100168530 D main.garr2

func SumOf(arr [3]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
