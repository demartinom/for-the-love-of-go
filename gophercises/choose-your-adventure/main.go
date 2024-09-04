package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	var arc Arc
	storyArc := readJson("gopher.json")
	err := json.Unmarshal(storyArc, &arc)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(arc.Intro)
	files := []string{"./web/html/base.tmpl.html"}
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
