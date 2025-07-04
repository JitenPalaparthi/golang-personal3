package main

import "fmt"

var Global float64 = 123123.1231

func main() {

	fmt.Printf("Address of Global %p\n", &Global)
	slice1 := []int{}
	fmt.Printf("Address of Slice1 %p\n", &slice1)
	slice2 := make([]int, 1000000000000)
	fmt.Printf("Address of Slice2 %p\n", &slice2)
	fmt.Printf("Ptr of Slice2     %p\n", &slice2[0])
	// Ptr -> Some address
	// Len -> 0
	// Cap -> 0
	slice1 = append(slice1, 10)
	fmt.Println("Len", len(slice1), "Cap", cap(slice1))
	slice1 = append(slice1, 11)
	fmt.Println("Len", len(slice1), "Cap", cap(slice1))
	slice1 = append(slice1, 12)
	fmt.Println("Len", len(slice1), "Cap", cap(slice1))
	slice1 = append(slice1, 13)
	fmt.Println("Len", len(slice1), "Cap", cap(slice1))
	slice1 = append(slice1, 14)
	fmt.Println("Len", len(slice1), "Cap", cap(slice1))

	str1 := "Hello World" // Where is this string stored
	fmt.Printf("Address of str1 %p\n", &str1)
	// str1 = "Hello Universe"

	str1 = str1 + "! How are you doing"

	for i := 1; i <= 10; i++ {
		str1 = str1 + fmt.Sprint(i)
	}
	fmt.Println(str1)

	{
		str2 := "Hello Earth!"
		println(str2)

		str3 := "Hello World"
		println(str3)
	}

	Global = 99933434.343343

	var str3 string
	fmt.Printf("Address of str3 %p\n", &str3)

}
