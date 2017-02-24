package db

import (
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db/query"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Product provides db functionality for Product
type Product struct {
	Session *mgo.Session
}

// Put creates or updates Product data
func (p Product) Put(dtProduct *data.Product) error {
	if dtProduct.ID == "" {
		dtProduct.ID = bson.NewObjectId()
	}

	dtProduct.PreSave()

	_, err := p.Session.DB("").C(productC).UpsertId(dtProduct.ID, dtProduct)
	return err
}

// Get method takes an id of the Product and returns a Product object
func (p Product) Get(id bson.ObjectId) (*data.Product, error) {
	dtProduct := &data.Product{}

	if err := p.Session.DB("").C(productC).FindId(id).One(dtProduct); err != nil {
		return nil, err
	}

	return dtProduct, nil
}

// List returns list of Product
func (p Product) List(q query.Query) ([]data.Product, error) {
	mq := make(p.Session.DB("").C(productC), q)

	products := []data.Product{}

	err := mq.All(&products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
