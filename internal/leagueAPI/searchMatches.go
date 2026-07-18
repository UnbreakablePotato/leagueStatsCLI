package leagueapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var Champions = make(map[int64]string)

var Items = make(map[int]string)

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
	Info struct {
		GameDuration int `json:"gameDuration"`
		Participants struct {
			ChampLevel         int    `json:"champLevel"`
			ChampionName       string `json:"championName"`
			Deaths             int    `json:"deaths"`
			Kills              int    `json:"kills"`
			GoldEarned         int64  `json:"goldEarned"`
			IndividualPosition string `json:"individualPosition"`
			Item0              int    `json:"item0"`
			Item1              int    `json:"item1"`
			Item2              int    `json:"item2"`
			Item3              int    `json:"item3"`
			Item4              int    `json:"item4"`
			Item5              int    `json:"item5"`
			Item6              int    `json:"item6"`
		} `json:"participants"`
	} `json:"info"`
}

type CurrentGameInfo struct {
	GameId             int64  `json:"gameId"`
	GameType           string `json:"gameType"`
	GameStartTime      int64  `json:"gameStartTime"`
	PlatformId         string `json:"platformId"`
	GameQeueueConfigId int    `json:"gameQueueConfigId"`
	BannedChampions    []struct {
		PickTurn   int   `json:"pickTurn"`
		ChampionId int64 `json:"championId"`
		TeamId     int64 `json:"teamId"`
	} `json:"bannedChampions"`
	Participants []struct {
		Puuid      string `json:"puuid"`
		ChampionId int64  `json:"championId"`
		Spell1Id   int64  `json:"spell1Id"`
		Spell2Id   int64  `json:"spell2Id"`
	} `json:"participants"`
}

var _ = godotenv.Load()

var apiKey, _ = os.LookupEnv("leagueAPI")

func SearcMatchID(region string, puuid string) []string {

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

var LiveGameInfo CurrentGameInfo

// Avoid caching here. Whether or not the game is currently going could change any
// second, so caching the info is wasteful
func CheckCurrentGameInfo() bool {
	region := ""
	switch Usr.Region {
	case "euw1":
		region = "europe"
	}

	fullUrl := "https://" + region + ".api.riotgames.com/lol/spectator/v5/active-games/by-summoner/" + Usr.Puuid

	res, err := http.Get(fullUrl)
	if err != nil {
		fmt.Printf("Get request failed to show current game: %v\n", err)
		return false
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Status code is not OK: %d\n", res.StatusCode)
		return false
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Byte translation failed when attempting to show current game info: %v\n", err)
		return false
	}

	if err := json.Unmarshal(data, &LiveGameInfo); err != nil {
		fmt.Printf("Failed to unmarshal current game info: %v", err)
		return false
	}

	return true
}

//https://euw1.api.riotgames.com/lol/spectator/v5/active-games/by-summoner/
