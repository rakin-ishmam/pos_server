package main

import (
	"net/http"

	"github.com/rakin-ishmam/pos_server/action"
	"github.com/rakin-ishmam/pos_server/response"
	"gopkg.in/mgo.v2"
)

// JSONRunner run action JSONAction
func JSONRunner(f func(w http.ResponseWriter, r *http.Request, ses *mgo.Session) action.JSONAction, session *mgo.Session) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		action := f(w, r, session)
		action.Do()
		dt, err := action.Result()

		if err != nil {
			response.ServeErr(w, r, err)
			return
		}

		response.JSON(w, r, dt)
	}

}
