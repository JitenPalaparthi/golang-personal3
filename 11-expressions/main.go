package main

import (
	"fmt"
	"math/big"
)

func main() {
	// swap variables
	a, b := 100, 200
	t := a
	a = b
	b = t

	// swap values
	c, d := 10, 20
	c, d = d, c

	println(c, d)

	x, y, z := 10, 20, 30
	x, y, z = y, z, x
	println(x, y, z)

	l, j := 9, 21

	l++
	j--
	// l += 1
	// j -= 1
	// l *= 1
	// j /= 1

	k := l + j*2 - l + j + 30 - l + 20 + (l * 2) + (j / 10)
	// (l * 2) -> 20
	// (j/10) -> 2
	// l + j*2 - l + j + 30 - l + 20 + 20 + 2
	// l+ 40 - l + j + 30 - l + 20 + 20 + 2
	// 50 - 10 + 20 + 30 -10 + 42
	// 40 + 40 + 42
	// 122
	println("k:=", k)

	// arithemetic
	// comparision > < == >= <= !=
	// logical && || !=

	//k = (l + j*2 - l + j + 30 - l + 20 + (l * 2) + (j / 10))

	// k1 := {
	// 	l + j*2 - l + j + 30 - l + 20 + (l * 2) + (j / 10)
	// }

	// k1 := func() int {
	// 	return l + j*2 - l + j + 30 - l + 20 + (l * 2) + (j / 10)
	// }

	a3, b3 := 999999999999999999, 999999999999999999
	c3 := a3 + b3
	println(c3)

	a4 := big.NewInt(int64(a3))
	b4 := big.NewInt(int64(b3))
	c4 := new(big.Int).Add(a4, b4)
	fmt.Println(c4)

	// c4 = a4 + b4

	//strct3:= struct1+struct2

	// ok := true + false

	// str1 := "hello"
	// str2 := "world"
	// str3 := str1 + str2

	// str4 := str1 - str2

	println("Big int")
	{
		a := new(big.Int)
		b := new(big.Int)
		a.SetString("111111111111111111111111111111", 10)
		//a1 := 111111111111111111111111111111
		b.SetInt64(1111111111)
		sum := new(big.Int).Add(a, b)
		fmt.Println(sum)
	}

}

// func add1(a, b uint8) uint16 {
// 	return 0
// }

// func add2(a, b int64) int64 {
// 	return 0
// }
