package user

import (
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

// Do update user data
func (u *Update) Do() {
	if err := u.AccessValidate(); err != nil {
		u.Err = err
		return
	}

	if err := u.Validate(); err != nil {
		u.Err = err
		return
	}

	dbUser := &db.User{Session: u.Session}

	dtUser, err := dbUser.Get(bson.ObjectIdHex(u.ReqPayload.ID.ID))
	if err != nil {
		u.Err = apperr.Database{
			Where:  "User",
			Action: "Update",
			Base:   err,
		}
		return
	}

	u.ResPayload.loadToData(dtUser)
	if err := dtUser.Validate(); err != nil {
		u.Err = err
		return
	}

	dtUser.ModifiedBy = u.Who.ID

	if err := dbUser.Put(dtUser); err != nil {
		u.Err = apperr.Database{
			Base:   err,
			Where:  "User",
			Action: "Update",
		}
		return
	}

	u.ResPayload = ID{ID: string(dtUser.ID)}
}

// AccessValidate returns error
// it checks access permission
func (u *Update) AccessValidate() error {
	if !u.Role.UserAccess.Can(data.AccessUpdate) {
		return apperr.Access{Where: "User", Permission: string(data.AccessUpdate)}
	}

	return nil
}

// Validate valdes action data
func (u Update) Validate() error {
	if !bson.IsObjectIdHex(u.ReqPayload.ID.ID) {
		return apperr.Validation{
			Where: "User",
			Field: "id",
			Cause: apperr.ValidationInvalid,
		}
	}

	return nil
}
