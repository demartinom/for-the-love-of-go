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
var story = readJson("gopher.json")
var err = json.Unmarshal(story, &storyText)

func home(w http.ResponseWriter, r *http.Request) {

	if err != nil {
		fmt.Println("error", err)
		return
	}
	files := []string{"./web/html/base.tmpl.html", "./web/html/home.tmpl.html"}
	handleTemplates(w, files, storyText.Intro)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Starting server at port %s\n", "8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
