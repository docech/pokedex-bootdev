package commands

import "os"

type exitCommand struct {
}

func NewExitCommand() exitCommand {
	return exitCommand{}
}

func (c exitCommand) Execute() error {
	os.Exit(0)
	return nil
}

func (c exitCommand) About() aboutCommand {
	return aboutCommand{
		name:        "exit",
		description: "Exit the Pokedex",
	}
}