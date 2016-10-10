package data

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Track provides information of creation and modification
type Track struct {
	ID         bson.ObjectId `bson:"_id, omitempty"`
	Deleted    bool          `bson:"deleted"`
	CreatedAt  time.Time     `bson:"created_at"`
	CreatedBy  bson.ObjectId `bson:"created_by, omitempty"`
	ModifiedAt time.Time     `bson:"modified_at"`
	ModifiedBy bson.ObjectId `bson:"modified_by, omitempty"`
}
