package api

import (
	"net/http"

	"github.com/rakin-ishmam/pos_server/action"
	mgo "gopkg.in/mgo.v2"
)

// CreateUser returns action for create user
func CreateUser(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// UpdateUser returns action for upddate user
func UpdateUser(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// DeleteUser returns action for delete user
func DeleteUser(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// ListUser returns action for list user
func ListUser(w http.ResponseWriter, h *http.Request, Session *mgo.Session) action.Action {

	return nil
}
