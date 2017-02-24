package api

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/rakin-ishmam/pos_server/action"
	"github.com/rakin-ishmam/pos_server/action/empty"
	"github.com/rakin-ishmam/pos_server/action/product"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db/query"
	mgo "gopkg.in/mgo.v2"
)

// CreateProduct returns action for create Product
func CreateProduct(w http.ResponseWriter, r *http.Request, session *mgo.Session) action.JSONAction {
	createDt := product.CreatePayload{}
	if errAc := jsonDecode(r, &createDt, "CreateProduct", "create payload"); errAc != nil {
		return errAc
	}

	usr := context.Get(r, "user").(data.User)
	rl := context.Get(r, "role").(data.Role)

	return product.NewCreate(createDt, usr, rl, session)
}

// UpdateProduct returns action for upddate Product
func UpdateProduct(w http.ResponseWriter, r *http.Request, session *mgo.Session) action.JSONAction {
	id, errAc := idFetch(r, "UpdateProduct")
	if errAc != nil {
		return errAc
	}

	updateDt := product.UpdatePayload{}

	if errAc = jsonDecode(r, &updateDt, "UpdateRole", "update payload"); errAc != nil {
		return errAc
	}

	updateDt.ID.ID = id

	usr := context.Get(r, "user").(data.User)
	rl := context.Get(r, "role").(data.Role)

	return product.NewUpdate(updateDt, usr, rl, session)
}

// DeleteProduct returns action for delete Product
func DeleteProduct(w http.ResponseWriter, r *http.Request, session *mgo.Session) action.JSONAction {
	return nil
}

// ListProduct returns action for list Product
func ListProduct(w http.ResponseWriter, r *http.Request, session *mgo.Session) action.JSONAction {
	qp := query.Product{}

	err := urlquery(&qp.GenInfo, r, urlQSkip("ListProduct"), urlQLimit("ListProduct"))
	if err != nil {
		return empty.NewJSON(err)
	}

	usr := context.Get(r, "user").(data.User)
	rl := context.Get(r, "role").(data.Role)

	return product.NewList(session, qp, usr, rl)
}

// FetchProduct returns action to get list of role
func FetchProduct(w http.ResponseWriter, r *http.Request, session *mgo.Session) action.JSONAction {
	id, errAc := idFetch(r, "FetchProduct")
	if errAc != nil {
		return errAc
	}

	usr := context.Get(r, "user").(data.User)
	rl := context.Get(r, "role").(data.Role)

	return product.NewFetch(id, usr, rl, session)
}
