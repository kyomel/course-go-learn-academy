package main

import "fmt"

// basic types(numbers, strings, booleans)
// var myInt int
// var myInt16 int16
// var myInt32 int32
// var myInt64 int64
// var myUint uint
// var myFloat32 float32
// var myFloat64 float64

// aggregate types (array, struct)

// reference types (pointers, slices, maps, functions, channels)

// interfaces type

// type myStruct struct {
// 	NumberOfTires int
// 	Luxury        bool
// 	BucketSeats   bool
// 	Make          string
// 	Model         string
// 	Year          int
// }

func main() {
	// basic types
	// myInt = 10
	// myUint = 20
	// myFloat32 = 10.1
	// myFloat64 = 100.1

	// log.Println(myInt, myUint, myFloat32, myFloat64)

	// myString := "Michael"

	// log.Println(myString)

	// myString = "John"

	// var myBool bool = true
	// myBool = false
	// log.Println(myBool)

	// aggregate types
	// var myStrings [3]string
	// var myStrings [3]int
	// myStrings[0] = "cat"
	// myStrings[1] = "dog"
	// myStrings[2] = "hamster"
	// myStrings[0] = 1
	// myStrings[1] = 5
	// myStrings[2] = 7
	// fmt.Println("First element in array is", myStrings[0])
	// var myCar myStruct //Old struct
	// myCar.NumberOfTires = 4
	// myCar.Luxury = false
	// myCar.Make = "Volkswagen"

	// reference types
	x := 10
	myFirstPointer := &x
	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)
	*myFirstPointer = 15
	fmt.Println("x is now", x)
	changeValueOfPointer(&x)
	fmt.Println("After function call, x is now", x)
}

func changeValueOfPointer(num *int) {
	*num = 25
}
