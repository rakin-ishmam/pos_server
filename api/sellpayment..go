package api

import (
	"net/http"

	"github.com/rakin-ishmam/pos_server/action"
	mgo "gopkg.in/mgo.v2"
)

// CreateSellPayment returns action for create SellPayment
func CreateSellPayment(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// UpdateSellPayment returns action for upddate SellPayment
func UpdateSellPayment(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// DeleteSellPayment returns action for delete SellPayment
func DeleteSellPayment(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// ListSellPayment returns action for list SellPayment
func ListSellPayment(w http.ResponseWriter, h *http.Request, Session *mgo.Session) action.Action {

	return nil
}
