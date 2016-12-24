package sellpayment

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
	Session    *mgo.Session
	ReqPayload UpdatePayload
	ResPayload geninfo.ID
	Who        data.User
	Role       data.Role
	Err        error
}

// Do update SellPayment data
func (u *Update) Do() {
	if err := u.AccessValidate(); err != nil {
		u.Err = err
		return
	}

	if err := u.Validate(); err != nil {
		u.Err = err
		return
	}

	dbPay := &db.SellPayment{Session: u.Session}

	dtPay, err := dbPay.Get(bson.ObjectIdHex(u.ReqPayload.ID.ID))
	if err != nil {
		u.Err = apperr.Database{
			Where:  "SellPayment",
			Action: "Update",
			Base:   err,
		}
		return
	}

	u.ResPayload.LoadToData(&dtPay.Track)
	if err := dtPay.Validate(); err != nil {
		u.Err = err
		return
	}

	dtPay.ModifiedBy = u.Who.ID

	if err := dbPay.Put(dtPay); err != nil {
		u.Err = apperr.Database{
			Base:   err,
			Where:  "SellPayment",
			Action: "Update",
		}
		return
	}

	u.ResPayload = geninfo.ID{ID: string(dtPay.ID)}
}

// Result returns result of thte action
func (u Update) Result() (io.Reader, error) {
	if u.Err != nil {
		return nil, u.Err
	}

	return converter.JSONtoBuff(u.ResPayload)
}

// AccessValidate returns error
// it checks access permission
func (u *Update) AccessValidate() error {
	if !u.Role.PaymentAccess.Can(data.AccessUpdate) {
		return apperr.Access{Where: "SellPayment", Permission: string(data.AccessUpdate)}
	}

	return nil
}

// Validate valdes action data
func (u Update) Validate() error {
	if !bson.IsObjectIdHex(u.ReqPayload.ID.ID) {
		return apperr.Validation{
			Where: "SellPayment",
			Field: "id",
			Cause: apperr.ValidationInvalid,
		}
	}

	return nil
}
