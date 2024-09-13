package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
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

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Function taken from html docs
	// Allows function to be referenced within itself
	var f func(*html.Node)

	f = func(n *html.Node) {
		// Checks if node is an anchor tag
		if n.Type == html.ElementNode && n.Data == "a" {
			// If it is an anchor tag, it looks over its attributes
			for _, a := range n.Attr {
				// If an href attribute is found, it prints the value of the href
				if a.Key == "href" {
					fmt.Println(a.Val)
				}

			}
		}
		//Allows func to visit all children and siblings of current node
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
