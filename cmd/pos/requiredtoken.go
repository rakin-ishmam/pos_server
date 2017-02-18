package main

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/rakin-ishmam/pos_server/action/user"
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/response"
	"gopkg.in/mgo.v2"
)

func requiredToken(f http.HandlerFunc, session *mgo.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) == 0 {
			response.ServeErr(w, r, apperr.NewAuthentication("token check", "invalid"))
			return
		}
		action := user.NewTokenToUser(session, token)
		action.Do()
		user, role, err := action.Result()
		if err != nil {
			response.ServeErr(w, r, err)
			return
		}

		context.Set(r, "user", user)
		context.Set(r, "role", role)

		f(w, r)
	}
}
