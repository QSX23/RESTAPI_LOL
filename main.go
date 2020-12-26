package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var people = []string{"qksix23", "TapCity", "Pyromantics", "Raiders0002", "j4k71"}

func main() {
	router := mux.NewRouter()

	db := PopulateData(people)

	fmt.Println(db)

	/*
		//all of the endpoints for this REST api
		router.HandleFunc("/players", GetPlayersEndpoint).Methods("GET")
		router.HandleFunc("/players/{id}", GetPlayerEndpoint).Methods("GET")
		router.HandleFunc("/players/{id}", CreatePlayerEndpoint).Methods("POST")
		router.HandleFunc("/players/{id}", UpdatePlayerEndpoint).Methods("PUT")
		router.HandleFunc("/players/{id}", DeletePlayerEndpoint).Methodss("DELETE")
	*/

	//listener for requests wrapped in log fatal to document failure
	log.Fatal(http.ListenAndServe(":8000", router))

}
