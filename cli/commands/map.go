package commands

import (
	"github.com/docech/pokedex-bootdev/cli/print"
	"github.com/docech/pokedex-bootdev/domain/pokedex"
)

type mapCommand struct {
	nextLocationAreas pokedex.NextLocationAreasFunc
	getLocationAreas pokedex.GetLocationAreasFunc
}

func NewMapCommand(next pokedex.NextLocationAreasFunc, get pokedex.GetLocationAreasFunc) cliCommand {
	return &mapCommand{
		nextLocationAreas: next,
		getLocationAreas: get,
	}
}

func (c *mapCommand) Execute(params ...string) error {
	if err := c.nextLocationAreas(); err != nil {
		return err
	}

	print.PrintLocationAreas(c.getLocationAreas())
	return nil
}

func (c mapCommand) About() aboutCommand {
	return aboutCommand{
		name: "map",
		description: "Displays X location areas from Pokemon world. Subsequent calls will display the next X areas.",
	}
}

