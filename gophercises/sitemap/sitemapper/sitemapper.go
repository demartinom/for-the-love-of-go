package sitemapper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/demartinom/link/parser"
)

func SiteMap(site string) {
	resp, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	parsedLinks := parser.Parse(resp.Body)

	for _, l := range parsedLinks {
		fmt.Printf("%+v\n", l)
	}
}
