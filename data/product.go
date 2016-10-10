package data

import (
	"github.com/rakin-ishmam/pos_server/apperr"
	"gopkg.in/mgo.v2/bson"
)

// Product represents pos product
type Product struct {
	Track

	Name        string        `bson:"name"`
	ProductType ProductType   `bson:"product_type"`
	Code        string        `bson:"code"`
	AvtFileID   string        `bson:"avt_file_id"`
	CategoryID  bson.ObjectId `bson:"category_id, omitempty"`
	SalePrice   float64       `bson:"sale_price"`
	BuyPrice    float64       `bson:"buy_price"`
	Quantity    int           `bson:"quantity"`
}

// Validate valids Product data
func (p *Product) Validate() error {
	if p.Code == "" {
		return apperr.Validation{
			Where: "Product",
			Field: "code",
			Cause: apperr.ValidationEmpty,
		}
	}

	return nil
}
