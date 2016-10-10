package inventory

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

	dbInv := &db.Inventory{Session: u.Session}

	dtInv, err := dbInv.Get(bson.ObjectIdHex(u.ReqPayload.ID.ID))
	if err != nil {
		u.Err = apperr.Database{
			Where:  "Inventory",
			Action: "Update",
			Base:   err,
		}
		return
	}

	u.ResPayload.loadToData(dtInv)
	if err := dtInv.Validate(); err != nil {
		u.Err = err
		return
	}

	dtInv.ModifiedBy = u.Who.ID
	dtInv.ModifiedAt = time.Now()

	if err := dbInv.Put(dtInv); err != nil {
		u.Err = apperr.Database{
			Base:   err,
			Where:  "Inventory",
			Action: "Update",
		}
		return
	}

	u.ResPayload = ID{ID: string(dtInv.ID)}
}

// AccessValidate returns error
// it checks access permission
func (u Update) AccessValidate() error {
	if !u.Role.InventoryAccess.Can(data.AccessUpdate) {
		return apperr.Access{Where: "Inventory", Permission: string(data.AccessUpdate)}
	}

	return nil
}

// Validate valdes action data
func (u Update) Validate() error {
	if !bson.IsObjectIdHex(u.ReqPayload.ID.ID) {
		return apperr.Validation{
			Where: "Inventory",
			Field: "id",
			Cause: apperr.ValidationInvalid,
		}
	}

	return nil
}
