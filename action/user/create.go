package user

import (
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
)

// Create manages create process of user
type Create struct {
	Session    *mgo.Session
	ReqPayload CreatePayload
	ResPayload ID
	Who        data.User
	Role       data.Role
	Err        error
}

// Do creates user
func (c *Create) Do() {
	if err := c.AccessValidate(); err != nil {
		c.Err = err
		return
	}

	dtUser := &data.User{}
	c.ReqPayload.loadToData(dtUser)

	dtUser.CreatedBy = c.Who.ID
	dtUser.ModifiedBy = c.Who.ID

	if err := dtUser.Validate(); err != nil {
		c.Err = err
		return
	}

	dbUser := db.User{Session: c.Session}
	err := dbUser.Put(dtUser)
	if err != nil {
		c.Err = apperr.Database{
			Base:   err,
			Where:  "User",
			Action: "Create",
		}
		return
	}

	c.ResPayload = ID{ID: string(dtUser.ID)}
}

// AccessValidate returns error. it checks access permission
func (c *Create) AccessValidate() error {
	if !c.Role.UserAccess.Can(data.AccessWrite) {
		return apperr.Access{Where: "User", Permission: string(data.AccessWrite)}
	}

	return nil
}
