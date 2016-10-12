package db

import (
	"github.com/rakin-ishmam/pos_server/data"
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
func (u User) List(ukip, limit int, filters ...QueryFilter) ([]data.User, error) {

	query := bson.M{}
	for _, step := range filters {
		query = step.Filter(query)
	}

	users := []data.User{}

	err := u.Session.DB("").C(userC).Find(query).Skip(ukip).Limit(limit).All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
