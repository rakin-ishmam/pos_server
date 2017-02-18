package query

// User provides queries on user
type User struct {
	chain []Applier
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

// Query returns list of Applier
func (u User) Query() []Applier {
	return u.chain
}
