// Package main that will launch into the command line
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	port int

	h1 = func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}

	h2 = func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}
)

func init() {
	log.Println(" initializing the application FN:init ")
	flag.IntVar(&port, "port", 8181, "port to run the application on")
	flag.Parse()
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	http.HandleFunc("/", h1)
	http.HandleFunc("/test", h2)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
