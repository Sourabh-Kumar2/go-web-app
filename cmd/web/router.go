package main

import "net/http"

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *Router) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc(pattern, handler)
}
