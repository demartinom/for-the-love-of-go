package sitemapper

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/demartinom/link/parser"
)

func SiteMap(site string) (*http.Response, error) {
	resp, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	parsedLinks := parser.Parse(resp.Body)
	for _, v := range parsedLinks {
		if !strings.HasPrefix(v.Href, "/") {
			continue
		} else {
			fmt.Printf("%+v\n", v)
		}
	}
	return nil, nil
}
