package commands

import "fmt"

type HelpResources struct {
	ProvideAbouts func () []aboutCommand
}

type helpCommand struct {
	res HelpResources
}

func NewHelpCommand(res HelpResources) *helpCommand {
	return &helpCommand{
		res: res,
	}
}

func (c *helpCommand) Execute() error {
	fmt.Println("\nThis is your Pokedex")
	fmt.Println("\nUsage:")
	for _, about := range c.res.ProvideAbouts() {
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