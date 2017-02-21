package user

import (
	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/auth"
	"github.com/rakin-ishmam/pos_server/config"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	"github.com/rakin-ishmam/pos_server/db/query"
	mgo "gopkg.in/mgo.v2"
)

// TokenToUser manage to get user and role from login token
type TokenToUser struct {
	session *mgo.Session
	token   string
	err     error
	resUser data.User
	resRole data.Role
}

// Do run action
func (t *TokenToUser) Do() {
	info, err := auth.Decode(t.token, config.TokenSecret)
	if err != nil {
		t.err = err
		return
	}

	usrQ := query.User{}
	usrQ.EqUserName(info.UserName)
	usrQ.SetSkip(0)
	usrQ.SetLimit(1)

	dbUsr := db.User{Session: t.session}
	dtUsrs, err := dbUsr.List(usrQ)
	if err != nil {
		t.err = apperr.NewDatabase(t.err, "user", "TokenToUser")
		return
	}

	if len(dtUsrs) == 0 {
		t.err = apperr.NewNotfound("token to user", "user")
		return
	}

	t.resUser = dtUsrs[0]

	dbRole := db.Role{Session: t.session}
	dtRole, err := dbRole.Get(t.resUser.RoleID)
	if err != nil {
		t.err = apperr.NewDatabase(err, "role", "TokenToUser")
		return
	} else if dtRole == nil {
		t.err = apperr.NewNotfound("token to user", "role")
		return
	}

	t.resRole = *dtRole

}

// Result returns User and user's Role
func (t TokenToUser) Result() (data.User, data.Role, error) {
	if t.err != nil {
		return t.resUser, t.resRole, t.err
	}

	return t.resUser, t.resRole, nil
}

// NewTokenToUser returns action TokenToUser
func NewTokenToUser(session *mgo.Session, token string) *TokenToUser {
	return &TokenToUser{session: session, token: token}
}
