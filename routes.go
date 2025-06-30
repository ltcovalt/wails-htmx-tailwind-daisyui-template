package main

import (
	"net/http"
)

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hx/header", Hx.Header)
	mux.HandleFunc("GET /hx/theme-controller", Hx.ThemeController)
	return mux
}
