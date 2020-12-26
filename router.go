package main

import "github.com/gorilla/mux"

//Router is the router for the application
func Router() *mux.Router {
	router := mux.NewRouter()

	//all the endpoints
	router.HandleFunc("/players", GetPlayersEndpoint).Methods("GET")
	router.HandleFunc("/players/{name}", GetPlayerEndpoint).Methods("GET")
	//router.HandleFunc("/players/{name}", CreateWholePlayerEndpoint).Methods("POST")
	router.HandleFunc("/players/{name}", CreatePlayerEndpoint).Methods("POST")
	//router.HandleFunc("/players/{name}", UpdatePlayerEndpoint).Methods("PUT")
	router.HandleFunc("/players/{name}", DeletePlayerEndpoint).Methods("DELETE")
	return router
}
