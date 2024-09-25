package sitemapper

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/demartinom/link/parser"
)

func SiteMap(site string) {
	resp, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	base := baseUrl.String()

	parsedLinks := parser.Parse(resp.Body)

	var hrefs []string

	for _, l := range parsedLinks {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}
}
