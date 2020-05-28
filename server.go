package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {}

func newServer() *server{
	return &server{}
}

func (s server) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

func (s server) Run() {
	http.HandleFunc("/", s.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
