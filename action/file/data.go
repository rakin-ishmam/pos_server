package file

import (
	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/data"
)

// CreatePayload stores data for File create
type CreatePayload struct {
	Location string `gorethink:"location"`
}

// LoadFromData copy data
func (c *CreatePayload) LoadFromData(dtFile *data.File) {
	c.Location = dtFile.Location
}

// LoadToData copy to data
func (c *CreatePayload) LoadToData(dtFile *data.File) {
	dtFile.Location = c.Location
}

// UpdatePayload stores update payload for File
type UpdatePayload struct {
	geninfo.ID
	CreatePayload
}

// LoadFromData copy data
func (u *UpdatePayload) LoadFromData(dtFile *data.File) {
	u.ID.LoadFromData(&dtFile.Track)
	u.CreatePayload.LoadFromData(dtFile)
}

// LoadToData copy to data
func (u *UpdatePayload) LoadToData(dtFile *data.File) {
	u.ID.LoadToData(&dtFile.Track)
	u.CreatePayload.LoadToData(dtFile)
}

// DetailPayload stores detail payload for Category
type DetailPayload struct {
	UpdatePayload
	geninfo.General
}

// LoadFromData copy data
func (d *DetailPayload) LoadFromData(dtFile *data.File) {
	d.UpdatePayload.LoadFromData(dtFile)

	d.General.LoadFromData(&dtFile.Track)
}

// LoadToData copy to data
func (d *DetailPayload) LoadToData(dtFile *data.File) {
	d.UpdatePayload.LoadToData(dtFile)

	d.General.LoadToData(&dtFile.Track)
}
