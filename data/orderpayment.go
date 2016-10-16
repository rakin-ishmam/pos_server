package data

import (
	"github.com/rakin-ishmam/pos_server/apperr"
	"gopkg.in/mgo.v2/bson"
)

// OrderPayment provides payment information
type OrderPayment struct {
	Track

	OrderID     bson.ObjectId `bson:"order_id,omitempty"`
	Amount      float64       `bson:"amount"`
	PaymentType PaymentType   `bson:"payment_type"`
	Number      string        `bson:"number"`
	Comment     string        `bson:"comment"`
}

// Validate valids OrderPayment data
func (o OrderPayment) Validate() error {
	if !o.PaymentType.Validate() {
		return apperr.Validation{
			Where: "OrderPayment",
			Field: "payment_type",
			Cause: apperr.ValidationInvalid,
		}
	}

	return nil
}
