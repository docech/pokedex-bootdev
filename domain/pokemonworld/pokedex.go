package pokemonworld

import (
	"errors"
	"fmt"
)

type Pokedex interface {
	CaughtPokemon(pokemon Pokemon)
	InspectPokemon(name string) error
}

type pokedex struct {
	caughtPokemons map[string]Pokemon
}

func NewPokedex() Pokedex {
	return &pokedex{
		caughtPokemons: make(map[string]Pokemon),
	}
}

func (p *pokedex) CaughtPokemon(pokemon Pokemon) {
	p.caughtPokemons[pokemon.Name] = pokemon
}

func (p *pokedex) InspectPokemon(name string) error {
	pokemon, ok := p.caughtPokemons[name]
	if !ok {
		return errors.New("you haven't caught that pokemon yet!")
	}

	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}
