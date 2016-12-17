package api

import (
	"net/http"

	"github.com/rakin-ishmam/pos_server/action"
	mgo "gopkg.in/mgo.v2"
)

// CreateOrder returns action for create Order
func CreateOrder(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// UpdateOrder returns action for upddate Order
func UpdateOrder(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// DeleteOrder returns action for delete Order
func DeleteOrder(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// ListOrder returns action for list Order
func ListOrder(w http.ResponseWriter, h *http.Request, Session *mgo.Session) action.Action {

	return nil
}
