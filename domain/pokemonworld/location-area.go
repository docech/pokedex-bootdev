package pokemonworld

type LocationAreaLink struct {
	Name string `json:"name"`
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