package role

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/rakin-ishmam/pos_server/data"
)

// ID stroe only id field of the role
type ID struct {
	ID string `json:"id, omitempty"`
}

func (i *ID) loadToData(dtRole *data.Role) {
	if bson.IsObjectIdHex(i.ID) {
		dtRole.ID = bson.ObjectIdHex(i.ID)
	}
}

func (i *ID) loadFromData(dtRole *data.Role) {
	i.ID = string(dtRole.ID)
}

// CreatePayload stores create data for role
type CreatePayload struct {
	Name            string          `json:"name"`
	UserAccess      data.AccessList `json:"user_access"`
	RoleAccess      data.AccessList `json:"role_access"`
	CategoryAccess  data.AccessList `json:"category_access"`
	CustomerAccess  data.AccessList `json:"customer_access"`
	InventoryAccess data.AccessList `json:"inventory_access"`
	ProductAccess   data.AccessList `json:"product_access"`
	SellAccess      data.AccessList `json:"sell_access"`
	PaymentAccess   data.AccessList `json:"payment_access"`
	FileAccess      data.AccessList `json:"file_access"`
}

func (c *CreatePayload) loadToData(dtRole *data.Role) {
	dtRole.Name = c.Name
	dtRole.UserAccess = c.UserAccess
	dtRole.RoleAccess = c.RoleAccess
	dtRole.CategoryAccess = c.CategoryAccess
	dtRole.InventoryAccess = c.InventoryAccess
	dtRole.ProductAccess = c.ProductAccess
	dtRole.SellAccess = c.SellAccess
	dtRole.PaymentAccess = c.PaymentAccess
	dtRole.FileAccess = c.FileAccess
}

func (c *CreatePayload) loadFromData(dtRole *data.Role) {
	c.Name = dtRole.Name
	c.UserAccess = dtRole.UserAccess
	c.RoleAccess = dtRole.RoleAccess
	c.CategoryAccess = dtRole.CategoryAccess
	c.InventoryAccess = dtRole.InventoryAccess
	c.ProductAccess = dtRole.ProductAccess
	c.SellAccess = dtRole.SellAccess
	c.PaymentAccess = dtRole.PaymentAccess
	c.FileAccess = dtRole.FileAccess
}

// UpdatePayload stores Update data for role
type UpdatePayload struct {
	ID

	CreatePayload
}

func (u *UpdatePayload) loadToData(dtRole *data.Role) {
	u.ID.loadToData(dtRole)
	u.CreatePayload.loadToData(dtRole)
}

func (u *UpdatePayload) loadFromData(dtRole *data.Role) {
	u.ID.loadFromData(dtRole)
	u.CreatePayload.loadFromData(dtRole)
}

// DetailPayload stores detail data for role
type DetailPayload struct {
	UpdatePayload

	Deleted    bool      `json:"deleted"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

func (d *DetailPayload) loadToData(dtRole *data.Role) {
	d.UpdatePayload.loadToData(dtRole)

	dtRole.Deleted = d.Deleted
	dtRole.CreatedAt = d.CreatedAt
	if bson.IsObjectIdHex(d.CreatedBy) {
		dtRole.CreatedBy = bson.ObjectIdHex(d.CreatedBy)
	}
	dtRole.ModifiedAt = d.ModifiedAt
	if bson.IsObjectIdHex(d.ModifiedBy) {
		dtRole.ModifiedBy = bson.ObjectIdHex(d.ModifiedBy)
	}
}

func (d *DetailPayload) loadFromData(dtRole *data.Role) {
	d.UpdatePayload.loadFromData(dtRole)

	d.Deleted = dtRole.Deleted
	d.CreatedAt = dtRole.CreatedAt
	d.CreatedBy = string(dtRole.CreatedBy)
	d.ModifiedAt = dtRole.ModifiedAt
	d.ModifiedBy = string(dtRole.ModifiedBy)
}
