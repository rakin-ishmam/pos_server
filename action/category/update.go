package category

import (
	"io"
	"time"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/converter"
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

	dbCat := &db.Category{Session: u.Session}

	dtCat, err := dbCat.Get(bson.ObjectIdHex(u.ReqPayload.ID.ID))
	if err != nil {
		u.Err = apperr.Database{
			Where:  "Category",
			Action: "Update",
			Base:   err,
		}
		return
	}

	u.ResPayload.LoadToData(&dtCat.Track)
	if err := dtCat.Validate(); err != nil {
		u.Err = err
		return
	}

	dtCat.ModifiedBy = u.Who.ID
	dtCat.ModifiedAt = time.Now()

	if err := dbCat.Put(dtCat); err != nil {
		u.Err = apperr.Database{
			Base:   err,
			Where:  "Category",
			Action: "Update",
		}
		return
	}

	u.ResPayload = geninfo.ID{ID: string(dtCat.ID)}
}

// Result returns result of thte action
func (u Update) Result() (io.Reader, error) {
	if u.Err != nil {
		return nil, u.Err
	}

	return converter.JSONtoBuff(u.ResPayload)
}

// AccessValidate returns error
// it checks access permission
func (u *Update) AccessValidate() error {
	if !u.Role.CategoryAccess.Can(data.AccessUpdate) {
		return apperr.Access{Where: "Category", Permission: string(data.AccessUpdate)}
	}

	return nil
}

// Validate valdes action data
func (u Update) Validate() error {
	if !bson.IsObjectIdHex(u.ReqPayload.ID.ID) {
		return apperr.Validation{
			Where: "Category",
			Field: "id",
			Cause: apperr.StrInvalid,
		}
	}

	return nil
}
