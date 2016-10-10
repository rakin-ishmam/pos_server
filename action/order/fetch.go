package order

import (
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	r "gopkg.in/dancannon/gorethink.v2"
)

// Fetch store fetch request data and store data
type Fetch struct {
	Session    *mgo.Session
	ReqID      string
	ResPayload DetailPayload
	Who        data.User
	Role       data.Role
	Err        error
}

// Do takes necessary steps to fetch Inventory data
func (f *Fetch) Do() {
	if err := f.AccessValidate(); err != nil {
		f.Err = err
		return
	}

	dbInv := db.Inventory{DB: db.DB{Session: f.Session}}
	dtInv, err := dbInv.Get(f.ReqID)
	if err != nil {
		f.Err = apperr.Database{
			Base:   err,
			Where:  "Inventory",
			Action: "Fetch",
		}
		return
	}

	f.ResPayload = DetailPayload{}
	f.ResPayload.loadFromData(dtInv)
}

// AccessValidate checks access permission
func (f *Fetch) AccessValidate() error {
	if !f.Role.InventoryAccess.Can(data.AccessRead) {
		return apperr.Access{Where: "Inventory", Permission: string(data.AccessRead)}
	}

	return nil
}
