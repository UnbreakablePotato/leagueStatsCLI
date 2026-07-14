package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	commandMap = map[string]command{
		"exit": {
			name:        "exit",
			description: "Closes the CLI program safely",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Lists every command available in the program",
			callback:    commandHelp,
		},
		"search": {
			name:        "search",
			description: "Searches a user and shows their overall soloq statistics",
			callbackS:   commandSearch,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("LS > ")
		input := []string{}

		isText := scanner.Scan()

		if !isText {
			fmt.Println("How did you even do that?")
		}

		text := scanner.Text()

		input = append(input, cleanInput(text)...)

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input", err)
			fmt.Fprintln(os.Stderr, err)
		}

		switch len(input) {
		case 1:
			v, ok := commandMap[input[0]]
			if !ok {
				fmt.Println("Command does not exist..")
				continue
			}
			if input[0] == v.name {
				commandMap[v.name].callback()
			} else {
				fmt.Println("Command does not exist..")
			}
		case 4:
			v, ok := commandMap[input[0]]
			if !ok {
				fmt.Println("Command does not exist..")
				continue
			}
			if input[0] == v.name {
				commandMap[v.name].callbackS(input[1], input[2], input[3])
			} else {
				fmt.Println("Command does not exist..")
			}
		}

	}

}
