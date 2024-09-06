package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Story struct {
	Intro     StoryArc `json:"intro"`
	NewYork   StoryArc `json:"new-york"`
	Debate    StoryArc `json:"debate"`
	SeanKelly StoryArc `json:"sean-kelly"`
	MarkBates StoryArc `json:"mark-bates"`
	Denver    StoryArc `json:"denver"`
	Home      StoryArc `json:"home"`
}

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

var storyText Story

func unmarshall() Story {
	story := readJson("gopher.json")
	err := json.Unmarshal(story, &storyText)
	if err != nil {
		fmt.Println("error", err)

	}
	return storyText
}

func home(w http.ResponseWriter, r *http.Request) {

	files := []string{"./web/html/base.tmpl.html", "./web/html/home.tmpl.html"}
	handleTemplates(w, files, storyText.Intro)
}

func newYork(w http.ResponseWriter, r *http.Request) {
	files := []string{"./web/html/base.tmpl.html", "./web/html/newyork.tmpl.html"}
	handleTemplates(w, files, storyText.NewYork)
}

func main() {
	unmarshall()
	http.HandleFunc("/", home)
	http.HandleFunc("/new-york", newYork)
	fmt.Printf("Starting server at port %s\n", "8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
