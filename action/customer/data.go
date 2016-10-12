package customer

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/data"
)

// CreatePayload stores create data of the Customer
type CreatePayload struct {
	Name      string   `json:"name"`
	AvtFileID string   `json:"avt_file_id"`
	Code      string   `json:"code"`
	Email     string   `json:"email"`
	Gender    string   `json:"gender"`
	Address   string   `json:"address"`
	Phone     []string `json:"phone"`
}

// LoadFromData copy data
func (c *CreatePayload) LoadFromData(dtCus *data.Customer) {
	c.Name = dtCus.Name
	c.AvtFileID = string(dtCus.AvtFileID)
	c.Code = dtCus.Code
	c.Email = dtCus.Email
	c.Gender = string(dtCus.Gender)
	c.Address = dtCus.Address
	c.Phone = dtCus.Phone
}

// LoadToData copy to data
func (c *CreatePayload) LoadToData(dtCus *data.Customer) {
	dtCus.Name = c.Name
	if bson.IsObjectIdHex(c.AvtFileID) {
		dtCus.AvtFileID = bson.ObjectIdHex(c.AvtFileID)
	}
	dtCus.Code = c.Code
	dtCus.Email = c.Email
	dtCus.Gender = data.Gender(c.Gender)
	dtCus.Address = c.Address
	dtCus.Phone = c.Phone
}

// UpdatePayload stores update payload of the Customer
type UpdatePayload struct {
	geninfo.ID
	CreatePayload
}

// LoadFromData copy data
func (u *UpdatePayload) LoadFromData(dtCus *data.Customer) {
	u.ID.LoadFromData(&dtCus.Track)
	u.CreatePayload.LoadFromData(dtCus)
}

// LoadToData copy to data
func (u *UpdatePayload) LoadToData(dtCus *data.Customer) {
	u.ID.LoadToData(&dtCus.Track)
	u.CreatePayload.LoadToData(dtCus)
}

// DetailPayload stores detail payload of the Customer
type DetailPayload struct {
	UpdatePayload

	geninfo.General
}

// LoadFromData copy data
func (d *DetailPayload) LoadFromData(dtCus *data.Customer) {
	d.UpdatePayload.LoadFromData(dtCus)

	d.General.LoadFromData(&dtCus.Track)
}

// LoadToData copy to data
func (d *DetailPayload) LoadToData(dtCus *data.Customer) {
	d.UpdatePayload.LoadToData(dtCus)

	d.General.LoadToData(&dtCus.Track)
}
