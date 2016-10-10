package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	var b bson.ObjectId
	b = ""
	fmt.Println(string(b) == "")

	// s := bson.NewObjectId().Hex()
	// bson.o
	// fmt.Println(bson.ObjectIdHex(""))

	// d := data.User{}
	// var t *data.Track
	// t = &d.Track
	// t.ID = bson.NewObjectId()
	//
	// fmt.Println(d)

	// m := bson.M{"name": 3}
	// m["id"] = "id"
	// fmt.Println(m)

	// err := errors.New("test")
	// err = apperr.Format(err, "ami", "hagu", "khai")
	// fmt.Println(err.Error())

	// session, err := r.Connect(r.ConnectOpts{
	// 	Address:  "localhost:28015",
	// 	Database: "padu",
	// })
	//
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// testUser := &data.User{}
	// testUser.ID = "test"
	//
	// err = db.Init(session)
	// fmt.Println(err)
	//
	// role := data.Role{}
	// role.RoleAccess = append(role.RoleAccess, "push")
	// dbRole := db.Role{}
	// dbRole.Session = session
	// err = dbRole.Put(&role, testUser)
	//
	// fmt.Println("role error", err)
	//
	// user := db.User{}
	// user.Session = session
	//
	// err = user.Put(&data.User{
	// 	Name: "nai",
	// 	Role: role,
	// }, testUser)
	// fmt.Println(err)

	// var test db.QueryFilter
	// test = &db.SkipStep{Skip: 1}
	//
	// fmt.Println(test, session)

	// dbUser := db.User{}
	// dbUser.Session = session
	// users, err := dbUser.List(&db.SkipStep{Skip: 1})
	// fmt.Println("test", len(users))
	// fmt.Println("err", err)

	// //

	// err = dbUser.Put(user)
	// //
	// fmt.Println(err)
	// fmt.Println(user.Role.RoleAccess)
}
