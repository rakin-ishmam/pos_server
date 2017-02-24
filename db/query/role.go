package query

import "gopkg.in/mgo.v2/bson"

// Role provides queries on role
type Role struct {
	chain []Applier
	GenInfo
}

// Filter make bson.M
func (r Role) Filter() bson.M {
	m := bson.M{}
	for _, v := range r.chain {
		v.Apply(m)
	}
	return m
}
