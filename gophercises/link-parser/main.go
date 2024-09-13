package main

import (
	"flag"
	"fmt"
)

func main() {
	siteFlag := flag.String("site", "https://gophercises.com/", "What site you want to parse")
	flag.Parse()
	fmt.Println("reading from site ", *siteFlag)
}
