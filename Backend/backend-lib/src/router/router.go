package router

import "net/http"

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{mux: http.NewServeMux()}
}
func (router *Router) Get(pattern string, handler http.HandlerFunc) {
	router.mux.HandleFunc("GET "+pattern, handler)
}
func (router *Router) Post(pattern string, handler http.HandlerFunc) {
	router.mux.HandleFunc("POST "+pattern, handler)
}
func (router *Router) Put(pattern string, handler http.HandlerFunc) {
	router.mux.HandleFunc("PUT "+pattern, handler)
}
func (router *Router) Delete(pattern string, handler http.HandlerFunc) {
	router.mux.HandleFunc("DELETE "+pattern, handler)
}
func (router *Router) ServeHTTP(responseWriter http.ResponseWriter, req *http.Request) {
	router.mux.ServeHTTP(responseWriter, req)
}
