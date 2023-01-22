package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Method  string
	Path    string
	Handler http.Handler
}

type Router struct {
	mux *mux.Router
}

func NewRouter() *Router {
	return &Router{
		mux: mux.NewRouter(),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *Router) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc(pattern, handler)
}

func (r *Router) AddRoute(route Route) {
	r.mux.NewRoute().Methods(route.Method).Path(route.Path).Handler(route.Handler)
}
