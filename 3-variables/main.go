package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num1 uint8 = 255 // 1 byte

	var num2 uint64 = uint64(num1) // 8 bytes there could be implicit type casting in some programming languages

	println(num2)

	var num3 int = 12312312312
	var num4 uint8 = uint8(num3)
	var num5 uint16 = uint16(num3)
	fmt.Printf("Binary of num3:  %b %d\n", num3, num3)
	fmt.Printf("Binary of num4:  %b %d\n", num4, num4)
	fmt.Printf("Binary of num5:  %b %d\n", num5, num5)

	//ok1 := true // short hand declaration

	//var num6 uint8 = uint8(ok1) // this cannot be done

	char1 := 'A' // int32 also called as rune
	char2 := 97  // int which is 8 bytes on 64bit machine
	fmt.Println(string(char1), string(char2))

	// in go functions or methods can return multiple values
	s1, c1 := getSqCb(10)
	fmt.Println("Square s1:", s1, "Cube c1:", c1)

	s2, _ := getSqCb(11) // _ blank identifier
	fmt.Println("Square s2:", s2)

	_, c2 := getSqCb(21)
	fmt.Println("Cube c2:", c2)

	//type conversion
	str2 := "12312"

	num7, err := strconv.Atoi(str2)

	if err != nil {
		println(err.Error())
	} else {
		println(num7)
	}

	str3 := "12f312"

	//strconv.ParseInt()
	num8, err := strconv.Atoi(str3)

	if err != nil {
		println(err.Error())
	} else {
		println(num8)
	}

	str4 := "12312.12312"

	float1, err := strconv.ParseFloat(str4, 64)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(float1)
	}

	//

}

func getSqCb(num int) (s int, c int) {
	s = num * num
	c = s * num
	return s, c
}

// type casting and type conversion
// 1. all numbers can be type casted between
// 2. there is no implicit type casting in Go
// 3.
