package api

import (
	"net/http"

	"github.com/rakin-ishmam/pos_server/action"
	mgo "gopkg.in/mgo.v2"
)

// CreateUser returns action for create user
func CreateUser(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// UpdateUser returns action for upddate user
func UpdateUser(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// DeleteUser returns action for delete user
func DeleteUser(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// ListUser returns action for list user
func ListUser(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// FetchUser returns action to get one user
func FetchUser(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {
	return nil
}
