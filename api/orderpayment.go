package api

import (
	"net/http"

	"github.com/rakin-ishmam/pos_server/action"
	mgo "gopkg.in/mgo.v2"
)

// CreateOrderPayment returns action for create OrderPayment
func CreateOrderPayment(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// UpdateOrderPayment returns action for upddate OrderPayment
func UpdateOrderPayment(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// DeleteOrderPayment returns action for delete OrderPayment
func DeleteOrderPayment(w http.ResponseWriter, h *http.Request, Session *mgo.Session) *action.Action {

	return nil
}

// ListOrderPayment returns action for list OrderPayment
func ListOrderPayment(w http.ResponseWriter, h *http.Request, Session *mgo.Session) action.Action {

	return nil
}
