package file

import (
	"time"

	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
)

// Create manages create process of the File
type Create struct {
	Session    *mgo.Session
	ReqPayload CreatePayload
	ResPayload ID
	Who        data.User
	Role       data.Role
	Err        error
}

// Do creates File
func (c *Create) Do() {
	if err := c.AccessValidate(); err != nil {
		c.Err = err
		return
	}

	dtFile := &data.File{}
	c.ReqPayload.loadToData(dtFile)

	dtFile.CreatedBy = c.Who.ID
	dtFile.ModifiedBy = c.Who.ID
	dtFile.CreatedAt = time.Now()
	dtFile.ModifiedAt = dtFile.CreatedAt

	if err := dtFile.Validate(); err != nil {
		c.Err = err
		return
	}

	dbFile := db.File{Session: c.Session}
	err := dbFile.Put(dtFile)
	if err != nil {
		c.Err = apperr.Database{
			Base:   err,
			Where:  "File",
			Action: "Create",
		}
		return
	}

	c.ResPayload = ID{ID: string(dtFile.ID)}
}

// AccessValidate returns error. it checks access permission
func (c *Create) AccessValidate() error {
	if !c.Role.FileAccess.Can(data.AccessWrite) {
		return apperr.Access{Where: "File", Permission: string(data.AccessWrite)}
	}

	return nil
}
