package api

import (
	"net/http"

	"github.com/rakin-ishmam/pos_server/action"
	mgo "gopkg.in/mgo.v2"
)

// CreateRole returns action for create Role
func CreateRole(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// UpdateRole returns action for upddate Role
func UpdateRole(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// DeleteRole returns action for delete Role
func DeleteRole(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// ListRole returns action for list Role
func ListRole(w http.ResponseWriter, h *http.Request, Session *mgo.Session) action.Action {

	return nil
}