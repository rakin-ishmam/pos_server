package role

import (
	"io"

	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/converter"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Fetch maneges fetch payload process
type Fetch struct {
	session    *mgo.Session
	who        data.User
	role       data.Role
	reqID      string
	resPayload *DetailPayload
	err        error
}

// Do fetch Role by ReqID
func (f *Fetch) Do() {
	if err := f.AccessValidate(); err != nil {
		f.err = err
		return
	}

	if err := f.Validate(); err != nil {
		f.err = err
		return
	}

	roleDB := db.Role{Session: f.session}

	dtRole, err := roleDB.Get(bson.ObjectIdHex(f.reqID))
	if err != nil {
		f.err = apperr.Database{
			Base:   err,
			Where:  "Role",
			Action: "Fetch",
		}
		return
	}

	f.resPayload = &DetailPayload{}
	f.resPayload.LoadFromData(dtRole)
}

// Result returns result of thte action
func (f Fetch) Result() (io.Reader, error) {
	if f.err != nil {
		return nil, f.err
	}

	return converter.JSONtoBuff(f.resPayload)
}

// AccessValidate returns error. it checkes access permission
func (f *Fetch) AccessValidate() error {
	if !f.role.RoleAccess.Can(data.AccessRead) {
		return apperr.Access{Where: "Role", Permission: string(data.AccessWrite)}
	}

	return nil
}

// Validate valids action data
func (f Fetch) Validate() error {
	if !bson.IsObjectIdHex(f.reqID) {
		return apperr.Validation{
			Where: "Role",
			Field: "id",
			Cause: apperr.StrInvalid,
		}
	}
	return nil
}

// NewFetch return role fetch action
func NewFetch(id string, who data.User, role data.Role, ses *mgo.Session) *Fetch {
	return &Fetch{
		reqID:   id,
		who:     who,
		role:    role,
		session: ses,
	}
}
