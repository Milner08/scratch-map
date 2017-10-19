package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/milner08/scratch-map/API/models"
	mgo "gopkg.in/mgo.v2"
)

type MapDataController struct {
	session *mgo.Session
}

// NewMapDataController provides a reference to a MapDataController with provided mongo session
func NewMapDataController(s *mgo.Session) *MapDataController {
	return &MapDataController{s}
}

//GetMapData gets the Map Data for the specified ID.
func (mdc MapDataController) GetMapData(w http.ResponseWriter, r *http.Request) {
	//Some kind of user check via RPC
	params := mux.Vars(r)

	mapData := models.MapData{}

	if err := mapData.GetMapData(mdc.session, params["id"]); err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(mapData)
}

func (mdc MapDataController) UpdateMapData(w http.ResponseWriter, r *http.Request) {
	//Some kind of user check via RPC
	mapData := models.MapData{}

	_ = json.NewDecoder(r.Body).Decode(&mapData)

	mapData.UpdateMapData(mdc.session)

	json.NewEncoder(w).Encode(mapData)
}

func (mdc MapDataController) CreateMapData(w http.ResponseWriter, r *http.Request) {
	//Some kind of user check via RPC
	mapData := models.MapData{}

	_ = json.NewDecoder(r.Body).Decode(&mapData)

	mapData.InsertNewMapData(mdc.session)

	json.NewEncoder(w).Encode(mapData)
}
