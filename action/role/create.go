package role

import (
	"io"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/converter"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
)

// Create manages Role create process
type Create struct {
	session    *mgo.Session
	who        data.User
	role       data.Role
	reqPayload CreatePayload
	resPayload geninfo.ID
	err        error
}

// Do insert new Role
func (c *Create) Do() {
	if err := c.AccessValidate(); err != nil {
		c.err = err
		return
	}

	dtRole := &data.Role{}
	c.reqPayload.LoadToData(dtRole)

	dtRole.BeforeCreate(c.who.ID)

	if err := dtRole.Validate(); err != nil {
		c.err = err
		return
	}

	roleDB := db.Role{Session: c.session}
	err := roleDB.Put(dtRole)
	if err != nil {
		c.err = apperr.Database{
			Base:   err,
			Where:  "Role",
			Action: "Create",
		}
		return
	}

	c.resPayload = geninfo.ID{ID: dtRole.ID.Hex()}
}

// Result returns result of thte action
func (c Create) Result() (io.Reader, error) {
	if c.err != nil {
		return nil, c.err
	}

	return converter.JSONtoBuff(c.resPayload)
}

// AccessValidate checks access permission
func (c *Create) AccessValidate() error {
	if !c.role.RoleAccess.Can(data.AccessWrite) {
		return apperr.Access{Where: "Role", Permission: string(data.AccessWrite)}
	}

	return nil
}

// NewCreate returns create action for role
func NewCreate(payload CreatePayload, who data.User, role data.Role, ses *mgo.Session) *Create {
	return &Create{
		who:        who,
		role:       role,
		reqPayload: payload,
		session:    ses,
	}
}
