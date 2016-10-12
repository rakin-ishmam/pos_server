package inventory

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/data"
)

// CreatePayload stores data for Inventory create
type CreatePayload struct {
	Code      string  `json:"code"`
	ProductID string  `json:"product_id"`
	SalePrice float64 `json:"sale_price"`
	BuyPrice  float64 `json:"buy_price"`
	Quantity  int     `json:"quantity"`
}

// LoadFromData copy data
func (c *CreatePayload) LoadFromData(dtInvt *data.Inventory) {
	c.Code = dtInvt.Code
	c.ProductID = string(dtInvt.ProductID)
	c.SalePrice = dtInvt.SalePrice
	c.BuyPrice = dtInvt.BuyPrice
	c.Quantity = dtInvt.Quantity
}

// LoadToData copy to data
func (c *CreatePayload) LoadToData(dtInvt *data.Inventory) {
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
	geninfo.ID
	CreatePayload
}

// LoadFromData copy data
func (u *UpdatePayload) LoadFromData(dtInvt *data.Inventory) {
	u.ID.LoadFromData(&dtInvt.Track)
	u.CreatePayload.LoadFromData(dtInvt)
}

// LoadToData copy to data
func (u *UpdatePayload) LoadToData(dtInvt *data.Inventory) {
	u.ID.LoadToData(&dtInvt.Track)
	u.CreatePayload.LoadToData(dtInvt)
}

// DetailPayload stores detail payload for Inventory
type DetailPayload struct {
	UpdatePayload

	geninfo.General
}

// LoadFromData copy data
func (d *DetailPayload) LoadFromData(dtInvt *data.Inventory) {
	d.UpdatePayload.LoadFromData(dtInvt)

	d.General.LoadFromData(&dtInvt.Track)
}

// LoadToData copy to data
func (d *DetailPayload) LoadToData(dtInvt *data.Inventory) {
	d.UpdatePayload.LoadToData(dtInvt)

	d.General.LoadToData(&dtInvt.Track)
}
