package customer

import (
	"time"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
)

// Create manages create process of the Customer
type Create struct {
	Session    *mgo.Session
	ReqPayload CreatePayload
	ResPayload geninfo.ID
	Who        data.User
	Role       data.Role
	Err        error
}

// Do creates Customer
func (c *Create) Do() {
	if err := c.AccessValidate(); err != nil {
		c.Err = err
		return
	}

	dtCus := &data.Customer{}
	c.ReqPayload.LoadToData(dtCus)

	dtCus.CreatedBy = c.Who.ID
	dtCus.ModifiedBy = c.Who.ID
	dtCus.CreatedAt = time.Now()
	dtCus.ModifiedAt = dtCus.CreatedAt

	if err := dtCus.Validate(); err != nil {
		c.Err = err
		return
	}

	dbCus := db.Customer{Session: c.Session}
	err := dbCus.Put(dtCus)
	if err != nil {
		c.Err = apperr.Database{
			Base:   err,
			Where:  "Customer",
			Action: "Create",
		}
		return
	}

	c.ResPayload = geninfo.ID{ID: string(dtCus.ID)}
}

// ActionErr returns error of the action
func (c Create) ActionErr() error {
	return c.Err
}

// Result returns result of thte action
func (c Create) Result() interface{} {
	return &c.ResPayload
}

// AccessValidate returns error. it checks access permission
func (c *Create) AccessValidate() error {
	if !c.Role.CustomerAccess.Can(data.AccessWrite) {
		return apperr.Access{Where: "Customer", Permission: string(data.AccessWrite)}
	}

	return nil
}
