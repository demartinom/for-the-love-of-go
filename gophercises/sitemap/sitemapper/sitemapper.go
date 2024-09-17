package sitemapper

import (
	"log"
	"net/http"
)

func SiteMap(site string) *http.Response {
	resp, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return resp
}
