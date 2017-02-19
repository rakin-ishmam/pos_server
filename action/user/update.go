package user

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

// Update manages update process of user
type Update struct {
	session    *mgo.Session
	reqPayload UpdatePayload
	resPayload geninfo.ID
	who        data.User
	role       data.Role
	err        error
}

// Do update user data
func (u *Update) Do() {
	if err := u.AccessValidate(); err != nil {
		u.err = err
		return
	}

	if err := u.Validate(); err != nil {
		u.err = err
		return
	}

	dbUser := &db.User{Session: u.session}

	dtUser, err := dbUser.Get(bson.ObjectIdHex(u.reqPayload.ID.ID))
	if err != nil {
		u.err = apperr.Database{
			Where:  "User",
			Action: "Update",
			Base:   err,
		}
		return
	}

	u.resPayload.LoadToData(&dtUser.Track)
	if err := dtUser.Validate(); err != nil {
		u.err = err
		return
	}

	dtUser.ModifiedBy = u.who.ID

	if err := dbUser.Put(dtUser); err != nil {
		u.err = apperr.Database{
			Base:   err,
			Where:  "User",
			Action: "Update",
		}
		return
	}

	u.resPayload = geninfo.ID{ID: string(dtUser.ID)}
}

// Result returns result of thte action
func (u Update) Result() (io.Reader, error) {
	if u.err != nil {
		return nil, u.err
	}

	return converter.JSONtoBuff(u.resPayload)
}

// AccessValidate returns error
// it checks access permission
func (u *Update) AccessValidate() error {
	if !u.role.UserAccess.Can(data.AccessUpdate) {
		return apperr.Access{Where: "User", Permission: string(data.AccessUpdate)}
	}

	return nil
}

// Validate valdes action data
func (u Update) Validate() error {
	if !bson.IsObjectIdHex(u.reqPayload.ID.ID) {
		return apperr.Validation{
			Where: "User",
			Field: "id",
			Cause: apperr.StrInvalid,
		}
	}

	return nil
}

// NewUpdateAction returns user Update action
func NewUpdateAction(reqPayload UpdatePayload, who data.User, role data.Role, session *mgo.Session) *Update {
	return &Update{
		reqPayload: reqPayload,
		who:        who,
		role:       role,
		session:    session,
	}
}
