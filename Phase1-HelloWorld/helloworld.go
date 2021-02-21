/*
 * Author: Jacky Lin
 * Date: 2021-02-17
 * Programming Language Name: Go 
 * Compile Method: 
 * 		1. Directly run the program: go run <filename>
 * 			e.g. go run helloworld.go
 * 		2. Compile to executable and run:
 * 			2.1 go build <filename>
 * 			2.2 ./<filename>
 * 			e.g. go build helloworld.go && ./helloworld
 * Citation: Go essential training Online Class: Linkedin learning, formerly Lynda.com. (n.d.). Retrieved February 17, 2021, from https://www.linkedin.com/learning/go-essential-training/
 */

// This is the comment in go language
// Go is organized in packages
package main

// This is the import in go language
// fmt package contains print method 
//that we would use to print "hello world"
import "fmt"

// In go we use func to define a function
// Function main will execute automatically in run time
func main () {
	// We use the format package.function to use a function 
	// from imported package
	// fmt.Println do the print in go
	fmt.Println("Hello, World.")
}
