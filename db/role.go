package db

import (
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db/query"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Role provides db functionality for Role
type Role struct {
	Session *mgo.Session
}

// Put creates or updates Role data
func (r Role) Put(dtRole *data.Role) error {
	if dtRole.ID == "" {
		dtRole.ID = bson.NewObjectId()
	}

	dtRole.PreSave()

	_, err := r.Session.DB("").C(roleC).UpsertId(dtRole.ID, dtRole)
	return err
}

// Get method takes an id of the Role and returns a Role object
func (r Role) Get(id bson.ObjectId) (*data.Role, error) {
	dtRole := &data.Role{}

	if err := r.Session.DB("").C(roleC).FindId(id).One(dtRole); err != nil {
		return nil, err
	}

	return dtRole, nil
}

// List takes filter steps and return list of Role
func (r Role) List(skip, limit int, filters ...query.Applier) ([]data.Role, error) {

	query := bson.M{}
	for _, step := range filters {
		step.Apply(query)
	}

	roles := []data.Role{}

	err := r.Session.DB("").C(roleC).Find(query).Skip(skip).Limit(limit).All(&roles)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
