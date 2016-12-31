package api

import (
	"net/http"

	"github.com/rakin-ishmam/pos_server/action"
	mgo "gopkg.in/mgo.v2"
)

// CreateOrderPayment returns action for create OrderPayment
func CreateOrderPayment(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// UpdateOrderPayment returns action for upddate OrderPayment
func UpdateOrderPayment(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// DeleteOrderPayment returns action for delete OrderPayment
func DeleteOrderPayment(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}

// ListOrderPayment returns action for list OrderPayment
func ListOrderPayment(w http.ResponseWriter, r *http.Request, Session *mgo.Session) action.JSONAction {

	return nil
}
