package data

import (
	"github.com/rakin-ishmam/pos_server/apperr"
	"gopkg.in/mgo.v2/bson"
)

// Category represents category of product
type Category struct {
	Track

	Name       string        `bson:"name"`
	CategoryID bson.ObjectId `bson:"category_id, omitempty"`
}

// PreSave takes the necessary step before saving data
func (c *Category) PreSave() {
	c.Track.Search = Spliter(c.Name)
}

// Validate valides Category
func (c Category) Validate() error {
	if c.Name == "" {
		return apperr.Validation{
			Where: "Category",
			Field: "name",
			Cause: apperr.ValidationEmpty,
		}
	}

	return nil
}
