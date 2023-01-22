package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type AddBookHandler struct{}

func (a *AddBookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	handleErr(err)
	var b book
	handleErr(json.Unmarshal(body, &b))
	b.ID = int64(len(books) + 1)
	books = append(books, b)
	w.WriteHeader(http.StatusCreated)
	handleErr(json.NewEncoder(w).Encode(b))
}
