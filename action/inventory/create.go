package inventory

import (
	"time"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
)

// Create manages create process of the Inventory
type Create struct {
	Session    *mgo.Session
	ReqPayload CreatePayload
	ResPayload geninfo.ID
	Who        data.User
	Role       data.Role
	Err        error
}

// Do creates Inventory
func (c *Create) Do() {
	if err := c.AccessValidate(); err != nil {
		c.Err = err
		return
	}

	dtInv := &data.Inventory{}
	c.ReqPayload.LoadToData(dtInv)

	dtInv.CreatedBy = c.Who.ID
	dtInv.ModifiedBy = c.Who.ID
	dtInv.CreatedAt = time.Now()
	dtInv.ModifiedAt = dtInv.CreatedAt

	if err := c.Validate(dtInv); err != nil {
		c.Err = err
		return
	}

	dbInvt := db.Inventory{Session: c.Session}
	err := dbInvt.Put(dtInv)
	if err != nil {
		c.Err = apperr.Database{
			Base:   err,
			Where:  "Inventory",
			Action: "Create",
		}
		return
	}

	c.ResPayload = geninfo.ID{ID: string(dtInv.ID)}
}

// ActionErr returns error of the action
func (c Create) ActionErr() error {
	return c.Err
}

// Result returns result of thte action
func (c Create) Result() interface{} {
	return c.ResPayload
}

// Validate takes Inventory data and returns error
// it checks data validation
func (c *Create) Validate(dtInv *data.Inventory) error {
	return dtInv.Validate()
}

// AccessValidate returns error. it checks access permission
func (c *Create) AccessValidate() error {
	if !c.Role.InventoryAccess.Can(data.AccessWrite) {
		return apperr.Access{Where: "Inventory", Permission: string(data.AccessWrite)}
	}

	return nil
}
