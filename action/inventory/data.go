package inventory

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/rakin-ishmam/pos_server/data"
)

// ID stroe only id field of the Inventory
type ID struct {
	ID string `json:"id, omitempty"`
}

func (i *ID) loadFromData(dtInvt *data.Inventory) {
	i.ID = string(dtInvt.ID)
}

func (i *ID) loadToData(dtInvt *data.Inventory) {
	if bson.IsObjectIdHex(i.ID) {
		dtInvt.ID = bson.ObjectIdHex(i.ID)
	}
}

// CreatePayload stores data for Inventory create
type CreatePayload struct {
	Code      string  `json:"code"`
	ProductID string  `json:"product_id"`
	SalePrice float64 `json:"sale_price"`
	BuyPrice  float64 `json:"buy_price"`
	Quantity  int     `json:"quantity"`
}

func (c *CreatePayload) loadFromData(dtInvt *data.Inventory) {
	c.Code = dtInvt.Code
	c.ProductID = string(dtInvt.ProductID)
	c.SalePrice = dtInvt.SalePrice
	c.BuyPrice = dtInvt.BuyPrice
	c.Quantity = dtInvt.Quantity
}

func (c *CreatePayload) loadToData(dtInvt *data.Inventory) {
	dtInvt.Code = c.Code
	if bson.IsObjectIdHex(c.ProductID) {
		dtInvt.ProductID = bson.ObjectIdHex(c.ProductID)
	}
	dtInvt.SalePrice = c.SalePrice
	dtInvt.BuyPrice = c.BuyPrice
	dtInvt.Quantity = c.Quantity
}

// UpdatePayload stores update payload for Inventory
type UpdatePayload struct {
	ID
	CreatePayload
}

func (u *UpdatePayload) loadFromData(dtInvt *data.Inventory) {
	u.ID.loadFromData(dtInvt)
	u.CreatePayload.loadFromData(dtInvt)
}

func (u *UpdatePayload) loadToData(dtInvt *data.Inventory) {
	u.ID.loadToData(dtInvt)
	u.CreatePayload.loadToData(dtInvt)
}

// DetailPayload stores detail payload for Inventory
type DetailPayload struct {
	UpdatePayload

	Deleted    bool      `json:"deleted"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

func (d *DetailPayload) loadFromData(dtInvt *data.Inventory) {
	d.UpdatePayload.loadFromData(dtInvt)

	d.Deleted = dtInvt.Deleted
	d.CreatedAt = dtInvt.CreatedAt
	d.CreatedBy = string(dtInvt.CreatedBy)
	d.ModifiedAt = dtInvt.ModifiedAt
	d.ModifiedBy = string(dtInvt.ModifiedBy)
}

func (d *DetailPayload) loadToData(dtInvt *data.Inventory) {
	d.UpdatePayload.loadToData(dtInvt)

	dtInvt.Deleted = d.Deleted
	dtInvt.CreatedAt = d.CreatedAt
	if bson.IsObjectIdHex(d.CreatedBy) {
		dtInvt.CreatedBy = bson.ObjectIdHex(d.CreatedBy)
	}
	dtInvt.ModifiedAt = d.ModifiedAt
	if bson.IsObjectIdHex(d.ModifiedBy) {
		dtInvt.ModifiedBy = bson.ObjectIdHex(d.ModifiedBy)
	}
}
