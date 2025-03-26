package main

import "fmt"

func main() {

	age := uint8(8)

	if age >= 18 {
		println("Eligible for vote, becase of age is ", age)
		fmt.Printf("Address of age:%p\n", &age)
	} else {
		println("Not eligible for vote, becase is age is ", age)
		fmt.Printf("Address of age:%p\n", &age)
	}

	if age := uint8(28); age >= 18 { // This is a new variable, which has the scope of the entire if and else
		println("Eligible for vote, becase of age is ", age)
		fmt.Printf("Address of age:%p\n", &age)
	} else {
		println("Not eligible for vote, becase is age is ", age)
		fmt.Printf("Address of age:%p\n", &age)
	}

	// a := 100
	// {
	// 	a := 200
	// 	println(a)
	// }
	// println(a)

}
