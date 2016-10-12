package category

import (
	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/data"
	"gopkg.in/mgo.v2/bson"
)

// CreatePayload stores data for Category create
type CreatePayload struct {
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
}

// LoadFromData copy data
func (c *CreatePayload) LoadFromData(dtCat *data.Category) {
	c.Name = dtCat.Name
	c.CategoryID = string(dtCat.CategoryID)
}

// LoadToData copy to data
func (c *CreatePayload) LoadToData(dtCat *data.Category) {
	dtCat.Name = c.Name
	if bson.IsObjectIdHex(c.CategoryID) {
		dtCat.CategoryID = bson.ObjectIdHex(c.CategoryID)
	}
}

// UpdatePayload stores update payload from Category
type UpdatePayload struct {
	geninfo.ID
	CreatePayload
}

// LoadFromData copy data
func (u *UpdatePayload) LoadFromData(dtCat *data.Category) {
	u.ID.LoadFromData(&dtCat.Track)
	u.CreatePayload.LoadFromData(dtCat)
}

// LoadToData copy to data
func (u *UpdatePayload) LoadToData(dtCat *data.Category) {
	u.ID.LoadToData(&dtCat.Track)
	u.CreatePayload.LoadToData(dtCat)
}

// DetailPayload stores detail payload from Category
type DetailPayload struct {
	UpdatePayload

	geninfo.General
}

// LoadFromData copy data
func (d *DetailPayload) LoadFromData(dtCat *data.Category) {
	d.UpdatePayload.LoadFromData(dtCat)

	d.General.LoadFromData(&dtCat.Track)
}

// LoadToData copy to data
func (d *DetailPayload) LoadToData(dtCat *data.Category) {
	d.UpdatePayload.LoadToData(dtCat)

	d.General.LoadToData(&dtCat.Track)
}
