package category

import (
	"io"
	"time"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/converter"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
)

// Create manages create process of Category
type Create struct {
	Session    *mgo.Session
	ReqPayload CreatePayload
	ResPayload geninfo.ID
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

	dtCat := &data.Category{}
	c.ReqPayload.LoadToData(dtCat)

	dtCat.CreatedBy = c.Who.ID
	dtCat.ModifiedBy = c.Who.ID
	dtCat.CreatedAt = time.Now()
	dtCat.ModifiedAt = dtCat.CreatedAt

	if err := dtCat.Validate(); err != nil {
		c.Err = err
		return
	}

	dbCat := db.Category{Session: c.Session}
	err := dbCat.Put(dtCat)
	if err != nil {
		c.Err = apperr.Database{
			Base:   err,
			Where:  "User",
			Action: "Create",
		}
		return
	}

	c.ResPayload = geninfo.ID{ID: string(dtCat.ID)}
}

// Result returns result of thte action
func (c Create) Result() (io.Reader, error) {
	if c.Err != nil {
		return nil, c.Err
	}

	return converter.JSONtoBuff(c.ResPayload)
}

// AccessValidate returns error. it checks access permission
func (c Create) AccessValidate() error {
	if !c.Role.CategoryAccess.Can(data.AccessWrite) {
		return apperr.Access{Where: "Category", Permission: string(data.AccessWrite)}
	}

	return nil
}
