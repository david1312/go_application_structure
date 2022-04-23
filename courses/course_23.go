package courses

import "fmt"

//** COMMENTS **//

// this is a single line comment

/*
 This is a block comment.
 a := 10
 fmt.Println(a)
*/

var name = "John Wick" // inline comment

//** NAMING CONVENTIONS IN GO **//

// Naming Conventions are important for code readability and maintainability.

// use short, concise names especially in shorter scopes
// common names for common types:
var s string   //string
var i int      //index
var num int    //number
var msg string //message
var v string   //value
var err error  //error value
var done bool  //bool, has been done?

// use mixedCase a.k.a camelCase instead of snake_case (variables and  functions)
var maxValue = 100  // recommended (camelCase)
var max_value = 100 // not recommended (snake_case)

// recommended
func writeToFile() {
	fmt.Println("recommended")
}

// not recommended
func write_to_file() {
	fmt.Println("not recommended")
}

// write acronyms in all caps
var writeToDB = true // recommended
var writeToDb = true // not recommended

func NamingConvention() {
	fmt.Println(name, s, i, num, msg, v, err, done, maxValue, max_value, writeToDB, writeToDb)
	writeToFile()
	write_to_file()
}
