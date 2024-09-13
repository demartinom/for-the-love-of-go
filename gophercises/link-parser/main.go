package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	//Test HTML
	exHtml := `<html>
	<body>
	<h1>Hello!</h1>
	<a href="/other-page">A link to another page</a>
	</body>
	</html>`
	// Function taken from html docs

	doc, err := html.Parse(strings.NewReader(exHtml))
	if err != nil {
		log.Fatal(err)
	}

	var reading func(*html.Node)
	reading = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
				}
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					fmt.Println(c.Data)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			reading(c)
		}

	}
	reading(doc)
}
