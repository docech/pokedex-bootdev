package pokemonworld

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

type DetailPokemonFunc = func(name string) error
type GetPokemonFunc = func() Pokemon