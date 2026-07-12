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

		if len(input) > 1 {
			fmt.Println("No current command supports more than one arguments...")
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
		}

	}

}
