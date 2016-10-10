package customer

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/rakin-ishmam/pos_server/data"
)

// ID stroe only id field of the Customer
type ID struct {
	ID string `json:"id, omitempty"`
}

func (i *ID) loadFromData(dtCus *data.Customer) {
	i.ID = string(dtCus.ID)
}

func (i *ID) loadToData(dtCus *data.Customer) {
	if bson.IsObjectIdHex(i.ID) {
		dtCus.ID = bson.ObjectIdHex(i.ID)
	}
}

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

func (c *CreatePayload) loadFromData(dtCus *data.Customer) {
	c.Name = dtCus.Name
	c.AvtFileID = string(dtCus.AvtFileID)
	c.Code = dtCus.Code
	c.Email = dtCus.Email
	c.Gender = string(dtCus.Gender)
	c.Address = dtCus.Address
	c.Phone = dtCus.Phone
}

func (c *CreatePayload) loadToData(dtCus *data.Customer) {
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
	ID
	CreatePayload
}

func (u *UpdatePayload) loadFromData(dtCus *data.Customer) {
	u.ID.loadFromData(dtCus)
	u.CreatePayload.loadFromData(dtCus)
}

func (u *UpdatePayload) loadToData(dtCus *data.Customer) {
	u.ID.loadToData(dtCus)
	u.CreatePayload.loadToData(dtCus)
}

// DetailPayload stores detail payload of the Customer
type DetailPayload struct {
	UpdatePayload

	Deleted    bool      `json:"deleted"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

func (d *DetailPayload) loadFromData(dtCus *data.Customer) {
	d.UpdatePayload.loadFromData(dtCus)

	d.Deleted = dtCus.Deleted
	d.CreatedAt = dtCus.CreatedAt
	d.CreatedBy = string(dtCus.CreatedBy)
	d.ModifiedAt = dtCus.ModifiedAt
	d.ModifiedBy = string(dtCus.ModifiedBy)
}

func (d *DetailPayload) loadToData(dtCus *data.Customer) {
	d.UpdatePayload.loadToData(dtCus)

	dtCus.Deleted = d.Deleted
	dtCus.CreatedAt = d.CreatedAt
	if bson.IsObjectIdHex(d.CreatedBy) {
		dtCus.CreatedBy = bson.ObjectIdHex(d.CreatedBy)
	}
	dtCus.ModifiedAt = d.ModifiedAt
	if bson.IsObjectIdHex(d.ModifiedBy) {
		dtCus.ModifiedBy = bson.ObjectIdHex(d.ModifiedBy)
	}
}
