package product

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/data"
)

// CreatePayload stores data for Product create
type CreatePayload struct {
	Name       string  `json:"name"`
	Code       string  `json:"code"`
	AvtFileID  string  `json:"avt_file_id"`
	CategoryID string  `json:"category_id"`
	SalePrice  float64 `json:"sale_price"`
	BuyPrice   float64 `json:"buy_price"`
	Quantity   int     `json:"quantity"`
}

// LoadFromData copy data
func (c *CreatePayload) LoadFromData(dtProd *data.Product) {
	c.Name = dtProd.Name
	c.Code = dtProd.Code
	c.AvtFileID = dtProd.AvtFileID
	c.CategoryID = string(dtProd.CategoryID)
	c.SalePrice = dtProd.SalePrice
	c.BuyPrice = dtProd.BuyPrice
	c.Quantity = dtProd.Quantity
}

// LoadToData copy to data
func (c *CreatePayload) LoadToData(dtProd *data.Product) {
	dtProd.Name = c.Name
	dtProd.Code = c.Code
	dtProd.AvtFileID = c.AvtFileID
	if bson.IsObjectIdHex(c.CategoryID) {
		dtProd.CategoryID = bson.ObjectIdHex(c.CategoryID)
	}
	dtProd.SalePrice = c.SalePrice
	dtProd.BuyPrice = c.BuyPrice
	dtProd.Quantity = c.Quantity
}

// UpdatePayload stores update payload for Product
type UpdatePayload struct {
	geninfo.ID
	CreatePayload
}

// LoadFromData copy data
func (u *UpdatePayload) LoadFromData(dtProd *data.Product) {
	u.ID.LoadFromData(&dtProd.Track)
	u.CreatePayload.LoadFromData(dtProd)
}

// LoadToData copy to data
func (u *UpdatePayload) LoadToData(dtProd *data.Product) {
	u.ID.LoadToData(&dtProd.Track)
	u.CreatePayload.LoadToData(dtProd)
}

// DetailPayload stores detail payload for Product
type DetailPayload struct {
	UpdatePayload

	geninfo.General
}

// LoadFromData copy data
func (d *DetailPayload) LoadFromData(dtProd *data.Product) {
	d.UpdatePayload.LoadFromData(dtProd)

	d.General.LoadFromData(&dtProd.Track)
}

// LoadToData copy to data
func (d *DetailPayload) LoadToData(dtProd *data.Product) {
	d.UpdatePayload.LoadToData(dtProd)

	d.General.LoadToData(&dtProd.Track)
}
