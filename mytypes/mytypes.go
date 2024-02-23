package mytypes

import "strings"

type MyInt int
type MyString string
type MyBuilder strings.Builder

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (string MyString) Len() int {
	return len(string)
}

func (builder MyBuilder) Hello() string {
	return "Hello, Gophers!"
}
