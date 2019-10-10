package main

import (
	"encoding/json"
	"io/ioutil"

	mgo "github.com/globalsign/mgo"
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

	file, err := ioutil.ReadFile(filename)
	check(err)

	var p interface{}
	err = json.Unmarshal(file, &p)
	check(err)

	// var x []interface{}
	// for i, el := range p {
	// 	x[i] = el
	// }

	err = c.Insert(p)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
