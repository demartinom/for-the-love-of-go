package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
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
