package data

import "github.com/rakin-ishmam/pos_server/apperr"

// Role represents user role
type Role struct {
	Track

	Name            string     `bson:"name"`
	UserAccess      AccessList `bson:"user_access"`
	RoleAccess      AccessList `bson:"role_access"`
	CategoryAccess  AccessList `bson:"category_access"`
	CustomerAccess  AccessList `bson:"customer_access"`
	InventoryAccess AccessList `bson:"inventory_access"`
	ProductAccess   AccessList `bson:"product_access"`
	SellAccess      AccessList `bson:"sell_access"`
	PaymentAccess   AccessList `bson:"payment_access"`
	FileAccess      AccessList `bson:"file_access"`
}

// PreSave takes the necessary step before saving data
func (r *Role) PreSave() {
	r.Track.Search = Spliter(r.Name)
}

// Validate returns Validation error if Role has invalid field
func (r Role) Validate() error {
	accessList := []struct {
		Field      string
		AccessList AccessList
	}{
		{"user_access", r.UserAccess},
		{"role_access", r.RoleAccess},
		{"category_access", r.CategoryAccess},
		{"customer_access", r.CustomerAccess},
		{"inventory_access", r.InventoryAccess},
		{"product_access", r.ProductAccess},
		{"aell_access", r.SellAccess},
		{"payment_access", r.PaymentAccess},
		{"file_access", r.FileAccess},
	}

	for _, v := range accessList {
		if !v.AccessList.Validate() {
			return apperr.Validation{Where: "Role", Field: v.Field, Cause: apperr.ValidationInvalid}
		}
	}

	return nil
}
