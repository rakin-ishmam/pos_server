package db

import (
	"github.com/rakin-ishmam/pos_server/db/query"
	mgo "gopkg.in/mgo.v2"
)

func make(c *mgo.Collection, q query.Query) *mgo.Query {
	bsn := q.Filter()
	query := c.Find(bsn)

	q.Skip(query)
	q.Limit(query)
	q.Order(query)

	return query
}
