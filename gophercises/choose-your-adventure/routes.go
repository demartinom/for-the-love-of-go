package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {

	files := []string{"./web/html/base.tmpl.html", "./web/html/home.tmpl.html"}
	handleTemplates(w, files, storyText.Intro)
}

func newYork(w http.ResponseWriter, r *http.Request) {
	files := []string{"./web/html/base.tmpl.html", "./web/html/newyork.tmpl.html"}
	handleTemplates(w, files, storyText.NewYork)
}

func denver(w http.ResponseWriter, r *http.Request) {
	files := []string{"./web/html/base.tmpl.html", "./web/html/denver.tmpl.html"}
	handleTemplates(w, files, storyText.Denver)
}
