package commands

import (
	"github.com/docech/pokedex-bootdev/cli/print"
	"github.com/docech/pokedex-bootdev/domain/pokedex"
)

type mapbCommand struct {
	prevLocationAreas pokedex.PreviousLocationAreasFunc
	getLocationAreas pokedex.GetLocationAreasFunc
}

func NewMapbCommand(prev pokedex.PreviousLocationAreasFunc, get pokedex.GetLocationAreasFunc) cliCommand {
	return &mapbCommand{
		prevLocationAreas: prev,
		getLocationAreas: get,
	}
}

func (c *mapbCommand) Execute() error {
	if err := c.prevLocationAreas(); err != nil {
		return err
	}

	print.PrintLocationAreas(c.getLocationAreas())
	return nil
}

func (c mapbCommand) About() aboutCommand {
	return aboutCommand{
		name:        "mapb",
		description: `Opposite of map command. Displays previous X location areas from Pokemon world.`,
	}
}