package sellpayment

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/rakin-ishmam/pos_server/data"
)

// ID stroe only id field of the SellPayment
type ID struct {
	ID string `json:"id, omitempty"`
}

func (i *ID) loadFromData(dtPay *data.SellPayment) {
	i.ID = string(dtPay.ID)
}

func (i *ID) loadToData(dtPay *data.SellPayment) {
	if bson.IsObjectIdHex(i.ID) {
		dtPay.ID = bson.ObjectIdHex(i.ID)
	}
}

// CreatePayload stores data for SellPayment create
type CreatePayload struct {
	SellID      string  `json:"sell_id"`
	Amount      float64 `json:"amount"`
	PaymentType string  `json:"payment_type"`
	Number      string  `json:"number"`
	Comment     string  `json:"comment"`
}

func (c *CreatePayload) loadFromData(dtPay *data.SellPayment) {
	c.SellID = string(dtPay.SellID)
	c.Amount = dtPay.Amount
	c.PaymentType = string(dtPay.PaymentType)
	c.Number = dtPay.Number
	c.Comment = dtPay.Comment
}

func (c *CreatePayload) loadToData(dtPay *data.SellPayment) {
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
	ID
	CreatePayload
}

func (u *UpdatePayload) loadFromData(dtPay *data.SellPayment) {
	u.ID.loadFromData(dtPay)
	u.CreatePayload.loadFromData(dtPay)
}

func (u *UpdatePayload) loadToData(dtPay *data.SellPayment) {
	u.ID.loadToData(dtPay)
	u.CreatePayload.loadToData(dtPay)
}

// DetailPayload stores detail payload for SellPayment
type DetailPayload struct {
	UpdatePayload

	Deleted    bool      `json:"deleted"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

func (d *DetailPayload) loadFromData(dtPay *data.SellPayment) {
	d.UpdatePayload.loadFromData(dtPay)

	d.Deleted = dtPay.Deleted
	d.CreatedAt = dtPay.CreatedAt
	d.CreatedBy = string(dtPay.CreatedBy)
	d.ModifiedAt = dtPay.ModifiedAt
	d.ModifiedBy = string(dtPay.ModifiedBy)
}

func (d *DetailPayload) loadToData(dtPay *data.SellPayment) {
	d.UpdatePayload.loadToData(dtPay)

	dtPay.Deleted = d.Deleted
	dtPay.CreatedAt = d.CreatedAt
	if bson.IsObjectIdHex(d.CreatedBy) {
		dtPay.CreatedBy = bson.ObjectIdHex(d.CreatedBy)
	}
	dtPay.ModifiedAt = d.ModifiedAt
	if bson.IsObjectIdHex(d.ModifiedBy) {
		dtPay.ModifiedBy = bson.ObjectIdHex(d.ModifiedBy)
	}
}
