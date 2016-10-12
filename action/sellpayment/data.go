package sellpayment

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/data"
)

// CreatePayload stores data for SellPayment create
type CreatePayload struct {
	SellID      string  `json:"sell_id"`
	Amount      float64 `json:"amount"`
	PaymentType string  `json:"payment_type"`
	Number      string  `json:"number"`
	Comment     string  `json:"comment"`
}

// LoadFromData copy data
func (c *CreatePayload) LoadFromData(dtPay *data.SellPayment) {
	c.SellID = string(dtPay.SellID)
	c.Amount = dtPay.Amount
	c.PaymentType = string(dtPay.PaymentType)
	c.Number = dtPay.Number
	c.Comment = dtPay.Comment
}

// LoadToData copty to data
func (c *CreatePayload) LoadToData(dtPay *data.SellPayment) {
	if bson.IsObjectIdHex(c.SellID) {
		dtPay.SellID = bson.ObjectIdHex(c.SellID)
	}
	dtPay.Amount = c.Amount
	dtPay.PaymentType = data.PaymentType(c.PaymentType)
	dtPay.Number = c.Number
	dtPay.Comment = c.Comment
}

// UpdatePayload stores update payload for SellPayment
type UpdatePayload struct {
	geninfo.ID
	CreatePayload
}

// LoadFromData copy data
func (u *UpdatePayload) LoadFromData(dtPay *data.SellPayment) {
	u.ID.LoadFromData(&dtPay.Track)
	u.CreatePayload.LoadFromData(dtPay)
}

// LoadToData copty to data
func (u *UpdatePayload) LoadToData(dtPay *data.SellPayment) {
	u.ID.LoadToData(&dtPay.Track)
	u.CreatePayload.LoadToData(dtPay)
}

// DetailPayload stores detail payload for SellPayment
type DetailPayload struct {
	UpdatePayload

	geninfo.General
}

// LoadFromData copy data
func (d *DetailPayload) LoadFromData(dtPay *data.SellPayment) {
	d.UpdatePayload.LoadFromData(dtPay)

	d.General.LoadFromData(&dtPay.Track)
}

// LoadToData copty to data
func (d *DetailPayload) LoadToData(dtPay *data.SellPayment) {
	d.UpdatePayload.LoadToData(dtPay)

	d.General.LoadToData(&dtPay.Track)
}
