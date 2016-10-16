package orderpayment

import (
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Fetch store fetch request data and store data
type Fetch struct {
	Session    *mgo.Session
	ReqID      string
	ResPayload DetailPayload
	Who        data.User
	Role       data.Role
	Err        error
}

// Do takes necessary steps to fetch OrderPayment data
func (f *Fetch) Do() {
	if err := f.AccessValidate(); err != nil {
		f.Err = err
		return
	}

	if err := f.Validate(); err != nil {
		f.Err = err
		return
	}

	dbPay := db.OrderPayment{Session: f.Session}
	dtPay, err := dbPay.Get(bson.ObjectIdHex(f.ReqID))
	if err != nil {
		f.Err = apperr.Database{
			Base:   err,
			Where:  "OrderPayment",
			Action: "Fetch",
		}
		return
	}

	f.ResPayload = DetailPayload{}
	f.ResPayload.LoadFromData(dtPay)
}

// ActionErr returns error of the action
func (f Fetch) ActionErr() error {
	return f.Err
}

// Result returns result of thte action
func (f Fetch) Result() interface{} {
	return &f.ResPayload
}

// AccessValidate checks access permission
func (f *Fetch) AccessValidate() error {
	if !f.Role.PaymentAccess.Can(data.AccessRead) {
		return apperr.Access{Where: "OrderPayment", Permission: string(data.AccessRead)}
	}

	return nil
}

// Validate valids action data
func (f Fetch) Validate() error {
	if !bson.IsObjectIdHex(f.ReqID) {
		return apperr.Validation{
			Where: "OrderPayment",
			Field: "id",
			Cause: apperr.ValidationInvalid,
		}
	}
	return nil
}
