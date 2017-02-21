package api

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/rakin-ishmam/pos_server/action"
	"github.com/rakin-ishmam/pos_server/action/user"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db/query"
	mgo "gopkg.in/mgo.v2"
	"github.com/rakin-ishmam/pos_server/action/empty"
)

// CreateUser returns action for create user
func CreateUser(w http.ResponseWriter, r *http.Request, session *mgo.Session) action.JSONAction {
	usrDt := user.CreatePayload{}
	if errAc := jsonDecode(r, &usrDt, "CreateUser", "create payload"); errAc != nil {
		return errAc
	}

	usr := context.Get(r, "user").(data.User)
	role := context.Get(r, "role").(data.Role)

	return user.NewCreate(usrDt, usr, role, session)
}

// UpdateUser returns action for upddate user
func UpdateUser(w http.ResponseWriter, r *http.Request, session *mgo.Session) action.JSONAction {
	id, errAc := idFetch(r, "UpdateUser")
	if errAc != nil {
		return errAc
	}

	updateDt := user.UpdatePayload{}

	if errAc = jsonDecode(r, &updateDt, "UpdateUser", "update payload"); errAc != nil {
		return errAc
	}

	updateDt.ID.ID = id

	usr := context.Get(r, "user").(data.User)
	role := context.Get(r, "role").(data.Role)

	return user.NewUpdateAction(updateDt, usr, role, session)
}

// DeleteUser returns action for delete user
func DeleteUser(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// ListUser returns action for list user
func ListUser(w http.ResponseWriter, r *http.Request, session *mgo.Session) action.JSONAction {
	uq := query.User{}

	err := urlquery(&uq.GenInfo, r, urlQSkip("ListUser"), urlQLimit("ListUser"))
	if err != nil {
		return empty.NewJSON(err)
	}
	
	usr := context.Get(r, "user").(data.User)
	role := context.Get(r, "role").(data.Role)

	return user.NewList(session, uq, usr, role)
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
