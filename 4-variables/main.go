package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var any1 any
	if any1 == nil { // interface{} -> empty interface
		fmt.Println("Zero value of any1 is", any1, "Type of", reflect.TypeOf(any1)) // nil
	}
	any1 = 100
	fmt.Println(any1, "Type of any1:", reflect.TypeOf(any1))
	any1 = "hello world"
	fmt.Println(any1, "Type of any1:", reflect.TypeOf(any1))
	any1 = true
	fmt.Println(any1, "Type of any1:", reflect.TypeOf(any1))
	any1 = 1231.123
	fmt.Println(any1, "Type of any1:", reflect.TypeOf(any1))
	any1 = uint8(123)
	fmt.Println(any1, "Type of any1:", reflect.TypeOf(any1))
	any1 = rune('A')
	fmt.Println(any1, "Type of any1:", reflect.TypeOf(any1))
	any1 = uintptr(12312312) // will discuss abt it later
	fmt.Println(any1, "Type of any1:", reflect.TypeOf(any1))

	str1 := "Hello World"
	str2 := "Hello Universe"
	var str3 string // Zero ""

	fmt.Println("Size of str1:", unsafe.Sizeof(str1))
	fmt.Println("Size of str2:", unsafe.Sizeof(str2))
	fmt.Println("Size of str3:", unsafe.Sizeof(str3))
	fmt.Println("Size of any1:", unsafe.Sizeof(any1))

}

// what is the zero value of any? nil
// why any is nil?

// Two data structures in Go
// 1. string
// 2. any

// 1. string -->

/*
str1{
ptr: 0x1111 (some address) // the size of any ptr on 64 bit machine is 8 bytes
len: 11 // len is int , the size of int is 8 bytes
}

str2{
ptr: 0x1112312 (some address) // the size of any ptr on 64 bit machine is 8 bytes
len: 14 // len is int , the size of int is 8 bytes
}

str3{
ptr: 0x01 (some address) // the size of any ptr on 64 bit machine is 8 bytes
len: 0 // len is int , the size of int is 8 bytes
}

why a dummy pointer to the str3 when there is no data ?
Go wants string to be very simplem data type
They dont want to complicate string by check nil bla bla bla..
They want string to be as easy as any other defined data type

*/

// when ever there is data struct like any, slice, map , chan, func, interface
// all the above data structures contains a ptr, it always checks wether this poitner is nil or not so
// can check whether the variable is nil or not

/*

any1{
dataPtr: The pointer to the actual data   // 8 bytes
typePtr: The pointer the underlining type // 8 bytes
}

all the types in the system (defined or userdedined are loaded in the system when applications
starts running. typePtr is the poiner of that particular data type)
*/
