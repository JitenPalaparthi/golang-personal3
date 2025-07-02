package main

import (
	"arena"
	"fmt"
	"math/rand/v2"
	"unsafe"
)

type Employee struct {
	Name     string    // 16
	Salaries []float32 // 24
	Address  string    // 16
	Gender   bool      // 1
	Married  bool
}

// GOEXPERIMENT=arenas go build

func main() {
	// // Create a new arena
	mem := arena.NewArena()
	// // Ensure the arena is freed when we're done
	//defer mem.Free() // This is free by your compilier marking not by your GC

	// // Allocate a 1MB slice in the arena
	const oneMB = 1024 * 1024 // 1 << 20
	slicePtr := arena.MakeSlice[int](mem, oneMB, oneMB)

	// 	sum = ProcessMillionElements()
	// 	println(sum)

	sum := ProcessMillionElementInput(slicePtr)
	println(sum)

	mem.Free() // delete free the memory , not the gc, the application

	// employees := make([]Employee, 100000)
	// mem1 := arena.NewArena()
	// defer mem1.Free()
	// const oneMB1 = 1 << 20
	// sliceEmployee := arena.MakeSlice[Employee](mem1, oneMB1, oneMB1)
	// fmt.Println(sliceEmployee)
	// fmt.Println(employees)

	fmt.Println("Size of Employee:", unsafe.Sizeof(Employee{}))

}

/*

slice := make([]int,1000000)
ptr -->

*/

func ProcessMillionElements() int {
	mem := arena.NewArena()
	// Ensure the arena is freed when we're done
	defer mem.Free() // This is free by your compilier marking not by your GC
	const oneMB = 1 << 20
	bigSlice := arena.MakeSlice[int](mem, oneMB, oneMB)
	//bigSlice := unsafe.Slice(slicePtr, oneMB) //
	println("len", len(bigSlice), "Cap:", cap(bigSlice))
	sum := 0
	for i := range bigSlice {
		bigSlice[i] = rand.IntN(1000)
	}
	for _, v := range bigSlice {
		if v%2 == 0 {
			sum += v
		}
	}
	return sum
}

func ProcessMillionElementInput(slice []int) int {
	sum := 0
	for i := range slice {
		slice[i] = rand.IntN(1000)
	}
	for _, v := range slice {
		if v%2 == 0 {
			sum += v
		}
	}
	return sum
}
