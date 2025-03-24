package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var num1 int = 12323

	var str1 string = string(num1) // type caste, based on the value it converts he unicode char

	fmt.Println(str1)

	var str2 = strconv.Itoa(num1) // type conversion .. as it is

	fmt.Println("str2:", str2, "type of str2:", reflect.TypeOf(str2))

	var str4 = fmt.Sprint(num1, true, "tring to convert ", 12321.12312)
	fmt.Println("str4:", str4, "type of str4:", reflect.TypeOf(str4))

	str5 := fmt.Sprintf("Bin:%b Xex:0X%X", num1, num1)

	fmt.Println("str5:", str5, "type of str5:", reflect.TypeOf(str5))

	a1, p1 := Rect(1.5, 2.6)
	fmt.Println("Area:", a1, "Perimeter:", p1)
	a2, s2, m2, d2 := Calc(12.4, 15.6)
	fmt.Println("Add:", a2, "Sub:", s2, "Mul:", m2, "Div:", d2)

	_, s3, _, _ := Calc(12.4, 15.6) // _ blank identifier
	fmt.Println("Sub:", s3)

	// type conversion

	str6, str7, str8 := "123213", "1312.23", "b"

	println(str6, str7, str8)

	num2, err := strconv.Atoi(str6)

	if err != nil { // there is some error
		println("error--->", err.Error())
	} else {
		println(num2)
	}

	float1, err := strconv.ParseFloat(str7, 32)
	if err != nil {
		println("error--->", err.Error())
	} else {
		fmt.Printf("%.2f\n", float1)
	}

	ok1, err := strconv.ParseBool(str8) //1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False
	if err != nil {
		println("error--->", err.Error())
	} else {
		fmt.Println(ok1)
	}

	str9 := "fffff"
	num3, err := strconv.ParseInt(str9, 16, 32)
	if err != nil {
		println("error--->", err.Error())
	} else {
		fmt.Println(num3)
	}

	str10 := "999999"

	num4, _ := strconv.Atoi(str10) // blank identifier
	fmt.Println(num4)

}

func Rect(l float32, b float32) (float32, float32) {
	return l * b, 2 * (l + b)
}

func Calc(i, j float32) (a float32, s float32, m float32, d float32) {
	a = i + j // not creating new variables
	s = i - j
	m = i * j
	d = i / j
	return a, s, m, d
	//return i + j, i - j, i * j, i / j
}
