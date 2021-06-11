package myapp

import "net/http"

// return NewHandler
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	return mux
}
