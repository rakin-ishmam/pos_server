package product

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/rakin-ishmam/pos_server/data"
)

// ID stroe only id field of the Product
type ID struct {
	ID string `json:"id, omitempty"`
}

func (i *ID) loadFromData(dtProd *data.Product) {
	i.ID = string(dtProd.ID)
}

func (i *ID) loadToData(dtProd *data.Product) {
	if bson.IsObjectIdHex(i.ID) {
		dtProd.ID = bson.ObjectIdHex(i.ID)
	}
}

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

func (c *CreatePayload) loadFromData(dtProd *data.Product) {
	c.Name = dtProd.Name
	c.Code = dtProd.Code
	c.AvtFileID = dtProd.AvtFileID
	c.CategoryID = string(dtProd.CategoryID)
	c.SalePrice = dtProd.SalePrice
	c.BuyPrice = dtProd.BuyPrice
	c.Quantity = dtProd.Quantity
}

func (c *CreatePayload) loadToData(dtProd *data.Product) {
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
	ID
	CreatePayload
}

func (u *UpdatePayload) loadFromData(dtProd *data.Product) {
	u.ID.loadFromData(dtProd)
	u.CreatePayload.loadFromData(dtProd)
}

func (u *UpdatePayload) loadToData(dtProd *data.Product) {
	u.ID.loadToData(dtProd)
	u.CreatePayload.loadToData(dtProd)
}

// DetailPayload stores detail payload for Product
type DetailPayload struct {
	UpdatePayload

	Deleted    bool      `json:"deleted"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

func (d *DetailPayload) loadFromData(dtProd *data.Product) {
	d.UpdatePayload.loadFromData(dtProd)

	d.Deleted = dtProd.Deleted
	d.CreatedAt = dtProd.CreatedAt
	d.CreatedBy = string(dtProd.CreatedBy)
	d.ModifiedAt = dtProd.ModifiedAt
	d.ModifiedBy = string(dtProd.ModifiedBy)
}

func (d *DetailPayload) loadToData(dtProd *data.Product) {
	d.UpdatePayload.loadToData(dtProd)

	dtProd.Deleted = d.Deleted
	dtProd.CreatedAt = d.CreatedAt
	dtProd.CreatedAt = d.CreatedAt
	if bson.IsObjectIdHex(d.CreatedBy) {
		dtProd.CreatedBy = bson.ObjectIdHex(d.CreatedBy)
	}
	dtProd.ModifiedAt = d.ModifiedAt
	if bson.IsObjectIdHex(d.ModifiedBy) {
		dtProd.ModifiedBy = bson.ObjectIdHex(d.ModifiedBy)
	}
}
