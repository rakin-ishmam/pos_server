package category

import (
	"time"

	"github.com/rakin-ishmam/pos_server/data"
	"gopkg.in/mgo.v2/bson"
)

// ID stroe only id field of the Category
type ID struct {
	ID string `json:"id, omitempty"`
}

func (i *ID) loadFromData(dtCat *data.Category) {
	i.ID = string(dtCat.ID)
}

func (i *ID) loadToData(dtCat *data.Category) {
	if bson.IsObjectIdHex(i.ID) {
		dtCat.ID = bson.ObjectIdHex(i.ID)
	}
}

// CreatePayload stores data for Category create
type CreatePayload struct {
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
}

func (c *CreatePayload) loadFromData(dtCat *data.Category) {
	c.Name = dtCat.Name
	c.CategoryID = string(dtCat.CategoryID)
}

func (c *CreatePayload) loadToData(dtCat *data.Category) {
	dtCat.Name = c.Name
	if bson.IsObjectIdHex(c.CategoryID) {
		dtCat.CategoryID = bson.ObjectIdHex(c.CategoryID)
	}
}

// UpdatePayload stores update payload from Category
type UpdatePayload struct {
	ID
	CreatePayload
}

func (u *UpdatePayload) loadFromData(dtCat *data.Category) {
	u.ID.loadFromData(dtCat)
	u.CreatePayload.loadFromData(dtCat)
}

func (u *UpdatePayload) loadToData(dtCat *data.Category) {
	u.ID.loadToData(dtCat)
	u.CreatePayload.loadToData(dtCat)
}

// DetailPayload stores detail payload from Category
type DetailPayload struct {
	UpdatePayload

	Deleted    bool      `json:"deleted"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

func (d *DetailPayload) loadFromData(dtCat *data.Category) {
	d.UpdatePayload.loadFromData(dtCat)

	d.Deleted = dtCat.Deleted
	d.CreatedAt = dtCat.CreatedAt
	d.CreatedBy = string(dtCat.CreatedBy)
	d.ModifiedAt = dtCat.ModifiedAt
	d.ModifiedBy = string(dtCat.ModifiedBy)
}

func (d *DetailPayload) loadToData(dtCat *data.Category) {
	d.UpdatePayload.loadToData(dtCat)

	dtCat.Deleted = d.Deleted
	dtCat.CreatedAt = d.CreatedAt
	if bson.IsObjectIdHex(d.CreatedBy) {
		dtCat.CreatedBy = bson.ObjectIdHex(d.CreatedBy)
	}
	dtCat.ModifiedAt = d.ModifiedAt
	if bson.IsObjectIdHex(d.ModifiedBy) {
		dtCat.ModifiedBy = bson.ObjectIdHex(d.ModifiedBy)
	}
}
