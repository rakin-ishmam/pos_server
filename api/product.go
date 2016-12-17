package api

import (
	"net/http"

	"github.com/rakin-ishmam/pos_server/action"
	mgo "gopkg.in/mgo.v2"
)

// CreateProduct returns action for create Product
func CreateProduct(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// UpdateProduct returns action for upddate Product
func UpdateProduct(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// DeleteProduct returns action for delete Product
func DeleteProduct(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// ListProduct returns action for list Product
func ListProduct(w http.ResponseWriter, h *http.Request, Session *mgo.Session) action.Action {

	return nil
}
