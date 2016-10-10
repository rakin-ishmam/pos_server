package orderpayment

import (
	"time"

	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Update manages update process of user
type Update struct {
	Session    *mgo.Session
	ReqPayload UpdatePayload
	ResPayload ID
	Who        data.User
	Role       data.Role
	Err        error
}

// Do update OrderPayment data
func (u *Update) Do() {
	if err := u.AccessValidate(); err != nil {
		u.Err = err
		return
	}

	if err := u.Validate(); err != nil {
		u.Err = err
		return
	}

	dbPay := &db.OrderPayment{Session: u.Session}

	dtPay, err := dbPay.Get(bson.ObjectIdHex(u.ReqPayload.ID.ID))
	if err != nil {
		u.Err = apperr.Database{
			Where:  "OrderPayment",
			Action: "Update",
			Base:   err,
		}
		return
	}

	u.ResPayload.loadToData(dtPay)
	if err := dtPay.Validate(); err != nil {
		u.Err = err
		return
	}

	dtPay.ModifiedBy = u.Who.ID
	dtPay.ModifiedAt = time.Now()

	if err := dbPay.Put(dtPay); err != nil {
		u.Err = apperr.Database{
			Base:   err,
			Where:  "OrderPayment",
			Action: "Update",
		}
		return
	}

	u.ResPayload = ID{ID: string(dtPay.ID)}
}

// AccessValidate returns error
// it checks access permission
func (u *Update) AccessValidate() error {
	if !u.Role.PaymentAccess.Can(data.AccessUpdate) {
		return apperr.Access{Where: "OrderPayment", Permission: string(data.AccessUpdate)}
	}

	return nil
}

// Validate valdes action data
func (u Update) Validate() error {
	if !bson.IsObjectIdHex(u.ReqPayload.ID.ID) {
		return apperr.Validation{
			Where: "OrderPayment",
			Field: "id",
			Cause: apperr.ValidationInvalid,
		}
	}

	return nil
}
