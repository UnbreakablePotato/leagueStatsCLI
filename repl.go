package main

import "strings"

func cleanInput(input string) []string {
	res := strings.Split(input, " ")

	clean := []string{}

	for i := range res {
		inter := strings.ReplaceAll(res[i], " ", "")
		clean = append(clean, inter)
	}

	return clean
}
