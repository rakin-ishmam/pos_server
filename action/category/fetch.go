package category

import (
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

// Do takes necessary steps to fetch Category data
func (f *Fetch) Do() {
	if err := f.AccessValidate(); err != nil {
		f.Err = err
		return
	}

	if err := f.Validate(); err != nil {
		f.Err = err
		return
	}

	dbCat := db.Category{Session: f.Session}
	dtCat, err := dbCat.Get(bson.ObjectIdHex(f.ReqID))
	if err != nil {
		f.Err = apperr.Database{
			Base:   err,
			Where:  "Category",
			Action: "Fetch",
		}
		return
	}

	f.ResPayload = DetailPayload{}
	f.ResPayload.LoadFromData(dtCat)
}

// ActionErr returns error of the action
func (f Fetch) ActionErr() error {
	return f.Err
}

// Result returns result of thte action
func (f Fetch) Result() interface{} {
	return f.ResPayload
}

// AccessValidate checks access permission
func (f Fetch) AccessValidate() error {
	if !f.Role.CategoryAccess.Can(data.AccessRead) {
		return apperr.Access{Where: "Category", Permission: string(data.AccessRead)}
	}

	return nil
}

// Validate valids action data
func (f Fetch) Validate() error {
	if !bson.IsObjectIdHex(f.ReqID) {
		return apperr.Validation{
			Where: "Category",
			Field: "id",
			Cause: apperr.ValidationInvalid,
		}
	}
	return nil
}
