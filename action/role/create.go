package role

import (
	"time"

	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
)

// Create manages Role create process
type Create struct {
	Session    *mgo.Session
	Who        data.User
	Role       data.Role
	ReqPayload CreatePayload
	ResPayload ID
	Err        error
}

// Do insert new Role
func (c *Create) Do() {
	if err := c.AccessValidate(); err != nil {
		c.Err = err
		return
	}

	dtRole := &data.Role{}
	c.ReqPayload.loadToData(dtRole)

	dtRole.CreatedBy = c.Who.ID
	dtRole.ModifiedBy = c.Who.ID
	dtRole.CreatedAt = time.Now()
	dtRole.ModifiedAt = dtRole.CreatedAt

	if err := dtRole.Validate(); err != nil {
		c.Err = err
		return
	}

	roleDB := db.Role{Session: c.Session}
	err := roleDB.Put(dtRole)
	if err != nil {
		c.Err = apperr.Database{
			Base:   err,
			Where:  "Role",
			Action: "Create",
		}
		return
	}

	c.ResPayload = ID{ID: string(dtRole.ID)}
}

// AccessValidate checks access permission
func (c *Create) AccessValidate() error {
	if !c.Role.RoleAccess.Can(data.AccessWrite) {
		return apperr.Access{Where: "Role", Permission: string(data.AccessWrite)}
	}

	return nil
}
