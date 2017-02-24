package product

import (
	"io"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/converter"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Update manages update process of user
type Update struct {
	session    *mgo.Session
	reqPayload UpdatePayload
	resPayload geninfo.ID
	who        data.User
	role       data.Role
	err        error
}

// Do update Category data
func (u *Update) Do() {
	if err := u.AccessValidate(); err != nil {
		u.err = err
		return
	}

	if err := u.Validate(); err != nil {
		u.err = err
		return
	}

	dbProd := &db.Product{Session: u.session}

	dtProd, err := dbProd.Get(bson.ObjectIdHex(u.reqPayload.ID.ID))
	if err != nil {
		u.err = apperr.Database{
			Where:  "Product",
			Action: "Update",
			Base:   err,
		}
		return
	}

	u.resPayload.LoadToData(&dtProd.Track)
	if err := dtProd.Validate(); err != nil {
		u.err = err
		return
	}

	dtProd.BeforeUpdate(u.who.ID)

	if err := dbProd.Put(dtProd); err != nil {
		u.err = apperr.Database{
			Base:   err,
			Where:  "Product",
			Action: "Update",
		}
		return
	}

	u.resPayload = geninfo.ID{ID: dtProd.ID.Hex()}
}

// Result returns result of thte action
func (u Update) Result() (io.Reader, error) {
	if u.err != nil {
		return nil, u.err
	}

	return converter.JSONtoBuff(u.resPayload)
}

// AccessValidate returns error
// it checks access permission
func (u *Update) AccessValidate() error {
	if !u.role.ProductAccess.Can(data.AccessUpdate) {
		return apperr.Access{Where: "Product", Permission: string(data.AccessUpdate)}
	}

	return nil
}

// Validate valdes action data
func (u Update) Validate() error {
	if !bson.IsObjectIdHex(u.reqPayload.ID.ID) {
		return apperr.Validation{
			Where: "Product",
			Field: "id",
			Cause: apperr.StrInvalid,
		}
	}

	return nil
}

// NewUpdate returns action to update role
func NewUpdate(payload UpdatePayload, who data.User, role data.Role, ses *mgo.Session) *Update {
	return &Update{
		who:        who,
		role:       role,
		reqPayload: payload,
		session:    ses,
	}
}
