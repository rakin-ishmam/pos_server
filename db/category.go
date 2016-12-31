package db

import (
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db/query"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Category provides db functionality for category
type Category struct {
	Session *mgo.Session
}

// Put creates or updates category data
func (c Category) Put(dtCategory *data.Category) error {
	if dtCategory.ID == "" {
		dtCategory.ID = bson.NewObjectId()
	}

	dtCategory.PreSave()

	_, err := c.Session.DB("").C(categoryC).UpsertId(dtCategory.ID, dtCategory)
	return err
}

// Get method takes an id of the Category and returns a Category object
func (c *Category) Get(id bson.ObjectId) (*data.Category, error) {
	dtCat := &data.Category{}

	if err := c.Session.DB("").C(categoryC).FindId(id).One(dtCat); err != nil {
		return nil, err
	}

	return dtCat, nil
}

// List takes filter steps and return list of Category
func (c Category) List(skip, limit int, filters ...query.Applier) ([]data.Category, error) {

	query := bson.M{}
	for _, step := range filters {
		step.Apply(query)
	}

	categories := []data.Category{}

	err := c.Session.DB("").C(categoryC).Find(query).Skip(skip).Limit(limit).All(&categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
