package query

import "gopkg.in/mgo.v2/bson"

// EqStep provides equal query step
type EqStep struct {
	Key   string
	Value interface{}
}

// Apply equal step
func (e EqStep) Apply(m bson.M) {
	m[e.Key] = e.Value
}
