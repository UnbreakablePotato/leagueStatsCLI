package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/UnbreakablePotato/leagueStatsCLI/internal/cache"
	leagueapi "github.com/UnbreakablePotato/leagueStatsCLI/internal/leagueAPI"
	"github.com/joho/godotenv"
)

type command struct {
	name        string
	description string
	callback    func() error
	callbackS   func(region string, gamename string, tag string) error
}

var commandMap map[string]command

func commandExit() error {
	fmt.Println("Closing leagueStatsCLI\nGoodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Usage:")
	fmt.Print("\n\n")
	if len(commandMap) < 1 {
		return errors.New("No commands in the current iteration of the program")
	}

	for k, v := range commandMap {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

/*
example input for searching for profile

search region gamename tag
*/

var puuidCache = cache.NewCache(60000000000)

func commandSearch(region string, gamename string, tag string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	leagueapi.SearchPuuid(region, gamename, tag)

	entry, ok := puuidCache.Get(leagueapi.Usr.Puuid)

	if !ok {
		fmt.Println("Cache miss in commandSearch")
		apiKey, check := os.LookupEnv("leagueAPI")
		if !check {
			fmt.Println("Cannot find apikey")
		}

		fullurl := "https://euw1.api.riotgames.com/lol/league/v4/entries/by-puuid/" + leagueapi.Usr.Puuid + "?api_key=" + apiKey

		res, err := http.Get(fullurl)
		if err != nil {
			fmt.Printf("Search request failed: %v\n", err)
			return err
		}

		if res.StatusCode != http.StatusOK {
			fmt.Printf("Status code is not OK in commandSearch function: %v\n", res.StatusCode)
			return errors.New("Status code is not OK")
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("[]byte translation failed: %v\n", err)
			return err
		}

		if err := json.Unmarshal(data, &leagueapi.ShallowStats); err != nil {
			fmt.Printf("Unmarshal failed: %v\n", err)
			return err
		}

		puuidCache.Add(leagueapi.Usr.Puuid, data)
	} else {
		fmt.Println("Cache hit in commandSearch")
		if err := json.Unmarshal(entry.Val, &leagueapi.ShallowStats); err != nil {
			fmt.Printf("Unmarshal failed: %v\n", err)
			return err
		}
	}

	winRate := 0.0

	wins := leagueapi.ShallowStats[0].Wins
	losses := leagueapi.ShallowStats[0].Losses

	fmt.Printf("Showing ranked statistics for: %s\n", leagueapi.Usr.GameName)
	fmt.Printf("  - Rank: %s %s\n", leagueapi.ShallowStats[0].Tier, leagueapi.ShallowStats[0].Rank)
	fmt.Printf("  - LP: %d\n", leagueapi.ShallowStats[0].LeaguePoints)
	fmt.Printf("  - Wins: %d\n", leagueapi.ShallowStats[0].Wins)
	fmt.Printf("  - Losses: %d\n", leagueapi.ShallowStats[0].Losses)
	if wins > 0 || losses > 0 {
		winRate = (float64(wins) / (float64(wins) + float64(losses))) * 100.0
		fmt.Printf("  - Winrate: %d%%\n", int(winRate))
	}

	matchIds := leagueapi.SearcMatchID(leagueapi.Usr.Region, leagueapi.Usr.Puuid)

	if err := leagueapi.ShowShallowMatch(matchIds); err != nil {
		fmt.Printf("Could not show matches: %v\n", err)
	}

	return nil
}
