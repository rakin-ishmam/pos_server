package api

import (
	"net/http"

	"github.com/rakin-ishmam/pos_server/action"
	mgo "gopkg.in/mgo.v2"
)

// CreateCategory returns action for create Category
func CreateCategory(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// UpdateCategory returns action for upddate Category
func UpdateCategory(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// DeleteCategory returns action for delete Category
func DeleteCategory(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// ListCategory returns action for list Category
func ListCategory(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}
