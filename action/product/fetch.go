package product

import (
	"io"

	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/converter"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Fetch store fetch request data and store data
type Fetch struct {
	session    *mgo.Session
	reqID      string
	resPayload DetailPayload
	who        data.User
	role       data.Role
	err        error
}

// Do takes necessary steps to fetch Product data
func (f *Fetch) Do() {
	if err := f.AccessValidate(); err != nil {
		f.err = err
		return
	}

	if err := f.Validate(); err != nil {
		f.err = err
		return
	}

	dbProd := db.Product{Session: f.session}
	dtProd, err := dbProd.Get(bson.ObjectIdHex(f.reqID))
	if err != nil {
		f.err = apperr.Database{
			Base:   err,
			Where:  "Prod",
			Action: "Fetch",
		}
		return
	}

	f.resPayload = DetailPayload{}
	f.resPayload.LoadFromData(dtProd)
}

// Result returns result of thte action
func (f Fetch) Result() (io.Reader, error) {
	if f.err != nil {
		return nil, f.err
	}

	return converter.JSONtoBuff(f.resPayload)
}

// AccessValidate checks access permission
func (f *Fetch) AccessValidate() error {
	if !f.role.ProductAccess.Can(data.AccessRead) {
		return apperr.Access{Where: "Product", Permission: string(data.AccessRead)}
	}

	return nil
}

// Validate valids action data
func (f Fetch) Validate() error {
	if !bson.IsObjectIdHex(f.reqID) {
		return apperr.Validation{
			Where: "Product",
			Field: "id",
			Cause: apperr.StrInvalid,
		}
	}
	return nil
}

// NewFetch returns action to fetch role
func NewFetch(id string, who data.User, role data.Role, ses *mgo.Session) *Fetch {
	return &Fetch{
		who:     who,
		role:    role,
		reqID:   id,
		session: ses,
	}
}
