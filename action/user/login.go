package user

import (
	"bytes"
	"io"

	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/db"
	mgo "gopkg.in/mgo.v2"
)

// Login manages login process
type Login struct {
	Session    *mgo.Session
	ReqPayload LoginPayload
	token      string
	err        error
}

// Do generate token
func (l *Login) Do() {

	dbUser := db.User{Session: l.Session}
	dtUsers, err := dbUser.List(0, 1, l.ReqPayload.query())

	if err != nil {
		l.err = apperr.Database{
			Base:   err,
			Where:  "User",
			Action: "Login",
		}
		return
	}

	if len(dtUsers) == 0 {
		l.err = apperr.Notfound{
			Where: apperr.StrLogin,
			What:  apperr.StrUser + "/" + apperr.StrPassword,
		}
		return
	}

	l.token = dtUsers[0].UserName
}

// Result returns result of the action
func (l Login) Result() (io.Reader, error) {
	if l.err != nil {
		return nil, l.err
	}

	buff := bytes.NewBuffer([]byte(l.token))
	return buff, nil
}
