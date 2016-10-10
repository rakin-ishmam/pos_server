package customer

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

	dbCus := &db.Customer{Session: u.Session}

	dtCus, err := dbCus.Get(bson.ObjectIdHex(u.ReqPayload.ID.ID))
	if err != nil {
		u.Err = apperr.Database{
			Where:  "Customer",
			Action: "Update",
			Base:   err,
		}
		return
	}

	u.ResPayload.loadToData(dtCus)
	if err := dtCus.Validate(); err != nil {
		u.Err = err
		return
	}

	dtCus.ModifiedBy = u.Who.ID
	dtCus.ModifiedAt = time.Now()

	if err := dbCus.Put(dtCus); err != nil {
		u.Err = apperr.Database{
			Base:   err,
			Where:  "Customer",
			Action: "Update",
		}
		return
	}

	u.ResPayload = ID{ID: string(dtCus.ID)}
}

// AccessValidate returns error
// it checks access permission
func (u Update) AccessValidate() error {
	if !u.Role.CustomerAccess.Can(data.AccessUpdate) {
		return apperr.Access{Where: "Customer", Permission: string(data.AccessUpdate)}
	}

	return nil
}

// Validate valdes action data
func (u Update) Validate() error {
	if !bson.IsObjectIdHex(u.ReqPayload.ID.ID) {
		return apperr.Validation{
			Where: "Customer",
			Field: "id",
			Cause: apperr.ValidationInvalid,
		}
	}

	return nil
}
