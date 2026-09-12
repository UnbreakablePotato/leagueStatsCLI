package csmapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Champion struct {
	ChampionId              int
	ChampionIcon            string
	RecommendedPerks        []int
	RecommendedSpells       []int    // example [1, 2]
	RecommendedAbilityOrder []string // example ["Q","W","E"]
	RecommendedStartItems   []int    // should consist of a list of start item ids
	RecommendedItems        []int    // should consist of a list of item ids
	BestMatchups            []int    // should consist of a list of champions ids
	WorstMatchups           []int    // should consist of a list of champions ids
}

type Build struct {
	ID           int
	ChampionID   int
	ChampionName string
	Position     string

	Item0 int
	Item1 int
	Item2 int
	Item3 int
	Item4 int
	Item5 int
	Item6 int

	SummonerSpell1 int
	SummonerSpell2 int

	Keystone int
	Perk1    int
	Perk2    int
	Perk3    int
	Perk4    int
	Perk5    int
	Perk6    int

	Games int
	Wins  int
}

func BuildReq(url string, b *Build) (*Build, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error sending requet: %s\n", err)
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return b, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return b, err
	}

	if err := json.Unmarshal(data, &b); err != nil {
		fmt.Printf("Error: %s\n", err)
		return b, err
	}

	return b, nil
}
