package main

func main() {

	// no need to give break in go
	switch day := uint8(4); day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("noday")
	}

	day := uint8(4)
	switch {
	case day == 1:
		println("Sunday")
	case day == 2:
		println("Monday")
	case day == 3:
		println("Tuesday")
	case day == 4:
		println("Wednesday")
	case day == 5:
		println("Thursday")
	case day == 6:
		println("Friday")
	case day == 7:
		println("Saturday")
	default:
		println("noday")
	}

	char := 'A'

	switch char {
	case 65, 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		println(string(char), "is a vovel")
	default:
		println(string(char), "is either a consonent or a special char")
	}

	// empty switch

	num := 50

	switch {
	case num >= 0 && num <= 50: // check true or false
		println(num, "is between or eaual 0 and 50")
	case num > 50 && num <= 100:
		println(num, "is between 50 and 100")
	case num > 100:
		println(num, "is more than 100")
	default:
		println(num, "is negative value")
	}

	age, gender := uint8(18), 'F'

	switch {
	case age >= 18 && (gender == 70 || gender == 'f'):
		println("she is eligible for marriage becase age is ", age)
	case age >= 21 && (gender == 'M' || gender == 'm'):
		println("he is eligible for marriage becase age is ", age)
	default:
		println("not eligible for marriage")
	}

	// char1 := 'A'
	// char2 := int32(22000)
	// char3 := char1 + char2
	// char4 := char2 - char1
	// // char5 := uint8('A')
	// // char6 := uint8('喯')
	// // char7 := uint8(21935)

	// println(char3, char4)
	// str1 := "A"
	// str2 := "喯"
	// println(len(str1), str1)
	// println(len(str2), str2)

}
