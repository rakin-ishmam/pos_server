package main

import (
	"gopkg.in/mgo.v2"
)

func allRoutes(session *mgo.Session) []Route {
	rs := []Route{}

	loginRoutes(rs, session)
	categoryRoutes(rs, session)
	customerRoutes(rs, session)
	inventoryRoutes(rs, session)
	orderPaymentRoutes(rs, session)
	productRoutes(rs, session)
	roleRoutes(rs, session)
	sellRoutes(rs, session)
	userRoutes(rs, session)

	return rs
}

func loginRoutes(rs []Route, session *mgo.Session) {

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

func userRoutes(rs []Route, session *mgo.Session) {

}
