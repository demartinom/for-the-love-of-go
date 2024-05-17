// Package for doing stuff
package main

import "fmt"

// Running Println in main rather than in a separate function
// func main() {
// 	fmt.Println("Hello, world")
// }

const englishHelloPrefix = "Hello, "

// Says hello
func Hello(name string) string {
	return englishHelloPrefix + name // Say hello
}

// Main function
func main() {
	fmt.Println(Hello("world")) // A thing
}
