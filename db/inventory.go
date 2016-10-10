package db

import (
	"github.com/rakin-ishmam/pos_server/data"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Inventory provides db functionality for Inventory
type Inventory struct {
	Session *mgo.Session
}

// Put creates or updates Inventory data
func (f *Inventory) Put(dtInv *data.Inventory) error {
	if dtInv.ID == "" {
		dtInv.ID = bson.NewObjectId()
	}

	_, err := f.Session.DB("").C(inventoryC).UpsertId(dtInv.ID, dtInv)
	return err
}

// Get method takes an id of the Inventory and returns a Inventory object
func (f Inventory) Get(id bson.ObjectId) (*data.Inventory, error) {
	dtInv := &data.Inventory{}

	if err := f.Session.DB("").C(inventoryC).FindId(id).One(dtInv); err != nil {
		return nil, err
	}

	return dtInv, nil
}

// List takes filter steps and return list of Inventory
func (f Inventory) List(skip, limit int, filters ...QueryFilter) ([]data.Inventory, error) {

	query := bson.M{}
	for _, step := range filters {
		query = step.Filter(query)
	}

	inventories := []data.Inventory{}

	err := f.Session.DB("").C(inventoryC).Find(query).Skip(skip).Limit(limit).All(&inventories)
	if err != nil {
		return nil, err
	}

	return inventories, nil
}
