package print

import (
	"fmt"

	"github.com/docech/pokedex-bootdev/domain/pokedex"
)

func PrintLocationAreas(locationAreas []pokedex.LocationAreaLink) {
	for _, area := range locationAreas {
		fmt.Println(area.Name)
	}
}

func PrintPokemonEncountersInLocationArea(area pokedex.LocationArea) {
	fmt.Println("Found Pokemon:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
}