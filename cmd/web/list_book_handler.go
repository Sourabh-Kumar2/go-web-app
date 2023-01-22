package main

import (
	"encoding/json"
	"net/http"
)

type ListBookHandler struct{}

func (ListBookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	handleErr(json.NewEncoder(w).Encode(books))
}
