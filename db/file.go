package db

import (
	"github.com/rakin-ishmam/pos_server/data"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// File provides db functionality for File
type File struct {
	Session *mgo.Session
}

// Put creates or updates File data
func (f *File) Put(dtFile *data.File) error {
	if dtFile.ID == "" {
		dtFile.ID = bson.NewObjectId()
	}

	_, err := f.Session.DB("").C(fileC).UpsertId(dtFile.ID, dtFile)
	return err
}

// Get method takes an id of the File and returns a File object
func (f File) Get(id bson.ObjectId) (*data.File, error) {
	dtFile := &data.File{}

	if err := f.Session.DB("").C(fileC).FindId(id).One(dtFile); err != nil {
		return nil, err
	}

	return dtFile, nil
}

// List takes filter steps and return list of File
func (f File) List(skip, limit int, filters ...QueryFilter) ([]data.File, error) {

	query := bson.M{}
	for _, step := range filters {
		query = step.Filter(query)
	}

	files := []data.File{}

	err := f.Session.DB("").C(fileC).Find(query).Skip(skip).Limit(limit).All(&files)
	if err != nil {
		return nil, err
	}

	return files, nil
}
