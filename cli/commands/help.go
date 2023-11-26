package commands

import "fmt"

type HelpDeps struct {
	ProvideAbouts func () []aboutCommand
}

type helpCommand struct {
	deps HelpDeps
}

func NewHelpCommand(deps HelpDeps) cliCommand {
	return &helpCommand{
		deps: deps,
	}
}

func (c *helpCommand) Execute() error {
	fmt.Println("\nThis is your Pokedex")
	fmt.Println("\nUsage:")
	for _, about := range c.deps.ProvideAbouts() {
		fmt.Printf("%s - %s\n", about.name, about.description)
	}
	fmt.Print("\n")
	return nil
}

func (c helpCommand) About() aboutCommand {
	return aboutCommand{
		name:        "help",
		description: "Displays a help message",
	}
}