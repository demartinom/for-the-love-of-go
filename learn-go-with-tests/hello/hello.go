// Package for doing stuff
package main

import "fmt"

// Running Println in main rather than in a separate function
// func main() {
// 	fmt.Println("Hello, world")
// }

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "

// Says hello
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name // Say hello
}

// Main function
func main() {
	fmt.Println(Hello("world", "")) // A thing
}
