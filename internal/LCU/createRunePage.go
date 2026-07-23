package lcu

import (
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
	SelectedPerksIds []int  `json:"selectedPerksIds"`
	Current          bool   `json:"current"`
}

const lockPath = "/mnt/c/Riot Games/League of Legends/lockfile"

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

func GetRunePage() error {
	parseLockFile(lockPath)

	fullUrl := "https://127.0.0.1:" + port + "lol-perks/v1/currentpage"

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Basic riot:"+password)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Printf("Request failed with: %d\n", res.StatusCode)
		return errors.New("")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &curr); err != nil {
		return err
	}

	return nil
}

func DeleteRunePage() error {

	err := GetRunePage()
	if err != nil {
		return err
	}

	stringId := strconv.Itoa(curr.Id)

	fullUrl := "https://127.0.0.1:" + port + "/lol-perks/v1/pages" + stringId

	req, err := http.NewRequest("DELETE", fullUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Basic riot:"+password)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

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

func PostRunePage(page *Runepage) error {

	err := DeleteRunePage()
	if err != nil {
		fmt.Printf("DeleteRunePage returned an error: %s", err)
		return err
	}

	fullUrl := "https://127.0.0.1:" + port + "/lol-perks/v1/pages"

	req, err := http.NewRequest("POST", fullUrl, nil)
	if err != nil {
		return err
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &curr); err != nil {
		return err
	}

	return nil
}
