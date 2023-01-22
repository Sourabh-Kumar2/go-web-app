package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListBooksHandler(t *testing.T) {
	h := &ListBookHandler{}
	request := httptest.NewRequest(http.MethodGet, "/books", nil)
	response := httptest.NewRecorder()
	h.ServeHTTP(response, request)
	expBooks, err := json.Marshal(books)
	if err != nil {
		t.Error("unexpected error", err)
	}
	b, err := io.ReadAll(response.Body)
	if string(expBooks) == string(b) {
		t.Errorf("unexpected response got %v, want %v", string(b), string(expBooks))
	}
}
