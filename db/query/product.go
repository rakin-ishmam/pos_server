package query

import "gopkg.in/mgo.v2/bson"

// Product provides queries on product
type Product struct {
	chain []Applier
	GenInfo
}

// Filter make bson.M
func (p Product) Filter() bson.M {
	m := bson.M{}
	for _, v := range p.chain {
		v.Apply(m)
	}
	return m
}
