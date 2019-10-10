package main

import (
	"fmt"

	mgo "github.com/globalsign/mgo"
	bson "github.com/globalsign/mgo/bson"
)

const (
	hosts      = "localhost:27017"
	database   = "testdb"
	username   = ""
	password   = ""
	collection = "persons"
	filename   = "writer/testfile.json"
)

// type person struct {
// 	Name    string   `json:"name"`
// 	Surname string   `json:"surname"`
// 	Age     int      `json:"age"`
// 	Friends []string `json:"friends"`
// }

func main() {
	session, err := mgo.Dial(hosts)
	check(err)

	c := session.DB(database).C(collection)

	var p []interface{}
	// var p []person

	var filter bson.D
	err = c.Find(filter).All(&p)
	for i := range p {
		fmt.Println(p[i])
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
