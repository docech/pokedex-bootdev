package api

import (
	"time"

	"github.com/docech/pokedex-bootdev/api/internal"
	"github.com/docech/pokedex-bootdev/api/internal/http"
	"github.com/docech/pokedex-bootdev/api/internal/pokeapi"
	"github.com/docech/pokedex-bootdev/domain/pokemonworld"
)

type pokeAPI struct {
	LocationAreasResource internal.ListResource[pokemonworld.LocationAreaLink]
	LocationAreaResource  internal.DetailResource[string, pokemonworld.LocationArea]
	PokemonResource       internal.DetailResource[string, pokemonworld.Pokemon]
}

func PokeAPI() *pokeAPI {
	caching := http.CacheConfig{
		MaxAge: 60 * time.Second,
	}

	return &pokeAPI{
		LocationAreasResource: pokeapi.NewLocationAreasResource(
			"https://pokeapi.co/api/v2/location-area/",
			caching,
		),
		LocationAreaResource: pokeapi.NewLocationAreaResource(
			"https://pokeapi.co/api/v2/location-area/",
			caching,
		),
		PokemonResource: pokeapi.NewPokemonResource(
			"https://pokeapi.co/api/v2/pokemon/",
			caching,
		),
	}
}