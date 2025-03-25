package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"strconv"
)

func main() {
	var any1 any
	any1 = rand.IntN(1000)
	fmt.Println("Data ", any1, "Type:", reflect.TypeOf(any1))
	fmt.Printf("Data: %d Type: %T\n", any1, any1)
	var num1 int = any1.(int)
	println(num1)

	any1 = "Hello World!"

	var num2 int

	num2, ok := any1.(int) // type assert to int when the type is string
	if ok {
		println(num2)
	} else {
		println("the type of any1 is not int")
	}

	var (
		num3     = int8(100)
		num4     = float32(123.123)
		num5     = uint16(2312)
		num6     = uint64(1232131)
		num7     = int32(34534)
		num8     = rune(22000)
		num9     = 787678.89
		str2     = "12313"
		str3     = "433.34"
		str4     = "98754_abc" // this cannot be converted
		any2 any = 57412.23
		any3 any = 4344
		ok1      = true // if true add 1 else add 0 (do nothing)
	)
	// 2153382.583

	var sum float64

	sum = float64(num3) + float64(num4) + float64(num5) + float64(num6) + float64(num7) + float64(num8) + num9

	if ok1 {
		sum += float64(1)
	}

	if num, err := strconv.Atoi(str2); err == nil { // when there is no error
		sum += float64(num)
	} else {
		if float1, err := strconv.ParseFloat(str2, 64); err == nil {
			sum += float1
		}
	}

	if num, err := strconv.Atoi(str3); err == nil { // when there is no error
		sum += float64(num)
	} else {
		if float1, err := strconv.ParseFloat(str3, 64); err == nil {
			sum += float1
		}
	}

	if num, err := strconv.Atoi(str4); err == nil { // when there is no error
		sum += float64(num)
	} else {
		if float1, err := strconv.ParseFloat(str4, 64); err == nil {
			sum += float1
		}
	}

	if num, ok := any2.(int); ok {
		sum += float64(num)
	} else {
		if num, ok := any2.(float64); ok {
			sum += num
		}
	}

	if num, ok := any3.(int); ok {
		sum += float64(num)
	} else {
		if num, ok := any3.(float64); ok {
			sum += num
		}
	}

	fmt.Printf("Sum of all variables:%.3f\n", sum)

	// var str5 = "123.23"

	// if num, err := strconv.Atoi(str5); err == nil {
	// 	println("The value of str5:", num)
	// } else {
	// 	if num, err := strconv.ParseFloat(str5, 64); err == nil {
	// 		fmt.Printf("The value of str5:%.3f\n", num)
	// 	}
	// }

	//str := "Hello World"

	println("End of main func")
}

// internal structure of any1
// any Header
// DataPtr --> Ptr to the data
// TypePtr --> Ptr to the type

// str StringHeader
// ptr to the array
// Len of a array
