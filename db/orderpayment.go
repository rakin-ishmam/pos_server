package db

import (
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db/query"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// OrderPayment provides db functionality for OrderPayment
type OrderPayment struct {
	Session *mgo.Session
}

// Put creates or updates OrderPayment data
func (o OrderPayment) Put(dtOrderPay *data.OrderPayment) error {
	if dtOrderPay.ID == "" {
		dtOrderPay.ID = bson.NewObjectId()
	}

	_, err := o.Session.DB("").C(orderPaymentC).UpsertId(dtOrderPay.ID, dtOrderPay)
	return err
}

// Get method takes an id of the OrderPayment and returns a OrderPayment object
func (o OrderPayment) Get(id bson.ObjectId) (*data.OrderPayment, error) {
	dtOrderPay := &data.OrderPayment{}

	if err := o.Session.DB("").C(orderPaymentC).FindId(id).One(dtOrderPay); err != nil {
		return nil, err
	}

	return dtOrderPay, nil
}

// List takes filter steps and return list of OrderPayment
func (o OrderPayment) List(skip, limit int, filters ...query.Applier) ([]data.OrderPayment, error) {

	query := bson.M{}
	for _, step := range filters {
		step.Apply(query)
	}

	payments := []data.OrderPayment{}

	err := o.Session.DB("").C(orderPaymentC).Find(query).Skip(skip).Limit(limit).All(&payments)
	if err != nil {
		return nil, err
	}

	return payments, nil
}
