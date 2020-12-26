package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPlayersEndpoint(t *testing.T) {

	addData()

	req, err := http.NewRequest("GET", "/players", nil)
	fail(err, "failure", t)

	res := httptest.NewRecorder()

	Router().ServeHTTP(res, req)

	checkResponseCode(t, 200, res.Code)

	expectedResponse := `[{"name":"QKsix23","summonerLevel":298,"id":"-LplhonA_Y6liXp9Jb5q_nUUyAKOOvZf2bRUSkkJB2eNL98","Rank":[{"queueType":"RANKED_FLEX_SR","tier":"BRONZE","wins":9}]},{"name":"TapCity","summonerLevel":143,"id":"3RATWCstwnQX4ChTz3ChV2Eh4LqS_8i3jyQklEXfqA9k6N4","Rank":[{"queueType":"RANKED_SOLO_5x5","tier":"SILVER","wins":96},{"queueType":"RANKED_FLEX_SR","tier":"BRONZE","wins":12}]},{"name":"Pyromantics","summonerLevel":119,"id":"OHS-svee9-4HUF6XN2S-5-mfqwAiepFqU-HI-pmi6jxqhcI","Rank":[{"queueType":"RANKED_SOLO_5x5","tier":"SILVER","wins":46}]},{"name":"Raiders0002","summonerLevel":97,"id":"WAdCP3LahXU9DWsXDfvBRnsnX-sV-_y_1qf6WGxOr5o4HBs","Rank":[{"queueType":"RANKED_FLEX_SR","tier":"BRONZE","wins":5}]},{"name":"j4k71","summonerLevel":199,"id":"NQec0pjvcDcECJgup9E2m_w4NCfCAUKsdOYKAJPDQNVPNXA","Rank":[{"queueType":"RANKED_SOLO_5x5","tier":"SILVER","wins":96},{"queueType":"RANKED_FLEX_SR","tier":"BRONZE","wins":12}]}]` + "\n"

	x := res.Body.String()

	assert.Equal(t, expectedResponse, x)

	if x != expectedResponse {
		fmt.Println(x)
		t.Error("The response does not match expected")
	}

}

func TestGetPlayerEndpoint(t *testing.T) {

	//Make a request and look for response
	req, err := http.NewRequest("GET", "/players/qksix23", nil)
	fail(err, "failure", t)
	res := httptest.NewRecorder()

	//Call the router and sevrve response and request
	Router().ServeHTTP(res, req)

	checkResponseCode(t, 200, res.Code)
	x := res.Body.String()

	expectedResponse := `{"name":"QKsix23","summonerLevel":298,"id":"-LplhonA_Y6liXp9Jb5q_nUUyAKOOvZf2bRUSkkJB2eNL98","Rank":[{"queueType":"RANKED_FLEX_SR","tier":"BRONZE","wins":9}]}` + "\n"

	assert.Equal(t, expectedResponse, x)

}

func TestCreatePlayerEndpoint(t *testing.T) {
	//Make a request and look for response
	req, err := http.NewRequest("POST", "/players/Boshi", nil)
	fail(err, "failure", t)
	res := httptest.NewRecorder()

	//Call the router and sevrve response and request
	Router().ServeHTTP(res, req)

	checkResponseCode(t, 200, res.Code)
	x := res.Body.String()
	expectedResponse := `[{"name":"QKsix23","summonerLevel":298,"id":"-LplhonA_Y6liXp9Jb5q_nUUyAKOOvZf2bRUSkkJB2eNL98","Rank":[{"queueType":"RANKED_FLEX_SR","tier":"BRONZE","wins":9}]},{"name":"TapCity","summonerLevel":143,"id":"3RATWCstwnQX4ChTz3ChV2Eh4LqS_8i3jyQklEXfqA9k6N4","Rank":[{"queueType":"RANKED_SOLO_5x5","tier":"SILVER","wins":96},{"queueType":"RANKED_FLEX_SR","tier":"BRONZE","wins":12}]},{"name":"Pyromantics","summonerLevel":119,"id":"OHS-svee9-4HUF6XN2S-5-mfqwAiepFqU-HI-pmi6jxqhcI","Rank":[{"queueType":"RANKED_SOLO_5x5","tier":"SILVER","wins":46}]},{"name":"Raiders0002","summonerLevel":97,"id":"WAdCP3LahXU9DWsXDfvBRnsnX-sV-_y_1qf6WGxOr5o4HBs","Rank":[{"queueType":"RANKED_FLEX_SR","tier":"BRONZE","wins":5}]},{"name":"j4k71","summonerLevel":199,"id":"NQec0pjvcDcECJgup9E2m_w4NCfCAUKsdOYKAJPDQNVPNXA","Rank":[{"queueType":"RANKED_SOLO_5x5","tier":"SILVER","wins":96},{"queueType":"RANKED_FLEX_SR","tier":"BRONZE","wins":12}]},{"name":"Boshi","summonerLevel":198,"id":"GyOpp0qE-5oGnnBH4yboTZGftYIo2UQQ-HWMteHniA","Rank":[{"queueType":"RANKED_SOLO_5x5","tier":"PLATINUM","wins":12},{"queueType":"RANKED_FLEX_SR","tier":"SILVER","wins":21}]}]` + "\n"
	assert.Equal(t, expectedResponse, x)

}

func TestDeletePlayerEndpoint(t *testing.T) {
	//Make a request and look for response
	req, err := http.NewRequest("DELETE", "/players/Boshi", nil)
	fail(err, "failure", t)
	res := httptest.NewRecorder()

	//Call the router and sevrve response and request
	Router().ServeHTTP(res, req)

	checkResponseCode(t, 200, res.Code)

	expectedResponse := ``

	x := res.Body.String()

	assert.Equal(t, expectedResponse, x)

}

//handles errors
func fail(err error, s string, t *testing.T) {
	if err != nil {
		t.Fatal()
	}
}

func addData() {
	data := PopulateData(People)

	for _, item := range data {
		DB = append(DB, item)
	}
}

//checks response codes
func checkResponseCode(t *testing.T, expect, response int) {
	fmt.Printf("Response %d, Expected %d\n", response, expect)
	if expect != response {
		t.Errorf("Expected %d. Got %d\n", expect, response)
	}
}
