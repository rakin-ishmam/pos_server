package orderpayment

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/data"
)

// CreatePayload stores data for OrderPayment create
type CreatePayload struct {
	OrderID     string  `json:"order_id"`
	Amount      float64 `json:"amount"`
	PaymentType string  `json:"payment_type"`
	Number      string  `json:"number"`
	Comment     string  `json:"comment"`
}

// LoadFromData copy data
func (c *CreatePayload) LoadFromData(dtPay *data.OrderPayment) {
	c.OrderID = string(dtPay.OrderID)
	c.Amount = dtPay.Amount
	c.PaymentType = string(dtPay.PaymentType)
	c.Number = dtPay.Number
	c.Comment = dtPay.Comment
}

// LoadToData copy to data
func (c *CreatePayload) LoadToData(dtPay *data.OrderPayment) {
	if bson.IsObjectIdHex(c.OrderID) {
		dtPay.OrderID = bson.ObjectIdHex(c.OrderID)
	}
	dtPay.Amount = c.Amount
	dtPay.PaymentType = data.PaymentType(c.PaymentType)
	dtPay.Number = c.Number
	dtPay.Comment = c.Comment
}

// UpdatePayload stores update payload for OrderPayment
type UpdatePayload struct {
	geninfo.ID
	CreatePayload
}

// LoadFromData copy data
func (u *UpdatePayload) LoadFromData(dtPay *data.OrderPayment) {
	u.ID.LoadFromData(&dtPay.Track)
	u.CreatePayload.LoadFromData(dtPay)
}

// LoadToData copy to data
func (u *UpdatePayload) LoadToData(dtPay *data.OrderPayment) {
	u.ID.LoadToData(&dtPay.Track)
	u.CreatePayload.LoadToData(dtPay)
}

// DetailPayload stores detail payload for OrderPayment
type DetailPayload struct {
	UpdatePayload

	geninfo.General
}

// LoadFromData copy data
func (d *DetailPayload) LoadFromData(dtPay *data.OrderPayment) {
	d.UpdatePayload.LoadFromData(dtPay)

	d.General.LoadFromData(&dtPay.Track)
}

// LoadToData copy to data
func (d *DetailPayload) LoadToData(dtPay *data.OrderPayment) {
	d.UpdatePayload.LoadToData(dtPay)

	d.General.LoadToData(&dtPay.Track)
}
