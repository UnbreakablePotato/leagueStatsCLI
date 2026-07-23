package lcu

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Runepage struct {
	Name             string `json:"name"`
	Id               int    `json:"id"`
	PrimaryStyleId   int    `json:"primaryStyleId"`
	SubStyleId       int    `json:"subStyleId"`
	SelectedPerksIds []int  `json:"selectedPerkIds"`
	Current          bool   `json:"isActive"`
}

const lockPath = "C:\\Riot Games\\League of Legends\\lockfile"

var port string

var password string

func parseLockFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not open file: %s", err)
		return err
	}

	lockString := string(data)

	lockSlice := strings.Split(lockString, ":")

	port = lockSlice[2]

	password = lockSlice[3]

	return nil
}

var curr Runepage

var IsCurr bool

func GetRunePage() error {
	parseLockFile(lockPath)

	fullUrl := "https://127.0.0.1:" + port + "/lol-perks/v1/currentpage"

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		fmt.Printf("1: %s\n", err)
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
		fmt.Printf("2: %s\n", err)
		return err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Printf("Request failed with: %d\n", res.StatusCode)
		return errors.New("")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("3: %s\n", err)
		return err
	}

	if err := json.Unmarshal(data, &curr); err != nil {
		fmt.Printf("4: %s\n", err)
		return err
	}

	IsCurr = curr.Current

	return nil
}

func DeleteRunePage() error {

	err := GetRunePage()
	if err != nil {
		fmt.Printf("GetRunePage returned error: %s\n", err)
		return err
	}

	stringId := strconv.Itoa(curr.Id)

	fullUrl := "https://127.0.0.1:" + port + "/lol-perks/v1/pages" + stringId

	req, err := http.NewRequest("DELETE", fullUrl, nil)
	if err != nil {
		fmt.Printf("http.NewRequest failed: %s\n", err)
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
		fmt.Printf("Could not client.Do request in deleterunepage: %s\n", err)

		return err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("could not readall in deleterunepage: %s\n", err)
		return err
	}

	if err := json.Unmarshal(data, &curr); err != nil {
		fmt.Printf("could not unmarshal in deleterunape: %s\n", err)
		return err
	}

	/*fmt.Println("Current rune page has:")
	fmt.Printf("name: %s\n", curr.Name)
	fmt.Printf("id: %d\n", curr.Id)
	fmt.Printf("primary: %d\n", curr.PrimaryStyleId)
	fmt.Printf("sub: %d\n", curr.SubStyleId)
	fmt.Printf("Perk: %d\n", curr.SelectedPerksIds[0])
	fmt.Printf("Perk: %d\n", curr.SelectedPerksIds[1])
	fmt.Printf("Perk: %d\n", curr.SelectedPerksIds[2])
	fmt.Printf("Perk: %d\n", curr.SelectedPerksIds[3])
	fmt.Printf("Perk: %d\n", curr.SelectedPerksIds[4])
	fmt.Printf("Perk: %d\n", curr.SelectedPerksIds[5])
	fmt.Printf("Perk: %d\n", curr.SelectedPerksIds[6])
	fmt.Printf("Perk: %d\n", curr.SelectedPerksIds[7])
	fmt.Printf("Perk: %d\n", curr.SelectedPerksIds[8])
	fmt.Printf("Current: %v\n", curr.Current)*/

	return nil
}

func PostRunePage(page *Runepage) error {

	err := DeleteRunePage()
	if err != nil {
		fmt.Printf("DeleteRunePage returned an error: %s\n", err)
		return err
	}

	runejson, err := json.Marshal(page)
	if err != nil {
		return err
	}

	fullUrl := "https://127.0.0.1:" + port + "/lol-perks/v1/pages"

	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(runejson))
	if err != nil {
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
		return err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &curr); err != nil {
		return err
	}

	return nil
}
