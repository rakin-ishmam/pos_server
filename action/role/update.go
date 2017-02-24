package role

import (
	"io"

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
	session    *mgo.Session
	who        data.User
	role       data.Role
	reqPayload UpdatePayload
	resPayload geninfo.ID
	err        error
}

// Do updates data.Role
func (u *Update) Do() {
	if err := u.AccessValidate(); err != nil {
		u.err = err
		return
	}

	if err := u.Validate(); err != nil {
		u.err = err
		return
	}

	roleDB := db.Role{Session: u.session}

	dtRole, err := roleDB.Get(bson.ObjectIdHex(u.reqPayload.ID.ID))
	if err != nil {
		u.err = apperr.Database{
			Base:   err,
			Where:  "Role",
			Action: "Update",
		}
		return
	}

	u.reqPayload.LoadToData(dtRole)

	if err := dtRole.Validate(); err != nil {
		u.err = err
		return
	}

	dtRole.BeforeUpdate(u.who.ID)

	if err := roleDB.Put(dtRole); err != nil {
		u.err = apperr.Database{
			Base:   err,
			Where:  "Role",
			Action: "Update",
		}
		return
	}

	u.resPayload = geninfo.ID{ID: dtRole.ID.Hex()}
}

// Result returns result of thte action
func (u Update) Result() (io.Reader, error) {
	if u.err != nil {
		return nil, u.err
	}

	return converter.JSONtoBuff(u.resPayload)
}

// AccessValidate checks access permission
func (u *Update) AccessValidate() error {
	if !u.role.RoleAccess.Can(data.AccessUpdate) {
		return apperr.Access{Where: "Role", Permission: string(data.AccessUpdate)}
	}

	return nil
}

// Validate valdes action data
func (u Update) Validate() error {
	if !bson.IsObjectIdHex(u.reqPayload.ID.ID) {
		return apperr.Validation{
			Where: "Role",
			Field: "id",
			Cause: apperr.StrInvalid,
		}
	}

	return nil
}

// NewUpdate returns user update action
func NewUpdate(payload UpdatePayload, who data.User, role data.Role, ses *mgo.Session) *Update {
	return &Update{
		reqPayload: payload,
		who:        who,
		role:       role,
		session:    ses,
	}
}
