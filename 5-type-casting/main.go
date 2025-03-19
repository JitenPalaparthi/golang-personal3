package main

import (
	"fmt"
	"reflect"
)

type char = int32 // alias type int32 to char
// / int32 == rune == char
func main() {

	var num1 uint8 = 123

	var num2 uint64 = uint64(num1)

	var num3 = float32(12312.123) // type inferance

	char1 := 'A' + 10

	fmt.Println("Char1:", char1, "Type of char1:", reflect.TypeOf(char1))
	fmt.Printf("Char1:%v type of char1:%T\n", char1, char1)

	fmt.Println(num1, num2, num3)

	var char2 char = 'B'
	fmt.Printf("char2:%v type of char2:%T\n", char2, char2)

	var byte1 byte = 20
	fmt.Printf("byte1:%v type of byte1:%T\n", byte1, byte1)

	//char1 = 19000

	fmt.Println(string(char1))
	fmt.Println(string(int32(rune(char2))))

	// 0 - 65000

	// char3 := uint8('B')
	// char4 := uint16('䨸')
	// char5 := uint32('B')
	// char6 := uint64('B')

	// any number type can be type casted to another number type

	num4 := 1231.1231
	var num5 int32 = int32(num4)
	println(num5)
	// var ok1 bool = true

	// var num6 int8 = int8(ok1)

	//num6 := 999999999999999999 // 1101111000001011011010110011101001110110001111111111 11111111
	//println(string(num6))

	// �

	//var num7 uint8 = uint8(num6)

	//fmt.Printf("binary of num6:%b", num6)
	//fmt.Println(string(255))

	var num6 uint64 = 12345678     // 10111100 0110000101001110
	var num7 uint8 = uint8(num6)   // 78
	var num8 uint16 = uint16(num6) // 24910

	fmt.Println(num7)
	fmt.Println(num8)

	// var float1 float32 = 3.14 // SEEEEEEE EMMMMMMM MMMMMMMM MMMMMMMM / S E M
	// fmt.Printf("%b", float1)

	str2 := fmt.Sprint(num6)
	fmt.Println(str2)

}

// comverting or casting from one type to the other type
// type casting
// type conversion
// type assertion
