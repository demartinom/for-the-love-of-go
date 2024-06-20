package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	//CLI flag for setting the port
	addr := flag.String("addr", ":4000", "HTTP network address")
	//Parse the flag
	flag.Parse()

	mux := http.NewServeMux()

	//Create file server
	//Relative to project directory
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting on server %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
