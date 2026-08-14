package lcu

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
)

type SummonerSpell struct {
	Spell1ID int
	Spell2ID int
}

func SetSummonerSpells(spell1 int, spell2 int) error {
	fullUrl := "https://127.0.0.1:" + port + "lol/champ-select/v1/session/my-selection"

	req, err := http.NewRequest("PATCH", fullUrl, nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}

	req.SetBasicAuth("riot", password)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error: %d\n", res.StatusCode)
		return errors.New("Got non ok status code")
	}

	return nil
}
