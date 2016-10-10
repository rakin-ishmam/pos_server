package sellpayment

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

// Do takes necessary steps to fetch SellPayment data
func (f *Fetch) Do() {
	if err := f.AccessValidate(); err != nil {
		f.Err = err
		return
	}

	if err := f.Validate(); err != nil {
		f.Err = err
		return
	}

	dbPay := db.SellPayment{Session: f.Session}
	dtPay, err := dbPay.Get(bson.ObjectIdHex(f.ReqID))
	if err != nil {
		f.Err = apperr.Database{
			Base:   err,
			Where:  "SellPayment",
			Action: "Fetch",
		}
		return
	}

	f.ResPayload = DetailPayload{}
	f.ResPayload.loadFromData(dtPay)
}

// AccessValidate checks access permission
func (f *Fetch) AccessValidate() error {
	if !f.Role.PaymentAccess.Can(data.AccessRead) {
		return apperr.Access{Where: "SellPayment", Permission: string(data.AccessRead)}
	}

	return nil
}

// Validate valids action data
func (f Fetch) Validate() error {
	if !bson.IsObjectIdHex(f.ReqID) {
		return apperr.Validation{
			Where: "SellPayment",
			Field: "id",
			Cause: apperr.ValidationInvalid,
		}
	}
	return nil
}
