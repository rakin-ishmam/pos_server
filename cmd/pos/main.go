package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rakin-ishmam/pos_server/action/admin"
	"github.com/rakin-ishmam/pos_server/config"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	test()
	session, err := mongoSession()
	fmt.Println("mongo uri", config.MongoURI, "+", config.TokenSecret)
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
	fmt.Println("login")
}

func mongoSession() (*mgo.Session, error) {
	return mgo.Dial(config.MongoURI)
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
	var id bson.ObjectId
	id = ""
	if id == "" {
		fmt.Println("yes-----")
	} else {
		fmt.Println("no-----------")
	}
	fmt.Println(time.Now().AddDate(0, 1, 0), id)
}
