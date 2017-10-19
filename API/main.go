package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/milner08/scratch-map/API/controllers"
	"gopkg.in/mgo.v2"
)

// our main function
func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	mdc := controllers.NewMapDataController(session)

	router := mux.NewRouter()
	router.HandleFunc("/mapdata/{id}", mdc.GetMapData).Methods("GET")
	router.HandleFunc("/mapdata", mdc.UpdateMapData).Methods("PUT")
	router.HandleFunc("/mapdata", mdc.CreateMapData).Methods("POST")

	//Setup JWT, Logging and CORS
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	corsRouter := handlers.CORS(originsOk, headersOk, methodsOk)(loggedRouter)

	log.Fatal(http.ListenAndServe(":8000", corsRouter))

}
