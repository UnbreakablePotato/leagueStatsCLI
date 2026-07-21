package leagueapi

import (
	"encoding/json"
	"errors"
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
		Participants []struct {
			ChampLevel   int    `json:"champLevel"`
			ChampionName string `json:"championName"`
			Deaths       int    `json:"deaths"`
			Kills        int    `json:"kills"`
			Assists      int    `json:"assists"`
		} `json:"participants"`
	} `json:"info"`
}

type DeepMatch struct {
	Info struct {
		GameDuration int `json:"gameDuration"`
		Participants []struct {
			ChampLevel         int    `json:"champLevel"`
			ChampionName       string `json:"championName"`
			Deaths             int    `json:"deaths"`
			Kills              int    `json:"kills"`
			Assists            int    `json:"assists"`
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

	/*switch region {
	case "euw":
		region = "europe"
	}*/

	fullUrl := "https://" + region + ".api.riotgames.com/lol/match/v5/matches/by-puuid/" + puuid + "/ids?start=0&count=5&api_key=" + apiKey
	//fmt.Printf("%s\n", fullUrl)
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
		result[i] = strings.ReplaceAll(result[i], "\"", "")
		result[i] = strings.ReplaceAll(result[i], "[", "")
		result[i] = strings.ReplaceAll(result[i], "]", "")
		//fmt.Printf("debug: %s\n", result[i])
	}

	return result
}

/*
I wanna be able to show a list of matches with player names and champion names
and perhaps summonors for each

then a seperate command where you can look at just one match in far more detail
*/

var shallow ShallowMatch

func ShowShallowMatch(matchIds []string) error {

	for i := range matchIds {
		id := matchIds[i]

		fullUrl := "https://" + Usr.Region + ".api.riotgames.com/lol/match/v5/matches/" + id + "?api_key=" + apiKey
		//https://europe.api.riotgames.com/lol/match/v5/matches/EUW1_7922609631?api_key=
		//fmt.Printf("%s\n", fullUrl)
		res, err := http.Get(fullUrl)
		if err != nil {
			fmt.Printf("Get rquest failed when requesting match id: %s, error: %v", id, err)
			return errors.New("")

		}

		if res.StatusCode != http.StatusOK {
			fmt.Printf("Status code is not OK: %d\n", res.StatusCode)
			return errors.New("")
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("io.ReadAll failed in ShowShallowFunction: %v", err)
			return errors.New("")
		}

		if err := json.Unmarshal(data, &shallow); err != nil {
			fmt.Printf("Unmarshal failed in ShowShallowMatch: %v", err)
			return errors.New("")
		}

		fmt.Println("-----------------------------------------------")
		fmt.Print("-----------------------------------------------\n\n")

		fmt.Printf("MATCHID: %s\n\n", id)

		fmt.Printf("%s ---------------------------------------- %s\n", shallow.Info.Participants[0].ChampionName, shallow.Info.Participants[5].ChampionName)

		fmt.Printf("K: %d D: %d A: %d ---- K: %d D: %d A: %d\n\n", shallow.Info.Participants[0].Kills, shallow.Info.Participants[0].Deaths, shallow.Info.Participants[0].Assists,
			shallow.Info.Participants[5].Kills, shallow.Info.Participants[5].Deaths, shallow.Info.Participants[5].Assists)

		fmt.Printf("%s ---------------------------------------- %s\n", shallow.Info.Participants[1].ChampionName, shallow.Info.Participants[6].ChampionName)

		fmt.Printf("K: %d D: %d A: %d ---- K: %d D: %d A: %d\n\n", shallow.Info.Participants[1].Kills, shallow.Info.Participants[1].Deaths, shallow.Info.Participants[1].Assists,
			shallow.Info.Participants[6].Kills, shallow.Info.Participants[6].Deaths, shallow.Info.Participants[6].Assists)

		fmt.Printf("%s ---------------------------------------- %s\n", shallow.Info.Participants[2].ChampionName, shallow.Info.Participants[7].ChampionName)

		fmt.Printf("K: %d D: %d A: %d ---- K: %d D: %d A: %d\n\n", shallow.Info.Participants[2].Kills, shallow.Info.Participants[2].Deaths, shallow.Info.Participants[2].Assists,
			shallow.Info.Participants[7].Kills, shallow.Info.Participants[7].Deaths, shallow.Info.Participants[7].Assists)

		fmt.Printf("%s ---------------------------------------- %s\n", shallow.Info.Participants[3].ChampionName, shallow.Info.Participants[8].ChampionName)

		fmt.Printf("K: %d D: %d A: %d ---- K: %d D: %d A: %d\n\n", shallow.Info.Participants[3].Kills, shallow.Info.Participants[3].Deaths, shallow.Info.Participants[3].Assists,
			shallow.Info.Participants[8].Kills, shallow.Info.Participants[8].Deaths, shallow.Info.Participants[8].Assists)

		fmt.Printf("%s ---------------------------------------- %s\n", shallow.Info.Participants[4].ChampionName, shallow.Info.Participants[9].ChampionName)

		fmt.Printf("K: %d D: %d A: %d ---- K: %d D: %d A: %d\n", shallow.Info.Participants[4].Kills, shallow.Info.Participants[4].Deaths, shallow.Info.Participants[4].Assists,
			shallow.Info.Participants[9].Kills, shallow.Info.Participants[9].Deaths, shallow.Info.Participants[9].Assists)

		fmt.Println("-----------------------------------------------")
		fmt.Print("-----------------------------------------------\n\n")

	}

	return nil
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
	//fmt.Printf("%s\n", fullUrl)
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
