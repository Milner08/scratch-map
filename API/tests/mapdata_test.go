package tests

import (
	"testing"

	"github.com/milner08/scratch-map/API/controllers"
	mgo "gopkg.in/mgo.v2"
)

func TestGetMapData(t *testing.T) {
	mdc := controllers.NewMapDataController(SetupMongoDBSession())
	mdc.GetMapData(w, r)
}

func SetupMongoDBSession() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return session
}
