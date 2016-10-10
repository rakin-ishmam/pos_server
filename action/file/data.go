package file

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/rakin-ishmam/pos_server/data"
)

// ID stroe only id field of the Category
type ID struct {
	ID string `json:"id, omitempty"`
}

func (i *ID) loadFromData(dtFile *data.File) {
	i.ID = string(dtFile.ID)
}

func (i *ID) loadToData(dtFile *data.File) {
	if bson.IsObjectIdHex(i.ID) {
		dtFile.ID = bson.ObjectIdHex(i.ID)
	}
}

// CreatePayload stores data for File create
type CreatePayload struct {
	Location string `gorethink:"location"`
}

func (c *CreatePayload) loadFromData(dtFile *data.File) {
	c.Location = dtFile.Location
}

func (c *CreatePayload) loadToData(dtFile *data.File) {
	dtFile.Location = c.Location
}

// UpdatePayload stores update payload for File
type UpdatePayload struct {
	ID
	CreatePayload
}

func (u *UpdatePayload) loadFromData(dtFile *data.File) {
	u.ID.loadFromData(dtFile)
	u.CreatePayload.loadFromData(dtFile)
}

func (u *UpdatePayload) loadToData(dtFile *data.File) {
	u.ID.loadToData(dtFile)
	u.CreatePayload.loadToData(dtFile)
}

// DetailPayload stores detail payload for Category
type DetailPayload struct {
	UpdatePayload

	Deleted    bool      `json:"deleted"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

func (d *DetailPayload) loadFromData(dtFile *data.File) {
	d.UpdatePayload.loadFromData(dtFile)

	d.Deleted = dtFile.Deleted
	d.CreatedAt = dtFile.CreatedAt
	d.CreatedBy = string(dtFile.CreatedBy)
	d.ModifiedAt = dtFile.ModifiedAt
	d.ModifiedBy = string(dtFile.ModifiedBy)
}

func (d *DetailPayload) loadToData(dtFile *data.File) {
	d.UpdatePayload.loadToData(dtFile)

	dtFile.Deleted = d.Deleted
	dtFile.CreatedAt = d.CreatedAt
	if bson.IsObjectIdHex(d.CreatedBy) {
		dtFile.CreatedBy = bson.ObjectIdHex(d.CreatedBy)
	}
	dtFile.ModifiedAt = d.ModifiedAt
	if bson.IsObjectIdHex(d.ModifiedBy) {
		dtFile.ModifiedBy = bson.ObjectIdHex(d.ModifiedBy)
	}
}
