package data

import "gopkg.in/mgo.v2/bson"

// Inventory represents product inventory
type Inventory struct {
	Track

	Code      string        `bson:"code"`
	ProductID bson.ObjectId `bson:"product_id, omitempty"`
	SalePrice float64       `bson:"sale_price"`
	BuyPrice  float64       `bson:"buy_price"`
	Quantity  int           `bson:"quantity"`
}

// PreSave takes the necessary step before saving data
func (inv *Inventory) PreSave() {
	inv.Track.Search = Spliter(inv.Code)
}

// Validate valids Inventory data
func (inv Inventory) Validate() error {
	return nil
}
