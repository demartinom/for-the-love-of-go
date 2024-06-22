package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	//CLI flag for setting the port
	addr := flag.String("addr", ":4000", "HTTP network address")
	//Parse the flag
	flag.Parse()

	//Create loggers for writing messages
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	//Create file server
	//Relative to project directory
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Staring at server %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
