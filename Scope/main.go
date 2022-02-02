package main

import (
	"scopeapp/packageone"
)

var myVar = "This is a package level variable"

func main() {
	// variables
	// declare a package level variable for the main
	// package name myVar

	// declare a block variable for the main function
	// called blockVar

	// declare a package level variable in the packageone
	// package name PackageVar

	// create an exported function in packageone called PrintMe

	// in the main function, prin out the values of myVar,
	// blockvar, and PackageVar on one line using the PrintMe
	// function in packageone

	var blockVar = "This is the block level variable"

	packageone.PrintMe(myVar, blockVar)

}
