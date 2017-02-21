package query

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Applier apply query steps
type Applier interface {
	Apply(m bson.M)
}

// Query provides funtionality to make query
type Query interface {
	Filter() bson.M
	Skip(*mgo.Query)
	Limit(*mgo.Query)
	Order(*mgo.Query)
}
