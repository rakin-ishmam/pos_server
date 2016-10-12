package data

// Order provides info about pos order
type Order struct {
	Track

	Code         string         `bson:"code"`
	OrderProduct []OrderProduct `bson:"order_product"`
	TotalPrice   float64        `bson:"total_price"`
	TotalPaid    float64        `bson:"total_paid"`
	Discount     float64        `bson:"discount"`
	Delivered    bool           `bson:"delivered"`
}

// OrderProduct provides info about one product in the order
type OrderProduct struct {
	Product  Product `bson:"product"`
	Quantity int     `bson:"quantity"`
}

// PreSave takes the necessary step before saving data
func (o *Order) PreSave() {
	o.Track.Search = Spliter(o.Code)
}

// Validate valids data
func (o Order) Validate() error {
	return nil
}
