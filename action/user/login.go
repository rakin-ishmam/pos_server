package user

import (
	"io"
	"time"

	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/auth"
	"github.com/rakin-ishmam/pos_server/config"
	"github.com/rakin-ishmam/pos_server/converter"
	"github.com/rakin-ishmam/pos_server/db"
	"github.com/rakin-ishmam/pos_server/db/query"
	mgo "gopkg.in/mgo.v2"
)

// Login manages login process
type Login struct {
	Session    *mgo.Session
	ReqPayload LoginPayload
	token      Token
	err        error
}

// Do generate token
func (l *Login) Do() {
	userQ := query.User{}
	userQ.EqPassword(l.ReqPayload.Password).EqUserName(l.ReqPayload.Username)
	userQ.SetSkip(0)
	userQ.SetLimit(1)

	dbUser := db.User{Session: l.Session}
	dtUsers, err := dbUser.List(userQ)

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

	info := auth.Info{Username: dtUsers[0].Username, Exp: time.Now()}
	l.token.Token, l.err = auth.New(info, config.TokenSecret)
}

// Result returns result of the action
func (l Login) Result() (io.Reader, error) {
	if l.err != nil {
		return nil, l.err
	}

	return converter.JSONtoBuff(l.token)
}
