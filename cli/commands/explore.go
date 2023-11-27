package commands

import (
	"errors"
	"fmt"

	"github.com/docech/pokedex-bootdev/cli/print"
	"github.com/docech/pokedex-bootdev/domain/pokemonworld"
)

type exploreCommand struct {
	detailLocationArea pokemonworld.DetailLocationAreaFunc
	getLocationArea pokemonworld.GetLocationAreaFunc
}

func NewExploreCommand(detail pokemonworld.DetailLocationAreaFunc, get pokemonworld.GetLocationAreaFunc) cliCommand {
	return &exploreCommand{
		detailLocationArea: detail,
		getLocationArea: get,
	}
}

func (c *exploreCommand) Execute(params ...string) error {
	if len(params) > 1 {
		return errors.New("too many parameters, expected one area name")
	}

	fmt.Printf("Exploring %s...\n", params[0])
	err := c.detailLocationArea(params[0])
	if err != nil {
		return err
	}

	print.PrintPokemonEncountersInLocationArea(c.getLocationArea())
	return nil
}

func (c *exploreCommand) About() aboutCommand {
	return aboutCommand{
		name:        "explore",
		description: "Explore the Pokemon location area",
	}
}