package data

import (
	"github.com/rakin-ishmam/pos_server/apperr"
	"gopkg.in/mgo.v2/bson"
)

// SellPayment provides payment information of sell
type SellPayment struct {
	Track

	SellID      bson.ObjectId `bson:"sell_id,omitempty"`
	Amount      float64       `bson:"amount"`
	PaymentType PaymentType   `bson:"payment_type"`
	Number      string        `bson:"number"`
	Comment     string        `bson:"comment"`
}

// Validate valids OrderPayment data
func (s SellPayment) Validate() error {
	if !s.PaymentType.Validate() {
		return apperr.Validation{
			Where: "SellPayment",
			Field: "payment_type",
			Cause: apperr.StrInvalid,
		}
	}

	return nil
}
