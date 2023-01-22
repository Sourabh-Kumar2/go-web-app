package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddBookHandler(t *testing.T) {
	h := &AddBookHandler{}
	body := `{"name": "The Hitchhiker's Guide to the Galaxy","price": 1100,"currency": "$","author": "Douglas Adams"}`
	request := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(body))
	response := httptest.NewRecorder()
	h.ServeHTTP(response, request)
	var expBook book
	err := json.NewDecoder(response.Body).Decode(&expBook)
	if err != nil {
		t.Error("unexpected error", err)
	}
	if expBook.Name != "The Hitchhiker's Guide to the Galaxy" {
		t.Error("book is not in list")
	}
}
