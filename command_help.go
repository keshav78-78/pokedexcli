package main

import (
	"fmt"
)

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the pokedex help menu!")
	fmt.Println("Here are your avalable commands")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s \n", cmd.name, cmd.description)
	}

	fmt.Println("")
	return nil
}
