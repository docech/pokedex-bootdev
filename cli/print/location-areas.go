package print

import (
	"fmt"

	"github.com/docech/pokedex-bootdev/domain/pokemonworld"
)

func PrintLocationAreas(locationAreas []pokemonworld.LocationAreaLink) {
	for _, area := range locationAreas {
		fmt.Println(area.Name)
	}
}

func PrintPokemonEncountersInLocationArea(area pokemonworld.LocationArea) {
	fmt.Println("Found Pokemon:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
}