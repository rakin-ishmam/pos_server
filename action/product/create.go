package product

import (
	"io"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/converter"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
)

// Create manages create process of the Inventory
type Create struct {
	session    *mgo.Session
	reqPayload CreatePayload
	resPayload geninfo.ID
	who        data.User
	role       data.Role
	err        error
}

// Do creates Product
func (c *Create) Do() {
	if err := c.AccessValidate(); err != nil {
		c.err = err
		return
	}

	dtProd := &data.Product{}
	c.reqPayload.LoadToData(dtProd)

	dtProd.BeforeCreate(c.who.ID)

	if err := dtProd.Validate(); err != nil {
		c.err = err
		return
	}

	dbProd := db.Product{Session: c.session}
	err := dbProd.Put(dtProd)
	if err != nil {
		c.err = apperr.Database{
			Base:   err,
			Where:  "Product",
			Action: "Create",
		}
		return
	}

	c.resPayload = geninfo.ID{ID: dtProd.ID.Hex()}
}

// Result returns result of thte action
func (c Create) Result() (io.Reader, error) {
	if c.err != nil {
		return nil, c.err
	}

	return converter.JSONtoBuff(c.resPayload)
}

// AccessValidate returns error. it checks access permission
func (c *Create) AccessValidate() error {
	if !c.role.ProductAccess.Can(data.AccessWrite) {
		return apperr.Access{Where: "Product", Permission: string(data.AccessWrite)}
	}

	return nil
}

// NewCreate returns action to create role
func NewCreate(payload CreatePayload, who data.User, role data.Role, ses *mgo.Session) *Create {
	return &Create{
		who:        who,
		role:       role,
		reqPayload: payload,
		session:    ses,
	}
}
