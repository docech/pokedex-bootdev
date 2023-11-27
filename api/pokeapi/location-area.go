package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/docech/pokedex-bootdev/api"
	"github.com/docech/pokedex-bootdev/api/http"
	"github.com/docech/pokedex-bootdev/domain/pokemonworld"
)

type locationAreaResource struct {
	resourceUrl string
	resource *pokemonworld.LocationArea
	fetcher http.FetchFunc
}

func NewLocationAreaResource(resourceUrl string, cacheConfig http.CacheConfig) api.DetailResource[string, pokemonworld.LocationArea] {
	return &locationAreaResource{
		resourceUrl: resourceUrl,
		resource: &pokemonworld.LocationArea{},
		fetcher: http.CachedFetch(cacheConfig),
	}
}

func (c *locationAreaResource) Detail(areaName string) error {
	return c.fetchResource(fmt.Sprintf("%s%s", c.resourceUrl, areaName))
}

func (c *locationAreaResource) Data() pokemonworld.LocationArea {
	return *c.resource
}

func (c *locationAreaResource) fetchResource(url string) error {
	data, err := c.fetcher(url)
	
	if err != nil {
		return err
	}

	var resource pokemonworld.LocationArea
	if err := json.Unmarshal(data, &resource); err != nil {
		return err
	}

	c.resource = &resource

	return nil
}

