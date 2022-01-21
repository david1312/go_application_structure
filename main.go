package main

import (
	"fmt"
)

// package scoped variables and constants
var x int = 100

const y = 0

func main() {
	// Local Scoped Variables and Constants Declarations, calling functions etc
	var a int = 7
	var b float64 = 3.5
	const c int = 10

	// Println() function prints out a line to stdout
	// It belongs to package fmt
	fmt.Println("Hello Go world!") // => Hello Go world!
	fmt.Println(a, b, c)           // => 7 3.5 10
}
