package easydb

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string `bson:"name"`
	Sex  string `bson:"sex"`
}

func TestDb(t *testing.T) {
	database := New("testDB")

	Convey("Given a db", t, func() {

		Convey("Insert a data", func() {
			one := Person{"Tony", "male"}
			err := database.Insert("dbtestColl", &one)
			So(err, ShouldBeNil)
		})

		Convey("get a data 1", func() {
			var data []Person
			err := database.Get("dbtestColl", bson.M{"name": "Tony"}, &data)
			So(err, ShouldBeNil)
			So(data[0].Sex, ShouldEqual, "male")
		})

		Convey("update a data", func() {
			selector := bson.M{"name": "Tony"}
			data := bson.M{"$set": bson.M{"sex": "female"}}
			err := database.Update("dbtestColl", selector, data)
			So(err, ShouldBeNil)
		})

		Convey("get a data 2", func() {
			var data []Person
			err := database.Get("dbtestColl", bson.M{"name": "Tony"}, &data)
			So(err, ShouldBeNil)
			So(data[0].Sex, ShouldEqual, "female")
		})

		Convey("remove a data", func() {
			selector := bson.M{"name": "Tony"}
			err := database.Remove("dbtestColl", selector)
			So(err, ShouldBeNil)
		})

		Convey("get a data 3", func() {
			var data []Person
			err := database.Get("dbtestColl", bson.M{"name": "Tony"}, &data)
			So(err, ShouldBeNil)
			So(len(data), ShouldEqual, 0)
		})
	})
}
