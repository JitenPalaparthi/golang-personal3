package main

func main() {

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}
	println()

	i := 1

	for i <= 10 { // for loop like a while loop
		if i%2 == 0 {
			print(i, " ")
		}
		i++
	}
	println()

	for i := 1; ; i++ {
		if i > 10 {
			break
		}
		if i%2 == 0 {
			continue
		}
		print(i, " ")
	}
	println()

	// for i := 1; ; {
	// 	if i > 10 {
	// 		break
	// 	}
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	print(i, " ")
	// 	i++
	// }
	// println()

	i = 1
	a, b := 0, 1
	for {
		if i > 10 {
			break
		}
		print(a, " ")
		a, b = b, a+b
		i++
	}
	//a, b, c := 10, 20, 30
	// t := a
	// b = a
	// a = t
	//a, b, c = c, a, b

}
