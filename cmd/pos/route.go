package main

import "net/http"

// Route represents http pos api route
type Route struct {
	Method  string
	Name    string
	Path    string
	Handler http.HandlerFunc
}
