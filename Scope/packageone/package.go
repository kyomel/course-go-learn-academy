package packageone

import "fmt"

var PackageVar = "This is a package level variable in packageone"

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, PackageVar)
}
