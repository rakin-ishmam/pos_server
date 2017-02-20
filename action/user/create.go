package user

import (
	"io"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/converter"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
)

// Create manages create process of user
type Create struct {
	session    *mgo.Session
	reqPayload CreatePayload
	resPayload geninfo.ID
	who        data.User
	role       data.Role
	err        error
}

// Do creates user
func (c *Create) Do() {
	if err := c.AccessValidate(); err != nil {
		c.err = err
		return
	}

	dtUser := &data.User{}
	c.reqPayload.LoadToData(dtUser)

	dtUser.BeforeCreate(c.who.ID)

	if err := dtUser.Validate(); err != nil {
		c.err = err
		return
	}

	dbUser := db.User{Session: c.session}
	err := dbUser.Put(dtUser)
	if err != nil {
		c.err = apperr.Database{
			Base:   err,
			Where:  "User",
			Action: "Create",
		}
		return
	}

	c.resPayload = geninfo.ID{ID: dtUser.ID.Hex()}
}

// Result returns result of the action
func (c Create) Result() (io.Reader, error) {
	if c.err != nil {
		return nil, c.err
	}

	return converter.JSONtoBuff(c.resPayload)
}

// AccessValidate returns error. it checks access permission
func (c *Create) AccessValidate() error {
	if !c.role.UserAccess.Can(data.AccessWrite) {
		return apperr.Access{Where: "User", Permission: string(data.AccessWrite)}
	}

	return nil
}

// NewCreate returns user create action
func NewCreate(reqPayload CreatePayload, who data.User, role data.Role, ses *mgo.Session) *Create {
	return &Create{
		who:        who,
		role:       role,
		session:    ses,
		reqPayload: reqPayload,
	}
}
