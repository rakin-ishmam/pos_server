package api

import (
	"encoding/json"
	"net/http"

	"github.com/rakin-ishmam/pos_server/action"
	"github.com/rakin-ishmam/pos_server/action/empty"
	"github.com/rakin-ishmam/pos_server/action/user"
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

// Login returns action to get login
func Login(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	loginDt := user.LoginPayload{}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&loginDt); err != nil {
		acc := empty.JSON{Err: err}
		return &acc
	}

	return &user.Login{
		Session:    Session,
		ReqPayload: loginDt,
	}
}
