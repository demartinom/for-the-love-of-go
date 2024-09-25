package sitemapper

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

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

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	base := baseUrl.String()

}
