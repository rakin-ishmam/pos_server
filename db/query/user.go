package query

import "gopkg.in/mgo.v2/bson"

// User provides queries on user
type User struct {
	chain []Applier
	GenInfo
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

// Filter make bson.M
func (u User) Filter() bson.M {
	m := bson.M{}
	for _, v := range u.chain {
		v.Apply(m)
	}
	return m
}
