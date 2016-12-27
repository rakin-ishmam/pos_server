package main

import (
	"fmt"

	"github.com/rakin-ishmam/pos_server/action/admin"

	"gopkg.in/mgo.v2"
)

func main() {

	session, err := mongoSession()

	if err != nil {
		fmt.Println("connection error", err)
		return
	}

	if !createAdmin(session) {
		return
	}
}

func mongoSession() (*mgo.Session, error) {
	return mgo.Dial(mongoURI)
}

func createAdmin(ses *mgo.Session) bool {
	action := admin.Create{Session: ses}
	action.Do()
	if err := action.Err; err != nil {
		fmt.Println("admin create error", err)
		return false
	}
	return true
}
