package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestListBooksHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/books", nil)
	response := httptest.NewRecorder()
	booksHandler(response, request)
	expBooks, err := json.Marshal(books)
	if err != nil {
		t.Error("unexpected error", err)
	}
	b, err := io.ReadAll(response.Body)
	if string(expBooks) == string(b) {
		t.Errorf("unexpected response got %v, want %v", string(b), string(expBooks))
	}
}
func TestAddBookHandler(t *testing.T) {
	body := `{"name": "The Hitchhiker's Guide to the Galaxy","price": 1100,"currency": "$","author": "Douglas Adams"}`
	request := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(body))
	response := httptest.NewRecorder()
	booksHandler(response, request)
	var expBook book
	err := json.NewDecoder(response.Body).Decode(&expBook)
	if err != nil {
		t.Error("unexpected error", err)
	}
	if expBook.Name != "The Hitchhiker's Guide to the Galaxy" {
		t.Error("book is not in list")
	}
}
