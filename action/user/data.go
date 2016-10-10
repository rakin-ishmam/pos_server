package user

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/rakin-ishmam/pos_server/data"
)

// ID stroe only id field of the role
type ID struct {
	ID string `json:"id, omitempty"`
}

func (i *ID) loadToData(dtUser *data.User) {
	if bson.IsObjectIdHex(i.ID) {
		dtUser.ID = bson.ObjectIdHex(i.ID)
	}
}

func (i *ID) loadFromData(dtUser *data.User) {
	i.ID = string(dtUser.ID)
}

// CreatePayload stores create data for User
type CreatePayload struct {
	Name      string   `json:"name"`
	UserName  string   `json:"user_name"`
	Language  string   `json:"language"`
	AvtFileID string   `json:"avt_file_id"`
	RoleID    string   `json:"role_id"`
	Email     string   `json:"email"`
	Address   string   `json:"address"`
	Gender    string   `json:"gender"`
	Phone     []string `json:"phone"`
	Password  string   `json:"password"`
}

func (c *CreatePayload) loadToData(dtUser *data.User) {
	dtUser.Name = c.Name
	dtUser.UserName = c.UserName
	dtUser.Language = data.Language(c.Language)
	dtUser.AvtFileID = c.AvtFileID
	if bson.IsObjectIdHex(c.RoleID) {
		dtUser.RoleID = bson.ObjectIdHex(c.RoleID)
	}
	dtUser.Email = c.Email
	dtUser.Email = c.Address
	dtUser.Gender = data.Gender(c.Gender)
	dtUser.Phone = c.Phone
	dtUser.Password = c.Password
}

func (c *CreatePayload) loadFromData(dtUser *data.User) {
	c.Name = dtUser.Name
	c.UserName = dtUser.UserName
	c.Language = string(dtUser.Language)
	c.AvtFileID = dtUser.AvtFileID
	c.RoleID = string(dtUser.RoleID)
	c.Email = dtUser.Email
	c.Address = dtUser.Address
	c.Gender = string(dtUser.Gender)
	c.Phone = dtUser.Phone
	c.Password = dtUser.Password
}

// UpdatePayload stores update data for User
type UpdatePayload struct {
	ID

	Name      string   `json:"name"`
	UserName  string   `json:"user_name"`
	Language  string   `json:"language"`
	AvtFileID string   `json:"avt_file_id"`
	RoleID    string   `json:"role_id"`
	Email     string   `json:"email"`
	Address   string   `json:"address"`
	Gender    string   `json:"gender"`
	Phone     []string `json:"phone"`
}

func (u *UpdatePayload) loadToData(dtUser *data.User) {
	u.ID.loadToData(dtUser)

	dtUser.Name = u.Name
	dtUser.UserName = u.UserName
	dtUser.Language = data.Language(u.Language)
	dtUser.AvtFileID = u.AvtFileID
	if bson.IsObjectIdHex(u.RoleID) {
		dtUser.RoleID = bson.ObjectIdHex(u.RoleID)
	}
	dtUser.Email = u.Email
	dtUser.Email = u.Address
	dtUser.Gender = data.Gender(u.Gender)
	dtUser.Phone = u.Phone
}

func (u *UpdatePayload) loadFromData(dtUser *data.User) {
	u.ID.loadFromData(dtUser)

	u.Name = dtUser.Name
	u.UserName = dtUser.UserName
	u.Language = string(dtUser.Language)
	u.AvtFileID = dtUser.AvtFileID
	u.RoleID = string(dtUser.RoleID)
	u.Email = dtUser.Email
	u.Address = dtUser.Address
	u.Gender = string(dtUser.Gender)
	u.Phone = dtUser.Phone
}

// DetailPayload stores detail data for User
type DetailPayload struct {
	ID
	UpdatePayload

	Deleted    bool      `json:"deleted"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

func (d *DetailPayload) loadFromData(dtUser *data.User) {
	d.ID.loadFromData(dtUser)
	d.UpdatePayload.loadFromData(dtUser)

	d.Deleted = dtUser.Deleted
	d.CreatedAt = dtUser.CreatedAt
	d.CreatedBy = string(dtUser.CreatedBy)
	d.ModifiedAt = dtUser.ModifiedAt
	d.ModifiedBy = string(dtUser.ModifiedBy)
}
