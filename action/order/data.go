package order

import (
	"time"

	"github.com/rakin-ishmam/pos_server/data"
)

// ID stroe only id field of the Inventory
type ID struct {
	ID string `json:"id, omitempty"`
}

func (i *ID) loadFromData(dtInvt *data.Inventory) {
	i.ID = dtInvt.ID
}

func (i *ID) loadToData(dtInvt *data.Inventory) {
	dtInvt.ID = i.ID
}

type ProductCreate struct {
	Product  string `gorethink:"product"`
	Quantity int    `gorethink:"quantity"`
}

type ProductDetail struct {
	Product  product.DetailPayload `gorethink:"product"`
	Quantity int                   `gorethink:"quantity"`
}

// CreatePayload stores data for Inventory create
type CreatePayload struct {
	Code         string         `gorethink:"code"`
	OrderProduct []OrderProduct `gorethink:"order_product"`
	TotalPrice   float64        `gorethink:"total_price"`
	TotalPaid    float64        `gorethink:"total_paid"`
	Discount     float64        `gorethink:"discount"`
	Delivered    bool           `gorethink:"delivered"`
}

func (c *CreatePayload) loadFromData(dtInvt *data.Inventory) {
	c.ProductID = dtInvt.ProductID
	c.SalePrice = dtInvt.SalePrice
	c.BuyPrice = dtInvt.BuyPrice
	c.Quantity = dtInvt.Quantity
}

func (c *CreatePayload) loadToData(dtInvt *data.Inventory) {
	dtInvt.ProductID = c.ProductID
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
	d.CreatedBy = dtInvt.CreatedBy
	d.ModifiedAt = dtInvt.ModifiedAt
	d.ModifiedBy = dtInvt.ModifiedBy
}

func (d *DetailPayload) loadToData(dtInvt *data.Inventory) {
	d.UpdatePayload.loadToData(dtInvt)

	dtInvt.Deleted = d.Deleted
	dtInvt.CreatedAt = d.CreatedAt
	dtInvt.CreatedBy = d.CreatedBy
	dtInvt.ModifiedAt = d.ModifiedAt
	dtInvt.ModifiedBy = d.ModifiedBy
}
