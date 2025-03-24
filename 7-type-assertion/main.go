package main

import "fmt"

func main() {

	var str1 string = "Hello Wrold"
	println(str1)

	var any1 any
	if any1 == nil {
		println("any1 is nil")
	} else {
		fmt.Println(any1)
	}
	any1 = 12312
	//var num1 int = int(any1) // cannot do type casting, neither conversion

	var num1 int = any1.(int) // type assertion any.(type)
	println(num1)

	any1 = true
	//var num2 int = any1.(int) // type assertion any.(type)
	//println(num2)
	any1 = 1231.12312
	any1 = "Golang Training"
	any1 = any("Something")

}

/*
String header for str1
-----------------------
Ptr --> Pointer to the actual data
Len --> 11
*/

/*
any header for any1
-------------------
TypePtr   // 8 bytes on 64bit machine
DataPtr   // 8 bytes on 64bit machine
*/
