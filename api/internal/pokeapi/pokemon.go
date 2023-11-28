package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/docech/pokedex-bootdev/api/internal"
	"github.com/docech/pokedex-bootdev/api/internal/http"
	"github.com/docech/pokedex-bootdev/domain/pokemonworld"
)

type pokemonResource struct {
	resourceUrl string
	resource *pokemonworld.Pokemon
	fetcher http.FetchFunc
}

func NewPokemonResource(resourceUrl string, cacheConfig http.CacheConfig) internal.DetailResource[string, pokemonworld.Pokemon] {
	return &pokemonResource{
		resourceUrl: resourceUrl,
		resource: &pokemonworld.Pokemon{},
		fetcher: http.CachedFetch(cacheConfig),
	}
}

func (r *pokemonResource) Detail(pokemonName string) error {
	return r.fetchResource(fmt.Sprintf("%s%s", r.resourceUrl, pokemonName))
}

func (r *pokemonResource) Data() pokemonworld.Pokemon {
	return *r.resource
}

func (r *pokemonResource) fetchResource(url string) error {
	data, err := r.fetcher(url)
	
	if err != nil {
		return err
	}

	var resource pokemonworld.Pokemon
	if err := json.Unmarshal(data, &resource); err != nil {
		return err
	}

	r.resource = &resource

	return nil
}

