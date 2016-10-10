package data

// Sell provide information about pos sell
type Sell struct {
	Track

	SellProduct []SellProduct `bson:"sell_product"`
	TotalPrice  float64       `bson:"total_price"`
	TotalPaid   float64       `bson:"total_paid"`
	Discount    float64       `bson:"discount"`
}

// SellProduct provides info about one product in the sale
type SellProduct struct {
	Product  Product `bson:"product"`
	Quantity int     `bson:"quantity"`
}
