package role

import (
	"github.com/rakin-ishmam/pos_server/action/geninfo"
	"github.com/rakin-ishmam/pos_server/data"
)

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

// LoadToData copy to data
func (c *CreatePayload) LoadToData(dtRole *data.Role) {
	dtRole.Name = c.Name
	dtRole.UserAccess = c.UserAccess
	dtRole.RoleAccess = c.RoleAccess
	dtRole.CategoryAccess = c.CategoryAccess
	dtRole.CustomerAccess = c.CustomerAccess
	dtRole.InventoryAccess = c.InventoryAccess
	dtRole.ProductAccess = c.ProductAccess
	dtRole.SellAccess = c.SellAccess
	dtRole.PaymentAccess = c.PaymentAccess
	dtRole.FileAccess = c.FileAccess
}

// LoadFromData copy data
func (c *CreatePayload) LoadFromData(dtRole *data.Role) {
	c.Name = dtRole.Name
	c.UserAccess = dtRole.UserAccess
	c.RoleAccess = dtRole.RoleAccess
	c.CustomerAccess = dtRole.CustomerAccess
	c.CategoryAccess = dtRole.CategoryAccess
	c.InventoryAccess = dtRole.InventoryAccess
	c.ProductAccess = dtRole.ProductAccess
	c.SellAccess = dtRole.SellAccess
	c.PaymentAccess = dtRole.PaymentAccess
	c.FileAccess = dtRole.FileAccess
}

// UpdatePayload stores Update data for role
type UpdatePayload struct {
	geninfo.ID

	CreatePayload
}

// LoadToData copy to data
func (u *UpdatePayload) LoadToData(dtRole *data.Role) {
	u.ID.LoadToData(&dtRole.Track)
	u.CreatePayload.LoadToData(dtRole)
}

// LoadFromData copy data
func (u *UpdatePayload) LoadFromData(dtRole *data.Role) {
	u.ID.LoadFromData(&dtRole.Track)
	u.CreatePayload.LoadFromData(dtRole)
}

// DetailPayload stores detail data for role
type DetailPayload struct {
	UpdatePayload

	geninfo.General
}

// LoadToData copy to data
func (d *DetailPayload) LoadToData(dtRole *data.Role) {
	d.UpdatePayload.LoadToData(dtRole)

	d.General.LoadToData(&dtRole.Track)
}

// LoadFromData copy data
func (d *DetailPayload) LoadFromData(dtRole *data.Role) {
	d.UpdatePayload.LoadFromData(dtRole)

	d.General.LoadFromData(&dtRole.Track)
}
