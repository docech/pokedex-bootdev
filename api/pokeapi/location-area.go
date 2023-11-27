package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/docech/pokedex-bootdev/api"
	"github.com/docech/pokedex-bootdev/api/http"
	"github.com/docech/pokedex-bootdev/domain/pokedex"
)

type locationAreaResource struct {
	resourceUrl string
	resource pokedex.LocationArea
	fetcher http.FetchFunc
}

func NewLocationAreaResource(resourceUrl string, cacheConfig http.CacheConfig) api.DetailResource[string, pokedex.LocationArea] {
	return &locationAreaResource{
		resourceUrl: resourceUrl,
		resource: pokedex.LocationArea{},
		fetcher: http.CachedFetch(cacheConfig),
	}
}

func (c *locationAreaResource) Detail(areaName string) error {
	return c.fetchResource(fmt.Sprintf("%s%s", c.resourceUrl, areaName))
}

func (c *locationAreaResource) Data() pokedex.LocationArea {
	return c.resource
}

func (c *locationAreaResource) fetchResource(url string) error {
	data, err := c.fetcher(url)
	
	if err != nil {
		return err
	}

	var resource pokedex.LocationArea
	if err := json.Unmarshal(data, &resource); err != nil {
		return err
	}

	c.resource = resource

	return nil
}

