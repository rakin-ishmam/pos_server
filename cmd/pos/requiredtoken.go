package main

import (
	"fmt"
	"net/http"
)

func requiredToken(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("token", r.Header.Get("token"))
		f(w, r)
	}
}
