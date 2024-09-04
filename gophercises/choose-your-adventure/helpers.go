package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func readJson(file string) []byte {
	jsonFile, err := os.Open("gopher.json")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	return byteValue
}
