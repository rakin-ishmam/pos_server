package data

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Track provides information of creation and modification
type Track struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Search     []string      `bson:"search"`
	Deleted    bool          `bson:"deleted"`
	CreatedAt  time.Time     `bson:"created_at"`
	CreatedBy  bson.ObjectId `bson:"created_by,omitempty"`
	ModifiedAt time.Time     `bson:"modified_at"`
	ModifiedBy bson.ObjectId `bson:"modified_by,omitempty"`
}

// BeforeUpdate takes userID and  updates necessary fields
func (t *Track) BeforeUpdate(userID bson.ObjectId) {
	t.ModifiedAt = time.Now()
	if userID != "" {
		t.ModifiedBy = userID
	}
}

// BeforeCreate takes userID and updates necessary fields
func (t *Track) BeforeCreate(userID bson.ObjectId) {
	t.CreatedAt = time.Now()
	t.ModifiedAt = t.CreatedAt

	if userID != "" {
		t.ModifiedBy = userID
		t.CreatedBy = userID
	}
}
