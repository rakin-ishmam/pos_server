package order

import (
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	r "gopkg.in/dancannon/gorethink.v2"
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

	dbInv := &db.Inventory{DB: db.DB{Session: u.Session}}

	dtInv, err := dbInv.Get(u.ReqPayload.ID.ID)
	if err != nil {
		u.Err = apperr.Database{
			Where:  "Inventory",
			Action: "Update",
			Base:   err,
		}
		return
	}

	u.ResPayload.loadToData(dtInv)
	if err := u.Validate(dtInv); err != nil {
		u.Err = err
		return
	}

	dtInv.ModifiedBy = u.Who.ID

	if err := dbInv.Put(dtInv); err != nil {
		u.Err = apperr.Database{
			Base:   err,
			Where:  "Inventory",
			Action: "Update",
		}
		return
	}

	u.ResPayload = ID{ID: dtInv.ID}
}

// AccessValidate returns error
// it checks access permission
func (u *Update) AccessValidate() error {
	if !u.Role.InventoryAccess.Can(data.AccessUpdate) {
		return apperr.Access{Where: "Inventory", Permission: string(data.AccessUpdate)}
	}

	return nil
}

// Validate takes Inventory data and returns error
// it valids File data
func (u *Update) Validate(data *data.Inventory) error {
	return data.Validate()
}
