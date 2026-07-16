package leagueapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type ShallowMatch struct {
	Info struct {
		GameDuration int `json:"gameDuration"`
		Participants struct {
			ChampLevel   int    `json:"champLevel"`
			ChampionName string `json:"championName"`
		} `json:"participants"`
	} `json:"info"`
}

type DeepMatch struct {
}

func SearcMatchID(region string, puuid string) []string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey, check := os.LookupEnv("leagueAPI")
	if !check {
		fmt.Println("Cannot find apikey")
	}

	switch region {
	case "euw":
		region = "europe"
	}

	fullUrl := "https://" + region + ".api.riotgames.com/lol/match/v5/matches/by-puuid/" + puuid + "/ids?start=0&count=5&api_key=" + apiKey

	res, err := http.Get(fullUrl)
	if err != nil {
		fmt.Printf("matchId request failed: %v\n", err)
		return []string{}
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	interResult := string(data)

	result := strings.Split(interResult, ",")

	for i := range result {
		fmt.Printf("debug: %s", result[i])
	}

	return result
}

/*
I wanna be able to show a list of matches with player names and champion names
and perhaps summonors for each

then a seperate command where you can look at just one match in far more detail
*/

func ShowShallowMatch(matchIds []string) {

}

func ShowDeepMatch(matchId string) {

}
