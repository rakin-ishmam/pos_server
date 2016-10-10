package db

import (
	"github.com/rakin-ishmam/pos_server/data"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SellPayment provides db functionality for SellPayment
type SellPayment struct {
	Session *mgo.Session
}

// Put creates or updates SellPayment data
func (s SellPayment) Put(dtPay *data.SellPayment) error {
	if dtPay.ID == "" {
		dtPay.ID = bson.NewObjectId()
	}

	_, err := s.Session.DB("").C(sellPaymentC).UpsertId(dtPay.ID, dtPay)
	return err
}

// Get method takes an id of the SellPayment and returns a SellPayment object
func (s SellPayment) Get(id bson.ObjectId) (*data.SellPayment, error) {
	dtPay := &data.SellPayment{}

	if err := s.Session.DB("").C(sellPaymentC).FindId(id).One(dtPay); err != nil {
		return nil, err
	}

	return dtPay, nil
}

// List takes filter steps and return list of SellPayment
func (s SellPayment) List(skip, limit int, filters ...QueryFilter) ([]data.SellPayment, error) {

	query := bson.M{}
	for _, step := range filters {
		query = step.Filter(query)
	}

	payments := []data.SellPayment{}

	err := s.Session.DB("").C(sellPaymentC).Find(query).Skip(skip).Limit(limit).All(&payments)
	if err != nil {
		return nil, err
	}

	return payments, nil
}
