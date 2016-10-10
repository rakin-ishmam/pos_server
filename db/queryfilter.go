package db

import "gopkg.in/mgo.v2/bson"

// QueryFilter is implemented by struct who wants to provide filter step
type QueryFilter interface {
	Filter(m bson.M) bson.M
}
