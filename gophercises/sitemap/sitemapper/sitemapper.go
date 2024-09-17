package sitemapper

import (
	"log"
	"net/http"

	"github.com/demartinom/link/parser"
)

func SiteMap(site string) (*http.Response, error) {
	resp, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	parser.Parse(resp.Body)
	return nil, nil
}
