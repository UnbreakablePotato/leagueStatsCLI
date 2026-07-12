package main

import (
	"errors"
	"fmt"
	"os"
)

type command struct {
	name        string
	description string
	callback    func() error
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
