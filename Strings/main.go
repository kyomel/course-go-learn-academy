package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println()
	// name := "Hello Francesca"
	// fmt.Println("String:", name)
	// fmt.Println()

	// fmt.Println("Bytes")
	// for i := 0; i < len(name); i++ {
	// 	fmt.Printf("%x ", name[i])
	// }
	// fmt.Println()
	// fmt.Println()
	// fmt.Println("Index\tRune\tString")
	// for x, y := range name {
	// 	fmt.Println(x, "\t", y, "\t", string(y))
	// }
	// fmt.Println()
	// fmt.Println("Three ways to concatenate strings")
	// h := "Hello, "
	// w := "world."
	// using logic + not very efficient
	// myString := h + w
	// fmt.Println(myString)
	// fmt.Println()
	// using fmt more efficient
	// myString = fmt.Sprintf("%s%s", h, w)
	// fmt.Println(myString)
	// fmt.Println()
	// using stringbuilder very efficient
	// var sb strings.Builder
	// sb.WriteString(h)
	// sb.WriteString(w)
	// fmt.Println(sb.String())
	// fmt.Println()
	// Get string in array
	// fmt.Println()
	// name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// fmt.Println("Getting a substring")
	// fmt.Println(name[10:13])

	// indexing
	// courseName := "Learn Go for Beginners Crash Course"
	// fmt.Println(string(courseName[0]))
	// fmt.Println(string(courseName[6]))
	// for i := 0; i <= 21; i++ {
	// 	fmt.Println(string(courseName[i]))
	// }
	// fmt.Println()
	// for i := 13; i <= 21; i++ {
	// 	fmt.Println(string(courseName[i]))
	// }
	// fmt.Println()

	// string length
	// fmt.Println("Length of courseName is", len(courseName))
	// var mySlice []string
	// mySlice = append(mySlice, "one")
	// mySlice = append(mySlice, "two")
	// mySlice = append(mySlice, "three")
	// fmt.Println("mySlice has", len(mySlice), "elements")
	// fmt.Println("the last element is mySlice is", mySlice[len(mySlice)-1])
	// courses := []string{
	// 	"Learn Go for Beginners Crash Course",
	// 	"Learn Java for Beginners Crash Course",
	// 	"Learn Python for Beginners Crash Course",
	// 	"Learn C for Beginners Crash Course",
	// }
	// for _, x := range courses {
	// 	if strings.Contains(x, "Go") {
	// 		fmt.Println("Go is found in", x, "and index is", strings.Index(x, "Go"))
	// 	}
	// }
	// newString := "Go is a great programming language. Go for it!"
	// fmt.Println(strings.HasPrefix(newString, "Go"))
	// fmt.Println(strings.HasPrefix(newString, "Python"))
	// fmt.Println(strings.HasSuffix(newString, "!"))
	// fmt.Println(strings.Count(newString, "Go"))
	// fmt.Println(strings.Count(newString, "Fish"))
	// fmt.Println(strings.Index(newString, "Python"))
	// fmt.Println(strings.LastIndex(newString, "Go"))

	// string manipulation
	// if strings.Contains(newString, "Go") {
	// 	newString = strings.ReplaceAll(newString, "Go", "Golang")
	// newString = strings.Replace(newString, "Go", "Golang", 1)
	// }
	// fmt.Println(newString)
	// string comparison
	// a := "A"
	// if a == "A" {
	// 	fmt.Println("a is equal to A")
	// }
	// if "A" > "B" {
	// 	fmt.Println("A is greater than B")
	// } else {
	// 	fmt.Println("A is not greater than B")
	// }
	// if "Alpha" > "Absolute" {
	// 	fmt.Println("A is greater than B")
	// } else {
	// 	fmt.Println("A is not greater than B")
	// }
	// badEmail := " me@here.com "
	// badEmail = strings.TrimSpace(badEmail)
	// fmt.Printf("=%s=", badEmail)
	// fmt.Println()
	// string manipulation part 2
	// str := "alpha alpha alpha alpha alpha"
	// str = replaceNth(str, "alpha", "beta", 3)
	// fmt.Println(str)

	// case
	// newString := "the quick brown FOX jumped over the lazy red DOG"
	// fmt.Println(strings.ToLower(newString))
	// fmt.Println(strings.ToUpper(newString))
	// fmt.Println(strings.Title(newString))
	// fmt.Println(strings.Title(strings.ToLower(newString)))

	// case 2
	myString := "This is a clear EXAMPLE of why we search in one case only."
	searchString := strings.ToLower(myString)
	// option 1
	// if strings.Contains(searchString, "this") {
	// 	fmt.Println("Found it!")
	// } else {
	// 	fmt.Println("Did not find it!")
	// }
	// option 2
	if strings.Contains(strings.ToLower(searchString), "this") {
		fmt.Println("Found it!")
	} else {
		fmt.Println("Did not find it!")
	}

	// other case functions
	fmt.Println(strings.ToLower(myString))
	fmt.Println(strings.ToUpper(myString))
	fmt.Println(strings.Title(myString))
	fmt.Println(strings.Title(strings.ToLower(myString)))
}

// func replaceNth(s, old, new string, n int) string {
// 	// index
// 	i := 0
// 	for j := 0; j <= n; j++ {
// 		x := strings.Index(s[i:], old)
// 		if x < 0 {
// 			// we did not find it
// 			break
// 		}
// 		// have found it
// 		i = i + x
// 		if j == n {
// 			return s[:i] + new + s[i+len(old):]
// 		}
// 		i += len(old)
// 	}
// 	return s
// }
