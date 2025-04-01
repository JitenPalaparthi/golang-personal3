package main

const PI = 3.14 // are stored in text/code segment

var Global int = 100

var UGlobal int

func main() {
	var any1 any = uint8(12)
	println(Global)
	println(UGlobal)
	println(PI)
	//num := any1.(uint) // assert it to int
	//println(num)

	// var num2 int = rand.IntN(100)
	// println(num2)

	// fmt.Printf("Address of %p\n", &Global)
	// fmt.Printf("Address of %p\n", &num2)

	// num = 100
	// num = True
	// num = "Hello World"

	switch any1.(type) {
	case uint, int, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
		println("it is a number type")
	default:
		println("it is not a number type")
	}

	// switch v := any1.(type) {
	// case uint, int, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
	// 	println(v * v)
	// default:
	// 	println("it is not a number type")
	// }

	switch v := any1.(type) {
	case uint:
		println(any1.(uint) * any1.(uint))
	case int:
		println(v * v)
	case uint8:
		println(v * v)
	case uint16:
		println(v * v)
	case uint32:
		println(v * v)
	case uint64:
		println(v * v)
	case int8:
		println(v * v)
	case int16:
		println(v * v)
	case int32:
		println(v * v)
	case int64:
		println(v * v)
	case float32:
		println(v * v)
	case float64:
		println(v * v)

	default:
		println("not a number type")

	}

}

// 1000f3810 B runtime.uint64Type
// any Header
// DataPtr
// TypePtr

// Text Segment
// Data Segment
// BSS
// Stack Memory
// Heap Memory
