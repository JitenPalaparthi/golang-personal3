package main

var (
	globalSlice1 = make([]int, 100)
	globalslice2 = []int{}
)

func main() {

	println(globalSlice1)
	println(globalslice2)
	var slice1 []int // declaration of a slice

	if slice1 == nil {
		println("nil slice")
	}

	slice1 = make([]int, 5) // This is going to change the slice header
	//fmt.Println(slice1)     //  [0 0 0 0 0]
	println(slice1)

	slice2 := make([]int, 5)    // shorthand declaration
	slice3 := []int{10, 20, 30} // slice declaraion and instantiation with values
	println(slice2)
	println(slice3)
	{
		slice4 := []int{} // the ptr is not nil
		if slice4 == nil {
			println("slice4 is nil")
		} else {
			println(slice4)
		}
	}
	{
		slice5 := make([]int, 9999)
		println(slice5)
		{
			// var str1 string
			//	slice5 := make([]int, 9999)
			//	println(slice5)
			sum := SumOf(slice5) // call or pass by value
			println("Sumof slice5:", sum)
		}

		{
			// var str1 string
			//slice5 := make([]int, 9999)
			//	println(slice5)
			sum := SumOf(slice5) // call or pass by value
			println("Sumof slice5:", sum)
		}
	}

	var num int
	println(&num)

	{
		a, b := 10, 20
		{
			c := a + b
			println(c)
		}
		// println(c)
	}
}

// a slice can be nil

// Slice header
// ------------
// Ptr --> 8 bytes --> nil
// Len --> 8 bytes --> 0
// Cap --> 8 bytes --> 0

// slice header with make

// Slice header
// ------------
// Ptr --> 8 bytes --> 0X400f100 (some address)
// Len --> 8 bytes --> 5
// Cap --> 8 bytes --> 5

func SumOf(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
