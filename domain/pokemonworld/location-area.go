package pokemonworld

import "fmt"

type LocationAreaLink struct {
	Name string `json:"name"`
}

func ShowSurroundingLocationAreas(locationAreas []LocationAreaLink) {
	for _, area := range locationAreas {
		fmt.Println(area.Name)
	}
}

type NextLocationAreasFunc = func() error
type PreviousLocationAreasFunc = func() error
type GetLocationAreasFunc = func() []LocationAreaLink

type LocationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type DetailLocationAreaFunc = func(name string) error
type GetLocationAreaFunc = func() LocationArea

func ExplorePokemonEncountersInLocationArea(area LocationArea) {
	fmt.Println("Found Pokemon:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
}