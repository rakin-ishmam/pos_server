package role

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

// Update manages Role update process
type Update struct {
	Session    *mgo.Session
	Who        data.User
	Role       data.Role
	ReqPayload UpdatePayload
	ResPayload geninfo.ID
	Err        error
}

// Do updates data.Role
func (u *Update) Do() {
	if err := u.AccessValidate(); err != nil {
		u.Err = err
		return
	}

	if err := u.Validate(); err != nil {
		u.Err = err
		return
	}

	roleDB := db.Role{Session: u.Session}

	dtRole, err := roleDB.Get(bson.ObjectIdHex(u.ReqPayload.ID.ID))
	if err != nil {
		u.Err = apperr.Database{
			Base:   err,
			Where:  "Role",
			Action: "Update",
		}
		return
	}

	u.ReqPayload.LoadToData(dtRole)

	if err := dtRole.Validate(); err != nil {
		u.Err = err
		return
	}

	dtRole.ModifiedBy = u.Who.ID
	dtRole.ModifiedAt = time.Now()

	if err := roleDB.Put(dtRole); err != nil {
		u.Err = apperr.Database{
			Base:   err,
			Where:  "Role",
			Action: "Update",
		}
		return
	}

	u.ResPayload = geninfo.ID{ID: string(dtRole.ID)}
}

// Result returns result of thte action
func (u Update) Result() (io.Reader, error) {
	if u.Err != nil {
		return nil, u.Err
	}

	return converter.JSONtoBuff(u.ResPayload)
}

// AccessValidate checks access permission
func (u *Update) AccessValidate() error {
	if !u.Role.RoleAccess.Can(data.AccessUpdate) {
		return apperr.Access{Where: "Role", Permission: string(data.AccessUpdate)}
	}

	return nil
}

// Validate valdes action data
func (u Update) Validate() error {
	if !bson.IsObjectIdHex(u.ReqPayload.ID.ID) {
		return apperr.Validation{
			Where: "Role",
			Field: "id",
			Cause: apperr.ValidationInvalid,
		}
	}

	return nil
}
