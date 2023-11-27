package commands

import (
	"github.com/docech/pokedex-bootdev/cli/print"
	"github.com/docech/pokedex-bootdev/domain/pokemonworld"
)

type mapbCommand struct {
	prevLocationAreas pokemonworld.PreviousLocationAreasFunc
	getLocationAreas pokemonworld.GetLocationAreasFunc
}

func NewMapbCommand(prev pokemonworld.PreviousLocationAreasFunc, get pokemonworld.GetLocationAreasFunc) cliCommand {
	return &mapbCommand{
		prevLocationAreas: prev,
		getLocationAreas: get,
	}
}

func (c *mapbCommand) Execute(params ...string) error {
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