package leagueapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/UnbreakablePotato/leagueStatsCLI/internal/cache"
	"github.com/joho/godotenv"
)

type user struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type entries struct {
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	Puuid        string `json:"puuid"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Veteran      bool   `json:"veteran"`
	FreshBlood   bool   `json:"freshBlood"`
	HotStreak    bool   `json:"hotStreak"`
}

var Usr user

var ShallowStats []entries

var UsrCache = cache.NewCache(60000000000)

func SearchPuuid(region string, gamename string, tagline string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	switch region {
	case "euw":
		region = "europe"
	}

	fullId := gamename + tagline

	entry, ok := UsrCache.Get(fullId)

	if !ok {
		fmt.Printf("%s not in  cache\n", fullId)
		safeGameName := url.PathEscape(gamename)
		safeTagLine := url.PathEscape(tagline)

		apiKey, check := os.LookupEnv("leagueAPI")
		if !check {
			fmt.Println("Cannot find apikey")
		}
		fullurl := "https://" + region + ".api.riotgames.com/riot/account/v1/accounts/by-riot-id/" + safeGameName + "/" + safeTagLine + "?api_key=" + apiKey
		//fmt.Printf("%s\n", fullurl)

		res, err := http.Get(fullurl)
		if err != nil {
			fmt.Printf("Search request failed: %v\n", err)
			return err
		}

		if res.StatusCode != http.StatusOK {
			fmt.Printf("Status code is not OK: %d\n", res.StatusCode)
			return errors.New("Status code is not OK")
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("[]byte translation failed: %v\n", err)
			return err
		}

		if err := json.Unmarshal(data, &Usr); err != nil {
			fmt.Printf("Unmarshal failed: %v\n", err)
			return err
		}
		UsrCache.Add(fullId, data)

		return nil

	}

	if err := json.Unmarshal(entry.Val, &Usr); err != nil {
		fmt.Printf("Unmarshal failed from cache: %v\n", err)
		return err
	}
	fmt.Println("Success: Cache access in searchUser")

	//fmt.Printf("Debug: %s\n", Usr.GameName)
	//fmt.Printf("Debug: %s\n", apiKey)

	return nil
}

/*
func SearchPuuid(region string, gamename string, tagline string) error {

	if region == "euw" {
		region = "europe"
	}

	safeGameName := url.PathEscape(gamename)
	safeTagLine := url.PathEscape(tagline)

	apiKey := os.Getenv("leagueAPI")
	fullurl := "https://" + region + ".api.riotgames.com/riot/account/v1/accounts/by-riot-id/" + safeGameName + "/" + safeTagLine
	fmt.Printf("%s\n", fullurl)
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		fmt.Printf("Search request failed: %v\n", err)
		return err
	}

	req.Header.Set("X-Riot-Token", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error during Do method: %v", err)
		return err
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Status code is not OK: %d\n", res.StatusCode)
		return errors.New("Status code is not OK")
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("[]byte translation failed: %v\n", err)
		return err
	}

	if err := json.Unmarshal(data, &Usr); err != nil {
		fmt.Printf("Unmarshal failed: %v\n", err)
		return err
	}

	fmt.Printf("Debug: %s\n", Usr.GameName)
	fmt.Printf("Debug: %s\n", apiKey)

	return nil
}
*/
