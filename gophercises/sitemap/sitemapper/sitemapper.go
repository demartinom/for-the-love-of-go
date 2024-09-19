package sitemapper

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/demartinom/link/parser"
)

func SiteMap(site string) (*http.Response, error) {
	var localLinks []parser.Link

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
			localLinks = append(localLinks, v)
		}
	}
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	if err := enc.Encode(localLinks); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return nil, nil
}
