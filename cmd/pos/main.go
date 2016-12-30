package main

import (
	"fmt"

	"github.com/rakin-ishmam/pos_server/action/admin"

	"github.com/gorilla/mux"
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

	r := mux.NewRouter()
	applyRoutes(r, session)

}

func applyRoutes(r *mux.Router, session *mgo.Session) {
	for _, v := range allRoutes(session) {
		r.Methods(v.Metod).Name(v.Name).Path(v.Path).HandlerFunc(v.Handler)
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
