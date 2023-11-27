package pokeapi

import (
	"encoding/json"
	"errors"

	"github.com/docech/pokedex-bootdev/api"
	"github.com/docech/pokedex-bootdev/api/http"
	"github.com/docech/pokedex-bootdev/domain/pokedex"
)

type locationAreasApiResource struct {
	Next     any         `json:"next"`
	Previous any         `json:"previous"`
	Results         []pokedex.LocationArea `json:"results"`
}

type locationAreasResource struct {
	resource	 	*locationAreasApiResource
	fetcher 		 http.FetchFunc
}

func NewLocationAreasResource(resourceUrl string, cacheConfig http.CacheConfig) api.ListResource[pokedex.LocationArea] {
	return &locationAreasResource{
		resource: &locationAreasApiResource{
			Next: resourceUrl,
			Previous: nil,
			Results: []pokedex.LocationArea{},
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

func (c *locationAreasResource) Data() []pokedex.LocationArea {
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

