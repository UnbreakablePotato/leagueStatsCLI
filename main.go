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
		"import": {
			name:         "import",
			description:  "Imports a runepage",
			callBackRune: commandImportRunePage,
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
		case 13:
			v, ok := commandMap[input[0]]
			if !ok {
				fmt.Println("Command does not exist..")
				continue
			}
			if input[0] == v.name {
				commandMap[v.name].callBackRune(input[1], input[2], input[3], input[4], input[5], input[6], input[7], input[8], input[9], input[10], input[11], input[12])
			} else {
				fmt.Println("Command does not exist..")
			}
		}

	}

}
