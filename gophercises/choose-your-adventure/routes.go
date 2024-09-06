package main

import "net/http"

var files = []string{"./web/html/base.tmpl.html", "./web/html/page.tmpl.html"}

func intro(w http.ResponseWriter, r *http.Request) {
	handleTemplates(w, files, storyText.Intro)
}

func newYork(w http.ResponseWriter, r *http.Request) {
	handleTemplates(w, files, storyText.NewYork)
}

func debate(w http.ResponseWriter, r *http.Request) {
	handleTemplates(w, files, storyText.Debate)
}

func seanKelly(w http.ResponseWriter, r *http.Request) {
	handleTemplates(w, files, storyText.SeanKelly)
}

func markBates(w http.ResponseWriter, r *http.Request) {
	handleTemplates(w, files, storyText.MarkBates)
}

func denver(w http.ResponseWriter, r *http.Request) {
	handleTemplates(w, files, storyText.Denver)
}
func home(w http.ResponseWriter, r *http.Request) {
	handleTemplates(w, files, storyText.Home)
}
