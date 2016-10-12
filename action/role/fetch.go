package role

import (
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Fetch maneges fetch payload process
type Fetch struct {
	Session    *mgo.Session
	Who        data.User
	Role       data.Role
	ReqID      string
	ResPayload *DetailPayload
	Err        error
}

// Do fetch Role by ReqID
func (f *Fetch) Do() {
	if err := f.AccessValidate(); err != nil {
		f.Err = err
		return
	}

	if err := f.Validate(); err != nil {
		f.Err = err
		return
	}

	roleDB := db.Role{Session: f.Session}

	dtRole, err := roleDB.Get(bson.ObjectIdHex(f.ReqID))
	if err != nil {
		f.Err = apperr.Database{
			Base:   err,
			Where:  "Role",
			Action: "Fetch",
		}
		return
	}

	f.ResPayload = &DetailPayload{}
	f.ResPayload.LoadFromData(dtRole)
}

// AccessValidate returns error. it checkes access permission
func (f *Fetch) AccessValidate() error {
	if !f.Role.RoleAccess.Can(data.AccessRead) {
		return apperr.Access{Where: "Role", Permission: string(data.AccessWrite)}
	}

	return nil
}

// Validate valids action data
func (f Fetch) Validate() error {
	if !bson.IsObjectIdHex(f.ReqID) {
		return apperr.Validation{
			Where: "Role",
			Field: "id",
			Cause: apperr.ValidationInvalid,
		}
	}
	return nil
}
