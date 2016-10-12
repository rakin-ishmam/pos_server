package db

import (
	"github.com/rakin-ishmam/pos_server/data"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Customer provides db functionality for Customer
type Customer struct {
	Session *mgo.Session
}

// Put creates or updates Customer data
func (c *Customer) Put(dtCus *data.Customer) error {
	if dtCus.ID == "" {
		dtCus.ID = bson.NewObjectId()
	}

	dtCus.PreSave()

	_, err := c.Session.DB("").C(customerC).UpsertId(dtCus.ID, dtCus)
	return err
}

// Get method takes an id of the Customer and returns a Customer object
func (c Customer) Get(id bson.ObjectId) (*data.Customer, error) {
	dtCus := &data.Customer{}

	if err := c.Session.DB("").C(customerC).FindId(id).One(dtCus); err != nil {
		return nil, err
	}

	return dtCus, nil
}

// List takes filter steps and return list of Customer
func (c Customer) List(skip, limit int, filters ...QueryFilter) ([]data.Customer, error) {

	query := bson.M{}
	for _, step := range filters {
		query = step.Filter(query)
	}

	customers := []data.Customer{}

	err := c.Session.DB("").C(customerC).Find(query).Skip(skip).Limit(limit).All(&customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}
