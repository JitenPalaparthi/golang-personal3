package main

func main() {

	str1 := " Nice to meet you" // str1 is a collection of bytes
	str2 := "很高兴认识你"            // Nice to meet you

	println(str2)

	for i := 0; i < len(str1); i++ {
		print(string(str1[i]), " ")
	}
	println()

	for i := 0; i < len(str2); i++ {
		print(string(str2[i]), " ")
	}
	println()

	for i, v := range str2 {
		println("i:", i, "char:", string(v))
	}

	for i := range 10 { // 1.22 onwards
		println(i)
	}

}
