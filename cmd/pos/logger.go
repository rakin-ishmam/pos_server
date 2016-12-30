package main

import (
	"log"
	"net/http"
)

func logRoute(r Route) http.HandlerFunc {
	return func(w http.ResponseWriter, h *http.Request) {
		log.Printf("start: Name=%s, Method=%s, Path=%s", r.Name, r.Metod, r.Path)
		r.Handler(w, h)
		log.Printf("end: Name=%s, Method=%s, Path=%s", r.Name, r.Metod, r.Path)
	}
}
