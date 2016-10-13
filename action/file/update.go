package file

import (
	"time"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
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
	ResPayload geninfo.ID
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

	dbFile := &db.File{Session: u.Session}

	dtFile, err := dbFile.Get(bson.ObjectIdHex(u.ReqPayload.ID.ID))
	if err != nil {
		u.Err = apperr.Database{
			Where:  "File",
			Action: "Update",
			Base:   err,
		}
		return
	}

	u.ResPayload.LoadToData(&dtFile.Track)
	if err := dtFile.Validate(); err != nil {
		u.Err = err
		return
	}

	dtFile.ModifiedBy = u.Who.ID
	dtFile.ModifiedAt = time.Now()

	if err := dbFile.Put(dtFile); err != nil {
		u.Err = apperr.Database{
			Base:   err,
			Where:  "File",
			Action: "Update",
		}
		return
	}

	u.ResPayload = geninfo.ID{ID: string(dtFile.ID)}
}

// ActionErr returns error of the action
func (u Update) ActionErr() error {
	return u.Err
}

// Result returns result of thte action
func (u Update) Result() interface{} {
	return u.ResPayload
}

// AccessValidate returns error
// it checks access permission
func (u *Update) AccessValidate() error {
	if !u.Role.FileAccess.Can(data.AccessUpdate) {
		return apperr.Access{Where: "File", Permission: string(data.AccessUpdate)}
	}

	return nil
}

// Validate valdes action data
func (u Update) Validate() error {
	if !bson.IsObjectIdHex(u.ReqPayload.ID.ID) {
		return apperr.Validation{
			Where: "File",
			Field: "id",
			Cause: apperr.ValidationInvalid,
		}
	}

	return nil
}
