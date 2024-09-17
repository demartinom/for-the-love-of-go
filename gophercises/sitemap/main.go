package main

import (
	"flag"

	"github.com/demartinom/go-learning/gophercises/sitemap/sitemapper"
)

func main() {
	site := flag.String("url", "https://httpstat.us/", "URL to get sitemap from")
	flag.Parse()

	sitemapper.SiteMap(*site)
}
