package api

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/rakin-ishmam/pos_server/action"
)

// CreateInventory returns action for create Inventory
func CreateInventory(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// UpdateInventory returns action for upddate Inventory
func UpdateInventory(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// DeleteInventory returns action for delete Inventory
func DeleteInventory(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// ListInventory returns action for list Inventory
func ListInventory(w http.ResponseWriter, h *http.Request, Session *mgo.Session) action.Action {

	return nil
}
