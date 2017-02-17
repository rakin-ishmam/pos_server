package user

import (
	"github.com/rakin-ishmam/pos_server/auth"
	"github.com/rakin-ishmam/pos_server/config"
	mgo "gopkg.in/mgo.v2"
)

// TokenToUser manage to get user and role from login token
type TokenToUser struct {
	session *mgo.Session
	token   string
	err     error
}

// Do run action
func (t *TokenToUser) Do() {
	info, err := auth.Decode(t.token, config.TokenSecret)
	if err != nil {
		t.err = err
		return
	}

}
