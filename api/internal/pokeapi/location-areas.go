package pokeapi

import (
	"encoding/json"
	"errors"

	"github.com/docech/pokedex-bootdev/api/internal"
	"github.com/docech/pokedex-bootdev/api/internal/http"
	"github.com/docech/pokedex-bootdev/domain/pokemonworld"
)

type locationAreasApiResource struct {
	Next     any         `json:"next"`
	Previous any         `json:"previous"`
	Results  []pokemonworld.LocationAreaLink `json:"results"`
}

type locationAreasResource struct {
	resource	 	*locationAreasApiResource
	fetcher 		 http.FetchFunc
}

func NewLocationAreasResource(resourceUrl string, cacheConfig http.CacheConfig) internal.ListResource[pokemonworld.LocationAreaLink] {
	return &locationAreasResource{
		resource: &locationAreasApiResource{
			Next: resourceUrl,
			Previous: nil,
			Results: []pokemonworld.LocationAreaLink{},
		},
		fetcher: http.CachedFetch(cacheConfig),
	}
}

func (c *locationAreasResource) Next() error {
	nextURL, ok := c.resource.Next.(string)
	if !ok {
		return errors.New("no next resource")
	}
	return c.fetchResource(nextURL)
}

func (c *locationAreasResource) Previous() error {
	previousURL, ok := c.resource.Previous.(string)
	if !ok {
		return errors.New("no previous resource")
	}
	return c.fetchResource(previousURL)
}

func (c *locationAreasResource) Data() []pokemonworld.LocationAreaLink {
	return c.resource.Results
}

func (c *locationAreasResource) fetchResource(url string) error {
	data, err := c.fetcher(url)
	
	if err != nil {
		return err
	}

	var resource locationAreasApiResource
	if err := json.Unmarshal(data, &resource); err != nil {
		return err
	}

	c.resource = &resource

	return nil
}

