package api

import (
	"net/http"

	"github.com/rakin-ishmam/pos_server/action"
	mgo "gopkg.in/mgo.v2"
)

// CreateCustomer returns action for create Customer
func CreateCustomer(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// UpdateCustomer returns action for upddate Customer
func UpdateCustomer(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// DeleteCustomer returns action for delete Customer
func DeleteCustomer(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// ListCustomer returns action for list Customer
func ListCustomer(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}
