package main

//Player is a structure for the data
type Player struct {
	Username string `json:"name,omitempty"`
	Level    int    `json:"summonerLevel,omitempty"`
	ID       string `json:"id,omitempty"`
	Rank     []Rank
}

//Rank data
type Rank struct {
	Queue  string `json:"queueType,omitempty"`
	Tier   string `json:"tier,omitempty"`
	Wins   int    `json:"wins,omitempty"`
	Streak bool   `json:"hotStreak,omitempty"`
}
