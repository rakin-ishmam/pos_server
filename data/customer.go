package data

import (
	"github.com/rakin-ishmam/pos_server/apperr"
	"gopkg.in/mgo.v2/bson"
)

// Customer provides customer info
type Customer struct {
	Track

	Name      string        `bson:"name"`
	AvtFileID bson.ObjectId `bson:"avt_file_id, omitempty"`
	Code      string        `bson:"code"`
	Email     string        `bson:"email"`
	Gender    Gender        `bson:"gender"`
	Address   string        `bson:"address"`
	Phone     []string      `bson:"phone"`
}

// Validate valids Customer data
func (c *Customer) Validate() error {
	if c.Name == "" {
		return apperr.Validation{
			Where: "Customer",
			Field: "name",
			Cause: apperr.ValidationEmpty,
		}
	}

	if !c.Gender.Validate() {
		return apperr.Validation{
			Where: "Customer",
			Field: "gender",
			Cause: apperr.ValidationInvalid,
		}
	}

	return nil
}
