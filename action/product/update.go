package product

import (
	"time"

	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Update manages update process of user
type Update struct {
	Session    *mgo.Session
	ReqPayload UpdatePayload
	ResPayload ID
	Who        data.User
	Role       data.Role
	Err        error
}

// Do update Category data
func (u *Update) Do() {
	if err := u.AccessValidate(); err != nil {
		u.Err = err
		return
	}

	if err := u.Validate(); err != nil {
		u.Err = err
		return
	}

	dbProd := &db.Product{Session: u.Session}

	dtProd, err := dbProd.Get(bson.ObjectIdHex(u.ReqPayload.ID.ID))
	if err != nil {
		u.Err = apperr.Database{
			Where:  "Product",
			Action: "Update",
			Base:   err,
		}
		return
	}

	u.ResPayload.loadToData(dtProd)
	if err := dtProd.Validate(); err != nil {
		u.Err = err
		return
	}

	dtProd.ModifiedBy = u.Who.ID
	dtProd.ModifiedAt = time.Now()

	if err := dbProd.Put(dtProd); err != nil {
		u.Err = apperr.Database{
			Base:   err,
			Where:  "Product",
			Action: "Update",
		}
		return
	}

	u.ResPayload = ID{ID: string(dtProd.ID)}
}

// AccessValidate returns error
// it checks access permission
func (u *Update) AccessValidate() error {
	if !u.Role.ProductAccess.Can(data.AccessUpdate) {
		return apperr.Access{Where: "Product", Permission: string(data.AccessUpdate)}
	}

	return nil
}

// Validate valdes action data
func (u Update) Validate() error {
	if !bson.IsObjectIdHex(u.ReqPayload.ID.ID) {
		return apperr.Validation{
			Where: "Product",
			Field: "id",
			Cause: apperr.ValidationInvalid,
		}
	}

	return nil
}
