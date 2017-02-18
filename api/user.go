package api

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/rakin-ishmam/pos_server/action"
	"github.com/rakin-ishmam/pos_server/action/user"
	"github.com/rakin-ishmam/pos_server/data"
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
func FetchUser(w http.ResponseWriter, r *http.Request, session *mgo.Session) action.JSONAction {
	id, errAc := idFetch(r, "fetch user")
	if errAc != nil {
		return errAc
	}

	var usr data.User
	var role data.Role
	usr = context.Get(r, "user").(data.User)
	role = context.Get(r, "role").(data.Role)

	return user.NewFetch(session, id, usr, role)
}

// Login returns action to get login
func Login(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	loginDt := user.LoginPayload{}

	if errAc := jsonDecode(r, &loginDt, "login", "username and password"); errAc != nil {
		return errAc
	}

	return &user.Login{
		Session:    Session,
		ReqPayload: loginDt,
	}
}
