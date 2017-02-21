package query

import mgo "gopkg.in/mgo.v2"

// GenInfo store general info of query
type GenInfo struct {
	skip  int
	limit int
	order string
}

// SetSkip set skip value
func (g *GenInfo) SetSkip(val int) {
	g.skip = val
}

// SetLimit set limit value
func (g *GenInfo) SetLimit(val int) {
	g.limit = val
}

// SetOrder set order val
func (g *GenInfo) SetOrder(val string) {
	g.order = val
}

// Skip set skip value in query
func (g GenInfo) Skip(q *mgo.Query) {
	q.Skip(g.skip)
}

// Limit set limit value in query
func (g GenInfo) Limit(q *mgo.Query) {
	q.Limit(g.limit)
}

// Order add order value to query
func (g GenInfo) Order(q *mgo.Query) {
	if g.order != "" {
		q.Sort(g.order)
	}
}
