package user

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db/query"
)

// LoginPayload store login info
type LoginPayload struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (l LoginPayload) query() []query.Applier {
	return []query.Applier{
		query.EqStep{Key: "user_name", Value: l.UserName},
		query.EqStep{Key: "password", Value: l.Password},
	}
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

// LoadToData copy to data
func (c *CreatePayload) LoadToData(dtUser *data.User) {
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

// LoadFromData copy data
func (c *CreatePayload) LoadFromData(dtUser *data.User) {
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
	geninfo.ID

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

// LoadToData copy to data
func (u *UpdatePayload) LoadToData(dtUser *data.User) {
	u.ID.LoadToData(&dtUser.Track)

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

// LoadFromData copy data
func (u *UpdatePayload) LoadFromData(dtUser *data.User) {
	u.ID.LoadFromData(&dtUser.Track)

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
	geninfo.ID
	UpdatePayload

	geninfo.General
}

// LoadFromData copy data
func (d *DetailPayload) LoadFromData(dtUser *data.User) {
	d.ID.LoadFromData(&dtUser.Track)
	d.UpdatePayload.LoadFromData(dtUser)

	d.General.LoadFromData(&dtUser.Track)
}
