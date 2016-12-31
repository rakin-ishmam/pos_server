package db

import (
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db/query"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Sell provides db functionality for Sell
type Sell struct {
	Session *mgo.Session
}

// Put creates or updates Sell data
func (s Sell) Put(dtSell *data.Sell) error {
	if dtSell.ID == "" {
		dtSell.ID = bson.NewObjectId()
	}

	_, err := s.Session.DB("").C(sellC).UpsertId(dtSell.ID, dtSell)
	return err
}

// Get method takes an id of the Sell and returns a Sell object
func (s Sell) Get(id bson.ObjectId) (*data.Sell, error) {
	dtSell := &data.Sell{}

	if err := s.Session.DB("").C(sellC).FindId(id).One(dtSell); err != nil {
		return nil, err
	}

	return dtSell, nil
}

// List takes filter steps and return list of Sell
func (s Sell) List(skip, limit int, filters ...query.Applier) ([]data.Sell, error) {

	query := bson.M{}
	for _, step := range filters {
		step.Apply(query)
	}

	sells := []data.Sell{}

	err := s.Session.DB("").C(sellC).Find(query).Skip(skip).Limit(limit).All(&sells)
	if err != nil {
		return nil, err
	}

	return sells, nil
}
