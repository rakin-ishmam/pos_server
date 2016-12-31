package user

import (
	"io"

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
	dtUsers, err := dbUser.List(0, 1, l.ReqPayload.query()[0])
	// err := dbUser.Put(dtUser)
	// if err != nil {
	// 	c.Err = apperr.Database{
	// 		Base:   err,
	// 		Where:  "User",
	// 		Action: "Create",
	// 	}
	// 	return
	// }

	// c.ResPayload = geninfo.ID{ID: string(dtUser.ID)}
}

// Result returns result of the action
func (l Login) Result() (io.Reader, error) {
	// if c.Err != nil {
	// 	return nil, c.Err
	// }

	// return converter.JSONtoBuff(c.ResPayload)
	return nil, nil
}
