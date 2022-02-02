package main

func main() {
	// answer := 7 + 3*4
	// fmt.Println("Answer:", answer)
	// answer = (7 + 3) * 4
	// fmt.Println("Answer now:", answer)
	// multiplication
	// area = 𝛑r2
	// var radius = 12.0
	// area := math.Pi * radius * radius
	// fmt.Println("Area is", area)
	// integer division
	// half := 1 / 2
	// fmt.Println("Half with integer division", half)
	// halFloat := 1.0 / 2.0
	// fmt.Println("Half with float", halFloat)
	// squaring (raising to the power)
	// badThreeeSquared := 3 ^ 2
	// fmt.Println("badthreeeSquared", badThreeeSquared)
	// goodThreeSquared := math.Pow(3.0, 2.0)
	// fmt.Println("goodthreeSquared", goodThreeSquared)
	// modulus
	// reminder := 50 % 3
	// fmt.Println("Reminder", reminder)
	// unary operators
	// x := 3
	// optional to use value x in other variables
	// var y = x
	// y++
	// x++
	// fmt.Println("x is now", x)
	// x--
	// x--
	// fmt.Println("x is now", x)

	// presedence
	// multiplication and division
	// a := 12.0 * 3.0 / 4.0
	// b := (12.0 * 3.0) / 4.0
	// c := 12.0 * (3.0 / 4.0)
	// fmt.Println("a", a, "b", b, "c", c)
	// integer division
	// unclear := 12 * (3 / 4)
	// fmt.Println("Unclear", unclear)
	// parenthesis
	// f := 12.0 / 3.0 / 4.0
	// fmt.Println("f", f)
	// f = 12.0 / (3.0 / 4.0)
	// fmt.Println("f now", f)
	// addition/subtraction
	// fmt.Println()
	// x := 12 + 3 - 4
	// y := (12 + 3) - 4
	// z := 12 + (3 - 4)
	// fmt.Println("x", x)
	// fmt.Println("y", y)
	// fmt.Println("z", z)
	// multiply
	// fmt.Println()
	// x = 12 + 3*4
	// y = (12 + 3) * 4
	// z = 12 + (3 * 4)
	// fmt.Println("x", x, "y", y, "z", z)

	// modulus
	// odoes one number divide exactly into another?
	// x := 12
	// y := 5
	// if x%y == 0 {
	// 	fmt.Println(y, "Divides exactly into", x)
	// } else {
	// 	fmt.Println(y, "Does not divide exactly into", x)
	// }
	// modulus month
	// thisMonth := 12
	// fmt.Println("The month after", thisMonth, "is", thisMonth+1)
	// for m := 1; m <= 12; m++ {
	// 	fmt.Println("The month after", m, "is", m%12+1)
	// }

	// relational and conditional operators
	// second := 31
	// minute := 1
	// if (minute < 59) && ((second + 1) > 59) {
	// 	minute++
	// }
	// fmt.Println(minute)

	// short circuit evaluation
	// a := 12
	// b := 0
	// option 1 (not best practice)
	// if b != 0 {
	// 	c := divideTwoNumbers(a, b)

	// 	if c == 2 {
	// 		fmt.Println("We found a two")
	// 	}
	// }
	// option 2 (not best practice)
	// if b != 0 && divideTwoNumbers(a, b) == 2 {
	// 	fmt.Println("Found 2")
	// }
	// if divideTwoNumbers(a, b) == 2 && b != 0 {
	// 	fmt.Println("Found 2")
	// }
	// c, err := divideTwoNumbers(a, b)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	if c == 2 {
	// 		fmt.Println("we found 2")
	// 	}
	// }

	// Assignment operator
	// x := 12
	// x++
	// fmt.Println(x)
	// y := 10
	// y /= 2
	// fmt.Println(y)
}

// func divideTwoNumbers(x, y int) (int, error) {
// 	if y == 0 {
// 		return 0, errors.New("cannot divide by 0")
// 	}
// 	return x / y, nil
// }
