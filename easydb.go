package easydb

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DbInfo is a object stored info of database.
type DbInfo struct {
	DbName string
	URL    string
}

// New return a dbInfo object.
// URL is 127.0.0.1:27017 in default.
func New(dbName string, URL ...string) DbInfo {
	newDB := DbInfo{dbName, "127.0.0.1:27017"}
	if len(URL) != 0 {
		newDB.URL = URL[0]
	}

	return newDB
}

// Get get specific data from database
func (database DbInfo) Get(collName string, selector bson.M, dataSlicePtr interface{}) error {
	session, err := mgo.Dial(database.URL)
	if err != nil {
		log.Print(err)
		return err
	}
	defer session.Close()

	db := session.DB(database.DbName)
	c := db.C(collName)

	err = c.Find(selector).All(dataSlicePtr)
	if err == mgo.ErrNotFound {
		return nil
	} else if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// Remove remove specific data.
func (database DbInfo) Remove(collName string, selector bson.M) error {
	session, err := mgo.Dial(database.URL)
	if err != nil {
		log.Print(err)
		return err
	}
	defer session.Close()

	db := session.DB(database.DbName)
	c := db.C(collName)

	err = c.Remove(selector)
	if err == mgo.ErrNotFound {
		return nil
	} else if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// Update update a specific data.
func (database DbInfo) Update(collName string, selector, data bson.M) error {
	session, err := mgo.Dial(database.URL)
	if err != nil {
		log.Print(err)
		return err
	}
	defer session.Close()

	db := session.DB(database.DbName)
	c := db.C(collName)

	err = c.Update(selector, data)
	if err == mgo.ErrNotFound {
		return nil
	} else if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// Insert add a data to database.
func (database DbInfo) Insert(collName string, dataPtr interface{}) error {
	session, err := mgo.Dial(database.URL)
	if err != nil {
		log.Print(err)
		return err
	}
	defer session.Close()

	db := session.DB(database.DbName)
	c := db.C(collName)

	err = c.Insert(dataPtr)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
