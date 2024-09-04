package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Arc struct {
	Intro     Title `json:"intro"`
	NewYork   Title `json:"new-york"`
	Debate    Title `json:"debate"`
	SeanKelly Title `json:"sean-kelly"`
	MarkBates Title `json:"mark-bates"`
	Denver    Title `json:"denver"`
	Home      Title `json:"home"`
}

type Title struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{"./web/html/base.tmpl.html"}

	var story Story
	err := json.Unmarshal([]byte(),%story)

	handleTemplates(w, files)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Starting server at port %s\n", "8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
