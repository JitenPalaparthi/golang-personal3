package main

func main() {

	// a, b := 10, 20

	// if a+b/10+(a*b)-(a/b)+20+a*1+b/20 == 100 && true || false {

	// }
	age, gender := uint8(19), 'F'
	//println(age >= 18 && (gender == 'F' || gender == 'f') && true && (true || false))
	if age >= 18 && (gender == 'F' || gender == 'f') && true && (true || false) {
		println("she is eligible for marriage becase of age ", age)
	} else if age >= 21 && (gender == 'M' || gender == 'm') {
		println("he is eligible for marriage becase of age ", age)
	} else {
		println("not eligible for marraige")
	}

	//println(age, gender)

}
