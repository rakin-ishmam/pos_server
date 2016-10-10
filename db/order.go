package db

import (
	"github.com/rakin-ishmam/pos_server/data"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Order provides db functionality for Order
type Order struct {
	Session *mgo.Session
}

// Put creates or updates Order data
func (o Order) Put(dtOrder *data.Order) error {
	if dtOrder.ID == "" {
		dtOrder.ID = bson.NewObjectId()
	}

	_, err := o.Session.DB("").C(orderC).UpsertId(dtOrder.ID, dtOrder)
	return err
}

// Get method takes an id of the Order and returns a Order object
func (o Order) Get(id bson.ObjectId) (*data.Order, error) {
	dtOrder := &data.Order{}

	if err := o.Session.DB("").C(orderC).FindId(id).One(dtOrder); err != nil {
		return nil, err
	}

	return dtOrder, nil
}

// List takes filter steps and return list of Order
func (o Order) List(skip, limit int, filters ...QueryFilter) ([]data.Order, error) {

	query := bson.M{}
	for _, step := range filters {
		query = step.Filter(query)
	}

	orders := []data.Order{}

	err := o.Session.DB("").C(orderC).Find(query).Skip(skip).Limit(limit).All(&orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
