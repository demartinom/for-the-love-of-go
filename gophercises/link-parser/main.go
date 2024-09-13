package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	siteFlag := flag.String("site", "https://gophercises.com/", "What site you want to parse")
	flag.Parse()
	fmt.Println("reading from site ", *siteFlag)

	resp, err := http.Get(*siteFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

}
