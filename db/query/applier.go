package query

import "gopkg.in/mgo.v2/bson"

// Applier apply query steps
type Applier interface {
	Apply(m bson.M)
}
