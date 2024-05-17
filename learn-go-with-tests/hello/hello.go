// Package for doing stuff
package main

import "fmt"

// Running Println in main rather than in a separate function
// func main() {
// 	fmt.Println("Hello, world")
// }

// Says hello
func Hello() string {
	return "Hello, world" // Say hello
}

// Main function
func main() {
	fmt.Println(Hello()) // A thing
}
