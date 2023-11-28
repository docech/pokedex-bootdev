package commands

import (
	"errors"

	"github.com/docech/pokedex-bootdev/domain/pokemonworld"
)

type inspectCommand struct {
	pokedex pokemonworld.Pokedex
}

func NewInspectCommand(pokedex pokemonworld.Pokedex) cliCommand {
	return &inspectCommand{
		pokedex: pokedex,
	}
}

func (c *inspectCommand) Execute(params ...string) error {
	if len(params) == 0 {
		return errors.New("missing pokemon name")
	}
	if len(params) > 1 {
		return errors.New("too many parameters, expected one pokemon name")
	}

	pokemonName := params[0]
	
	return c.pokedex.InspectPokemon(pokemonName)
}

func (c *inspectCommand) About() aboutCommand {
	return aboutCommand{
		name:        "inspect",
		description: "Inspects already caught Pokemons",
	}
}