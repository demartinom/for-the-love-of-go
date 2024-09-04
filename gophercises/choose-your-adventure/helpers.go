package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handleTemplates(w http.ResponseWriter, templates []string) {
	ts, err := template.ParseFiles(templates...)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Println(err)
	}
}
