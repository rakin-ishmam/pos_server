package main

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/response"
)

func panicRecover(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer catchPanic(w, r)
		f(w, r)
	}
}

func catchPanic(w http.ResponseWriter, r *http.Request) {
	err := recover()
	if err != nil {
		log.Println("error: catchErr: ", err)
		response.ServeErr(w, r, apperr.Internal{Where: "app panic"})
	}

	context.Clear(r)
}
