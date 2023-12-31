package commands

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/docech/pokedex-bootdev/domain/pokemonworld"
)

type catchCommand struct {
	pokedex pokemonworld.Pokedex
	detailPokemon pokemonworld.DetailPokemonFunc
	getPokemon pokemonworld.GetPokemonFunc
}

func NewCatchCommand(
	pokedex pokemonworld.Pokedex, 
	detail pokemonworld.DetailPokemonFunc, 
	get pokemonworld.GetPokemonFunc,
) cliCommand {
	return catchCommand{
		pokedex: pokedex,
		detailPokemon: detail,
		getPokemon: get,
	}
}

func (c catchCommand) Execute(params ...string) error {
	if len(params) == 0 {
		return errors.New("missing pokemon name")
	}
	if len(params) > 1 {
		return errors.New("too many parameters, expected one pokemon name")
	}
	
	pokemonName := params[0]
	err := c.detailPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing Pokeball at %s...\n", pokemonName)
	baseExperience := c.getPokemon().BaseExperience
	catchProbability := float64(baseExperience) / 300.0
	if rand.Float64() < catchProbability {
		fmt.Printf("%s was caught!\n", pokemonName)
		c.pokedex.CaughtPokemon(c.getPokemon())
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func (c catchCommand) About() aboutCommand {
	return aboutCommand{
		name:        "catch",
		description: "Catches a Pokemon",
	}
}