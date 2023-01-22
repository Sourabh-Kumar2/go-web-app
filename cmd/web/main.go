package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	router.AddRoute(Route{Method: http.MethodGet, Path: "/books", Handler: &ListBookHandler{}})
	router.AddRoute(Route{Method: http.MethodPost, Path: "/books", Handler: &AddBookHandler{}})
	log.Fatal(http.ListenAndServe(":5111", router))
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
