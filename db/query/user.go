package query

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User provides queries on user
type User struct {
	chain []Applier
	skip  int
	limit int
	order string
}

// EqUserName add EqStep for UserName
func (u *User) EqUserName(name string) *User {
	u.chain = append(u.chain, EqStep{
		Key:   "user_name",
		Value: name,
	})

	return u
}

// EqPassword add EqStep for Password
func (u *User) EqPassword(pass string) *User {
	u.chain = append(u.chain, EqStep{
		Key:   "password",
		Value: pass,
	})

	return u
}

// SetSkip set skip value
func (u *User) SetSkip(val int) {
	u.skip = val
}

// SetLimit set limit value
func (u *User) SetLimit(val int) {
	u.limit = val
}

// Filter make bson.M
func (u User) Filter() bson.M {
	m := bson.M{}
	for _, v := range u.chain {
		v.Apply(m)
	}
	return m
}

// Skip set skip value in query
func (u User) Skip(q *mgo.Query) {
	q.Skip(u.skip)
}

// Limit set limit value in query
func (u User) Limit(q *mgo.Query) {
	q.Limit(u.limit)
}

// Order add order value to query
func (u User) Order(q *mgo.Query) {
	if u.order != "" {
		q.Sort(u.order)
	}
}
