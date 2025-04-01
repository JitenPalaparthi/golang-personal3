package main

import "math/rand/v2"

func main() {

	// done := false

	// for i := 1; i <= 5; i++ {
	// 	if done {
	// 		break
	// 	}
	// 	for j := 2; j <= 5; j++ {
	// 		if i == j {
	// 			done = true
	// 			break
	// 		}
	// 		println("i: ", i, "j: ", j)

	// 	}
	// 	//println()
	// }
	// println()

outer:
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 5; j++ {
			if i == j {
				break outer
			}
			println("i: ", i, "j: ", j)
		}
		//println()
	}
	println()

outer2:
	for {
		num := rand.IntN(10000)
		switch {
		case num%2 == 1:
			println(num, "is odd number")
		case num%2 == 0:
			println(num, "is even number")
			// 	fallthrough
			// case num%10 == 0:
			if num%10 == 0 {
				println(num, "is divisible by 10 so exiting ")
				break outer2
			}
		}
	}

	// 	num := 1
	// loop:
	// 	if num%2 == 0 {
	// 		print(num, " ")
	// 	}
	// 	num++
	// 	if num <= 10 {
	// 		goto loop
	// 	}
	// 	println()

	count := 0

loop1:
	num := rand.IntN(1000)
	count++
	if count <= 10 {
		if num%2 == 0 {
			goto even
		} else {
			goto odd
		}
	} else {
		goto outer3
	}
even:
	println(num, " is even ")
	goto loop1
odd:
	println(num, " is odd ")
	goto loop1
outer3:
	println("exit")

}
