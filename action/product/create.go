package product

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

// Create manages create process of the Inventory
type Create struct {
	Session    *mgo.Session
	ReqPayload CreatePayload
	ResPayload geninfo.ID
	Who        data.User
	Role       data.Role
	Err        error
}

// Do creates Product
func (c *Create) Do() {
	if err := c.AccessValidate(); err != nil {
		c.Err = err
		return
	}

	dtProd := &data.Product{}
	c.ReqPayload.LoadToData(dtProd)

	dtProd.CreatedBy = c.Who.ID
	dtProd.ModifiedBy = c.Who.ID
	dtProd.CreatedAt = time.Now()
	dtProd.ModifiedAt = dtProd.CreatedAt

	dtProd.ProductType = data.LocalProductType

	if err := dtProd.Validate(); err != nil {
		c.Err = err
		return
	}

	dbProd := db.Product{Session: c.Session}
	err := dbProd.Put(dtProd)
	if err != nil {
		c.Err = apperr.Database{
			Base:   err,
			Where:  "Product",
			Action: "Create",
		}
		return
	}

	c.ResPayload = geninfo.ID{ID: string(dtProd.ID)}
}

// Result returns result of thte action
func (c Create) Result() (io.Reader, error) {
	if c.Err != nil {
		return nil, c.Err
	}

	return converter.JSONtoBuff(c.ResPayload)
}

// AccessValidate returns error. it checks access permission
func (c *Create) AccessValidate() error {
	if !c.Role.ProductAccess.Can(data.AccessWrite) {
		return apperr.Access{Where: "Product", Permission: string(data.AccessWrite)}
	}

	return nil
}
