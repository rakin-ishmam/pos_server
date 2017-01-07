package main

import (
	"fmt"
	"net/http"

	"github.com/rakin-ishmam/pos_server/action/admin"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	test()
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

	http.ListenAndServe(":8080", r)
}

func applyRoutes(r *mux.Router, session *mgo.Session) {
	for _, v := range allRoutes(session) {
		fmt.Println("login")
		r.Methods(v.Method).Name(v.Name).Path(v.Path).HandlerFunc(logRoute(v))
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

func test() {
	b := bson.M{}
	test2(b)
	fmt.Println(b)
}

func test2(b bson.M) {
	b["hagu"] = 1
}
