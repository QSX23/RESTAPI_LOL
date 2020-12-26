package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//GetSummonerNameLevel gets our names and level
func GetSummonerNameLevel(name string) Player {
	url := "https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/" + name + API_KEY
	//fmt.Println(url)

	timeClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	errors(err)

	res, err := timeClient.Do(req)
	errors(err)

	user, err := ioutil.ReadAll(res.Body)
	errors(err)

	player := Player{}
	jsonError := json.Unmarshal(user, &player)
	errors(jsonError)

	//fmt.Println(player)

	res.Body.Close()

	return player
}

//GetTierWinsHot gets a few more fields of data
func GetTierWinsHot(p Player) Player {
	url := "https://na1.api.riotgames.com/lol/league/v4/entries/by-summoner/" + p.ID + API_KEY
	//fmt.Println(url)

	timeClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	errors(err)

	res, err := timeClient.Do(req)
	errors(err)

	user, err := ioutil.ReadAll(res.Body)
	errors(err)

	jsonError := json.Unmarshal(user, &p.Rank)
	errors(jsonError)

	//fmt.Println(p)

	return p

}

//PopulateData populates our table
func PopulateData(people []string) []Player {

	db := make([]Player, len(people))

	//iterate over the people and create a struct for each with all data
	for _, x := range people {
		p := GetTierWinsHot(GetSummonerNameLevel(x))
		db = append(db, p)
	}

	return db
}

//handles errors
func errors(err error) {
	if err != nil {
		fmt.Print(err)
	}
}
