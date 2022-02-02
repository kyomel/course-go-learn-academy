package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type
// type Animal interface {
// 	Says() string
// 	HowManyLegs() int
// }

// // Dog is the type for dogs
// type Dog struct {
// 	Name         string
// 	Sound        string
// 	NumberOfLegs int
// }

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

// func (d *Dog) Says() string {
// 	return d.Sound
// }

// func (d *Dog) HowManyLegs() int {
// 	return d.NumberOfLegs
// }

// // Cat is the type for cats
// type Cat struct {
// 	Name         string
// 	Sound        string
// 	NumberOfLegs int
// 	HasTail      bool
// }

// func (c *Cat) Says() string {
// 	return c.Sound
// }

// func (c *Cat) HowManyLegs() int {
// 	return c.NumberOfLegs
// }

// type Animal struct {
// 	Name         string
// 	Sound        string
// 	NumberOfLegs int
// }

// // Animal function
// func (a *Animal) Says() {
// 	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
// 	fmt.Println()
// }

// func (a *Animal) HowManyLegs() {
// 	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
// 	fmt.Println()
// }

func main() {
	// Pointer Section
	// var myInt int
	// myInt = 10
	// fmt.Println(myInt)
	// x := 10
	// myFirstPointer := &x
	// fmt.Println("x is", x)
	// fmt.Println("myFirstPointer is", myFirstPointer)
	// *myFirstPointer = 15
	// fmt.Println("x is now", x)
	// changeValueOfPointer(&x)
	// fmt.Println("After function call, x is now", x)

	// Slices Section
	// var animals []string
	// animals = append(animals, "dog")
	// animals = append(animals, "fish")
	// animals = append(animals, "cat")
	// animals = append(animals, "horse")

	// fmt.Println(animals)

	// for i, x := range animals {
	// 	fmt.Println(i, x)
	// }

	// fmt.Println("Element 0 is", animals[0])

	// fmt.Println("First two elements are", animals[0:2])

	// fmt.Println("The slice is", len(animals), "elements long")

	// fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	// sort.Strings(animals)

	// fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))
	// fmt.Println(animals)

	// animals = DeleteFromSlice(animals, 1)
	// fmt.Println(animals)

	// maps
	// intMap := make(map[string]int)
	// intMap["one"] = 1
	// intMap["two"] = 2
	// intMap["three"] = 3
	// intMap["four"] = 4
	// intMap["five"] = 5

	// for key, value := range intMap {
	// 	fmt.Println(key, value)
	// }

	// // delete(intMap, "four")

	// el, ok := intMap["four"]
	// if ok {
	// 	fmt.Println(el, "is in map")
	// } else {
	// 	fmt.Println(el, "is in not map")
	// }

	// intMap["two"] = 4

	// functions
	// z := addTwoNumber(2, 4)
	// fmt.Println(z)
	// myTotal := sumMany(2, 3, 4, 5, 88, 7, -5)
	// fmt.Println(myTotal)
	// var dog Animal
	// dog.Name = "dog"
	// dog.Sound = "woof"
	// dog.Says()

	// cat := Animal{
	// 	Name:         "cat",
	// 	Sound:        "meow",
	// 	NumberOfLegs: 4,
	// }
	// cat.Says()
	// cat.HowManyLegs()

	// channel
	// go doSomething("Hello, world")
	// fmt.Println("This is another message")
	// for {
	// 	// do nothing
	// }
	// keyPressChan = make(chan rune)
	// go listenForKeyPress()
	// fmt.Println("Press any key, or q to quit")
	// _ = keyboard.Open()
	// defer func() {
	// 	keyboard.Close()
	// }()
	// for {
	// 	char, _, _ := keyboard.GetSingleKey()
	// 	if char == 'q' || char == 'Q' {
	// 		break
	// 	}
	// 	keyPressChan <- char
	// }

	// interfaces
	// dog := Dog{
	// 	Name:         "dog",
	// 	Sound:        "woof",
	// 	NumberOfLegs: 4,
	// }
	// Riddle(&dog)
	// var cat Cat
	// cat.Name = "cat"
	// cat.NumberOfLegs = 4
	// cat.Sound = "meow"
	// cat.HasTail = true
	// Riddle(&cat)

	// Expressions
	// age := 10
	// name := "Jack"
	// rightHanded := true
	// fmt.Printf("%s is %d years old. Right handed: %t\n", name, age, rightHanded)
	// ageInTenYears := age + 10
	// fmt.Printf("In ten years, %s will be %d years old\n", name, ageInTenYears)
	// isATeenager := age >= 13
	// fmt.Println(name, "is a teenager:", isATeenager)

	// Booleans
	// apples := 18
	// oranges := 9
	// // boolean expressions
	// fmt.Println(apples == oranges)
	// fmt.Println(apples != oranges)
	// // > < >= <=
	// fmt.Printf("%d > %d: %t\n", apples, oranges, apples > oranges)
	// fmt.Printf("%d > %d: %t\n", apples, oranges, apples < oranges)
	// fmt.Printf("%d > %d: %t\n", apples, oranges, apples >= oranges)
	// fmt.Printf("%d > %d: %t\n", apples, oranges, apples <= oranges)

	// Compound Booleans
	jack := Employee{
		Name:     "Jack Smith",
		Age:      27,
		Salary:   40000,
		FullTime: false,
	}
	jill := Employee{
		Name:     "Jill Jones",
		Age:      33,
		Salary:   60000,
		FullTime: true,
	}
	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)
	for _, x := range employees {
		if x.Age > 30 {
			fmt.Println(x.Name, "is 30 or older")
		} else {
			fmt.Println(x.Name, "is under 30")
		}

		if x.Age > 30 && x.Salary > 50000 {
			fmt.Println(x.Name, "makes more than 50000 and is over 30")
		} else {
			fmt.Println("Either", x.Name, "makes less than 50000 and is over 30")
		}

		if x.Age > 30 || x.Salary > 50000 {
			fmt.Println(x.Name, "makes more than 50000 or is over 30")
		} else {
			fmt.Println(x.Name, "makes less than 50000 and is over 30")
		}

		if (x.Age > 30 || x.Salary < 50000) && x.FullTime {
			fmt.Println(x.Name, "matches our unclear criteria")
		}
	}
}

// func changeValueOfPointer(num *int) {
// 	*num = 25
// }

// DeleteFromSlice
// func DeleteFromSlice(a []string, i int) []string {
// 	a[i] = a[len(a)-1]
// 	a[len(a)-1] = ""
// 	a = a[:len(a)-1]
// 	return a
// }

// func addTwoNumber(x, y int) int {
// 	return x + y
// }

// func sumMany(nums ...int) int {
// 	total := 0
// 	for _, x := range nums {
// 		total = total + x
// 	}
// 	return total
// }

// func doSomething(s string) {
// 	until := 0
// 	for {
// 		fmt.Println("s is", s)
// 		until = until + 1
// 		if until >= 5 {
// 			break
// 		}
// 	}
// }

// func listenForKeyPress() {
// 	for {
// 		key := <-keyPressChan
// 		fmt.Println("You pressed", string(key))
// 	}
// }

// to answer about Riddle Dog
// func Riddle(a Animal) {
// 	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it?`, a.Says(), a.HowManyLegs())
// 	fmt.Println(riddle)
// }
