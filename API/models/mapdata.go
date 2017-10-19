package models

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MapData type, used to export a list of visited countries to the client
type MapData struct {
	ID               bson.ObjectId `json:"id" bson:"_id"`
	VisitedCountries []string      `json:"visited_countries" bson:"visited_countries"`
}

//Get the map data associated with an ID. Return an error if it cant be found.
func (md *MapData) GetMapData(session *mgo.Session, id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("ID is not valid.")
	}

	oid := bson.ObjectIdHex(id)

	if err := session.DB("scratch_map").C("map_data").FindId(oid).One(&md); err != nil {
		return err
	}

	return nil
}

//Updates a map data object.
func (md *MapData) UpdateMapData(session *mgo.Session) {
	session.DB("scratch_map").C("map_data").UpdateId(md.ID, md)
}

//Creates a new map data object
func (md *MapData) InsertNewMapData(session *mgo.Session) {
	md.ID = bson.NewObjectId()

	session.DB("scratch_map").C("map_data").Insert(md)
}
