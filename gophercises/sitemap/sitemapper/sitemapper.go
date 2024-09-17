package sitemapper

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func SiteMap(site string) (*http.Response, error) {
	resp, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err: %s ", err)
		return nil, err
	}
	fmt.Println(string(body))
	return resp, nil
}
