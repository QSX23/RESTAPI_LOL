package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//People is the players to be populated into the database
var People = []string{"qksix23", "TapCity", "Pyromantics", "Raiders0002", "j4k71"}

//DB is the global variable for the database
var DB []Player

//GetPlayersEndpoint an enpoint with all players in above string listed
func GetPlayersEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(DB)
}

//GetPlayerEndpoint an endpoint for a specific player
func GetPlayerEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, item := range DB {
		if strings.ToLower(item.Username) == strings.ToLower(params["name"]) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Player{})
}

//CreateWholePlayerEndpoint creates a player from a whole entry
func CreateWholePlayerEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	var player Player

	_ = json.NewDecoder(req.Body).Decode(&player)
	player.Username = params["name"]
	DB = append(DB, player)
	json.NewEncoder(w).Encode(DB)
}

//CreatePlayerEndpoint creates a player from just a username
func CreatePlayerEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	//var player Player

	//_ = json.NewDecoder(req.Body).Decode(&player)
	p := GetTierWinsHot(GetSummonerNameLevel(params["name"]))
	DB = append(DB, p)
	json.NewEncoder(w).Encode(DB)
}

//UpdatePlayerEndpoint updates a player record. There is no need for this so the endpoint is commented out
func UpdatePlayerEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	var player Player

	for _, item := range DB {
		if strings.ToLower(item.Username) == strings.ToLower(params["name"]) {
			_ = json.NewDecoder(req.Body).Decode(&player)
			player.Username = params["name"]
			DB = append(DB, player)
			json.NewEncoder(w).Encode(DB)
			return
		}
	}

	json.NewEncoder(w).Encode(&Player{})

}

//DeletePlayerEndpoint deletes the named player
func DeletePlayerEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for i, item := range DB {
		if strings.ToLower(item.Username) == strings.ToLower(params["name"]) {
			DB = append(DB[:i], DB[i+1:]...)
			break
		}
	}

}

//the main program
func main() {

	data := PopulateData(People)

	for _, item := range data {
		DB = append(DB, item)
	}

	//router := mux.NewRouter()
	router := Router()

	fmt.Println("API running")

	//listener for requests wrapped in log fatal to document failure
	log.Fatal(http.ListenAndServe(":8000", router))

}
