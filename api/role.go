package api

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/rakin-ishmam/pos_server/action"
	"github.com/rakin-ishmam/pos_server/action/role"
	"github.com/rakin-ishmam/pos_server/data"
	mgo "gopkg.in/mgo.v2"
)

// CreateRole returns action for create Role
func CreateRole(w http.ResponseWriter, r *http.Request, session *mgo.Session) action.JSONAction {
	createDt := role.CreatePayload{}
	if errAc := jsonDecode(r, &createDt, "CreateRole", "create payload"); errAc != nil {
		return errAc
	}

	usr := context.Get(r, "user").(data.User)
	rl := context.Get(r, "role").(data.Role)

	return role.NewCreate(createDt, usr, rl, session)
}

// UpdateRole returns action for upddate Role
func UpdateRole(w http.ResponseWriter, r *http.Request, session *mgo.Session) action.JSONAction {
	id, errAc := idFetch(r, "UpdateRole")
	if errAc != nil {
		return errAc
	}

	updateDt := role.UpdatePayload{}

	if errAc = jsonDecode(r, &updateDt, "UpdateRole", "update payload"); errAc != nil {
		return errAc
	}

	updateDt.ID.ID = id

	usr := context.Get(r, "user").(data.User)
	rl := context.Get(r, "role").(data.Role)

	return role.NewUpdate(updateDt, usr, rl, session)
}

// DeleteRole returns action for delete Role
func DeleteRole(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// ListRole returns action for list Role
func ListRole(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// FetchRole returns action to get one role
func FetchRole(w http.ResponseWriter, r *http.Request, session *mgo.Session) action.JSONAction {
	id, errAc := idFetch(r, "FetchRole")
	if errAc != nil {
		return errAc
	}

	var usr data.User
	var rl data.Role
	usr = context.Get(r, "user").(data.User)
	rl = context.Get(r, "role").(data.Role)

	return role.NewFetch(id, usr, rl, session)
}
