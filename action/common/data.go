package common

import (
	"time"

	"github.com/rakin-ishmam/pos_server/data"
	"gopkg.in/mgo.v2/bson"
)

// ID stroe only id field
type ID struct {
	ID string `json:"id, omitempty"`
}

// LoadFromData copy from Data
func (i *ID) LoadFromData(dt *data.Track) {
	i.ID = string(dt.ID)
}

// LoadToData copy to Data
func (i *ID) LoadToData(dt *data.Track) {
	if bson.IsObjectIdHex(i.ID) {
		dt.ID = bson.ObjectIdHex(i.ID)
	}
}

// General stores general info of data
type General struct {
	Deleted    bool      `json:"deleted"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

// LoadFromData copy from Data
func (g *General) LoadFromData(dt *data.Track) {
	g.Deleted = dt.Deleted
	g.CreatedAt = dt.CreatedAt
	g.CreatedBy = string(dt.CreatedBy)
	g.ModifiedAt = dt.ModifiedAt
	g.ModifiedBy = string(dt.ModifiedBy)
}

// LoadToData copy to Data
func (g *General) LoadToData(dt *data.Track) {
	dt.Deleted = g.Deleted
	dt.CreatedAt = g.CreatedAt
	if bson.IsObjectIdHex(g.CreatedBy) {
		dt.CreatedBy = bson.ObjectIdHex(g.CreatedBy)
	}
	dt.ModifiedAt = g.ModifiedAt
	if bson.IsObjectIdHex(g.ModifiedBy) {
		dt.ModifiedBy = bson.ObjectIdHex(g.ModifiedBy)
	}
}
