package orderpayment

import (
	"time"

	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
)

// Create manages create process of the OrderPayment
type Create struct {
	Session    *mgo.Session
	ReqPayload CreatePayload
	ResPayload ID
	Who        data.User
	Role       data.Role
	Err        error
}

// Do creates OrderPayment
func (c *Create) Do() {
	if err := c.AccessValidate(); err != nil {
		c.Err = err
		return
	}

	dtPay := &data.OrderPayment{}
	c.ReqPayload.loadToData(dtPay)

	dtPay.CreatedBy = c.Who.ID
	dtPay.ModifiedBy = c.Who.ID
	dtPay.CreatedAt = time.Now()
	dtPay.ModifiedAt = dtPay.CreatedAt

	if err := dtPay.Validate(); err != nil {
		c.Err = err
		return
	}

	dbPay := db.OrderPayment{Session: c.Session}
	err := dbPay.Put(dtPay)
	if err != nil {
		c.Err = apperr.Database{
			Base:   err,
			Where:  "OrderPayment",
			Action: "Create",
		}
		return
	}

	c.ResPayload = ID{ID: string(dtPay.ID)}
}

// AccessValidate returns error. it checks access permission
func (c *Create) AccessValidate() error {
	if !c.Role.PaymentAccess.Can(data.AccessWrite) {
		return apperr.Access{Where: "OrderPayment", Permission: string(data.AccessWrite)}
	}

	return nil
}
