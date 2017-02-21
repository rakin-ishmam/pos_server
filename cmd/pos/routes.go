package main

import (
	"github.com/rakin-ishmam/pos_server/api"
	"gopkg.in/mgo.v2"
)

func allRoutes(session *mgo.Session) []Route {
	rs := []Route{}

	loginRoutes(&rs, session)
	categoryRoutes(rs, session)
	customerRoutes(rs, session)
	inventoryRoutes(rs, session)
	orderPaymentRoutes(rs, session)
	productRoutes(rs, session)
	roleRoutes(rs, session)
	sellRoutes(rs, session)
	userRoutes(&rs, session)

	return rs
}

func loginRoutes(rs *[]Route, session *mgo.Session) {
	*rs = append(*rs, Route{
		Name:    "User login",
		Method:  "POST",
		Path:    "/api/login",
		Handler: panicRecover(JSONRunner(api.Login, session)),
	})
}

func categoryRoutes(rs []Route, session *mgo.Session) {

}

func customerRoutes(rs []Route, session *mgo.Session) {

}

func inventoryRoutes(rs []Route, session *mgo.Session) {

}

func orderPaymentRoutes(rs []Route, session *mgo.Session) {

}

func productRoutes(rs []Route, session *mgo.Session) {

}

func roleRoutes(rs []Route, session *mgo.Session) {

}

func sellRoutes(rs []Route, session *mgo.Session) {

}

func userRoutes(rs *[]Route, session *mgo.Session) {
	//get list of the user
	*rs = append(*rs, Route{
		Name:    "List of user",
		Method:  "GET",
		Path:    "/api/user/list",
		Handler: panicRecover(requiredToken(JSONRunner(api.ListUser, session), session)),
	})

	//get one user by id
	*rs = append(*rs, Route{
		Name:    "Fetch user by id",
		Method:  "GET",
		Path:    "/api/user/{id}",
		Handler: panicRecover(requiredToken(JSONRunner(api.FetchUser, session), session)),
	})

	// update one user
	*rs = append(*rs, Route{
		Name:    "Update User",
		Method:  "PUT",
		Path:    "/api/user/{id}",
		Handler: panicRecover(requiredToken(JSONRunner(api.UpdateUser, session), session)),
	})

	// create one user
	*rs = append(*rs, Route{
		Name:    "Create User",
		Method:  "POST",
		Path:    "/api/user",
		Handler: panicRecover(requiredToken(JSONRunner(api.CreateUser, session), session)),
	})

}
