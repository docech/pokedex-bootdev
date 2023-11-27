package pokemonworld

type Pokedex interface {
	CaughtPokemon(pokemon Pokemon)
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