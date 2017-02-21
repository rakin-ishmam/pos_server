package db

import (
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db/query"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User provides db functionality for User
type User struct {
	Session *mgo.Session
}

// Put creates or updates User data
func (u User) Put(dtUser *data.User) error {
	if dtUser.ID == "" {
		dtUser.ID = bson.NewObjectId()
	}

	dtUser.PreSave()

	_, err := u.Session.DB("").C(userC).UpsertId(dtUser.ID, dtUser)
	return err
}

// Get method takes an id of the User and returns a User object
func (u User) Get(id bson.ObjectId) (*data.User, error) {
	dtUser := &data.User{}

	if err := u.Session.DB("").C(userC).FindId(id).One(dtUser); err != nil {
		return nil, err
	}

	return dtUser, nil
}

// List takes filter steps and return list of User
func (u User) List(q query.Query) ([]data.User, error) {

	mq := make(u.Session.DB("").C(userC), q)

	users := []data.User{}

	err := mq.All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
