package main

import "net/http"

// Route represents http pos api route
type Route struct {
	Metod   string
	Name    string
	Path    string
	Handler http.HandlerFunc
}
