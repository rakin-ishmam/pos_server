package sellpayment

import (
	"time"

	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
)

// Create manages create process of the SellPayment
type Create struct {
	Session    *mgo.Session
	ReqPayload CreatePayload
	ResPayload ID
	Who        data.User
	Role       data.Role
	Err        error
}

// Do creates SellPayment
func (c *Create) Do() {
	if err := c.AccessValidate(); err != nil {
		c.Err = err
		return
	}

	dtPay := &data.SellPayment{}
	c.ReqPayload.loadToData(dtPay)

	dtPay.CreatedBy = c.Who.ID
	dtPay.ModifiedBy = c.Who.ID
	dtPay.CreatedAt = time.Now()
	dtPay.ModifiedAt = dtPay.CreatedAt

	if err := c.Validate(dtPay); err != nil {
		c.Err = err
		return
	}

	dbPay := db.SellPayment{Session: c.Session}
	err := dbPay.Put(dtPay)
	if err != nil {
		c.Err = apperr.Database{
			Base:   err,
			Where:  "SellPayment",
			Action: "Create",
		}
		return
	}

	c.ResPayload = ID{ID: string(dtPay.ID)}
}

// Validate takes SellPayment data and returns error
// it checks data validation
func (c *Create) Validate(dtPay *data.SellPayment) error {
	return dtPay.Validate()
}

// AccessValidate returns error. it checks access permission
func (c *Create) AccessValidate() error {
	if !c.Role.PaymentAccess.Can(data.AccessWrite) {
		return apperr.Access{Where: "SellPayment", Permission: string(data.AccessWrite)}
	}

	return nil
}
