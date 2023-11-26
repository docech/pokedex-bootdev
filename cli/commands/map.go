package commands

import (
	"fmt"

	"github.com/docech/pokedex-bootdev/domain/pokedex"
)

type MapResources struct {
	ProvideLocationAreas func () ([]pokedex.LocationArea, error)
}

func printLocationAreas(locationAreas []pokedex.LocationArea) {
	for _, area := range locationAreas {
		fmt.Println(area.Name)
	}
}

func handleExecute(res MapResources) error {
	areas, err := res.ProvideLocationAreas()
	if err != nil {
		return err
	}
	printLocationAreas(areas)
	return nil
}

type mapCommand struct {
	res MapResources
}

func NewMapCommand(res MapResources) cliCommand {
	return &mapCommand{
		res: res,
	}
}

func (c *mapCommand) Execute() error {
	return handleExecute(c.res)
}

func (c mapCommand) About() aboutCommand {
	return aboutCommand{
		name: "map",
		description: "Displays X location areas from Pokemon world. Subsequent calls will display the next X areas.",
	}
}

type mapbCommand struct {
	res MapResources
}

func NewMapbCommand(res MapResources) cliCommand {
	return &mapbCommand{
		res: res,
	}
}

func (c *mapbCommand) Execute() error {
	return handleExecute(c.res)
}

func (c mapbCommand) About() aboutCommand {
	return aboutCommand{
		name: "mapb",
		description: `Opposite of map command. Displays previous X location areas from Pokemon world.`,
	}
}