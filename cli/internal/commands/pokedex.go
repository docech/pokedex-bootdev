package commands

import (
	"fmt"

	"github.com/docech/pokedex-bootdev/domain/pokemonworld"
)

type pokedexCommand struct {
	pokedex pokemonworld.Pokedex
}

func NewPokedexCommand(pokedex pokemonworld.Pokedex) cliCommand {
	return &pokedexCommand{
		pokedex: pokedex,
	}
}

func (c *pokedexCommand) Execute(params ...string) error {
	fmt.Println("Your Pokedex:")
	return c.pokedex.ListCaughtPokemons()
}

func (c *pokedexCommand) About() aboutCommand {
	return aboutCommand{
		name:        "pokedex",
		description: "Lists caught Pokemons",
	}
}