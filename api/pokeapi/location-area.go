package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/docech/pokedex-bootdev/api"
	"github.com/docech/pokedex-bootdev/api/http"
	"github.com/docech/pokedex-bootdev/domain/pokedex"
)

type locationAreaResource struct {
	value pokedex.LocationArea
}

func NewLocationAreaResource(resourceUrl string, cacheConfig http.CacheConfig) func (areaName string) (api.Resource[pokedex.LocationArea], error) {
	fetcher := fetchResource(http.CachedFetch(cacheConfig))

	return func (areaName string) (api.Resource[pokedex.LocationArea], error) {
		value, err := fetcher(fmt.Sprintf("%s%s", resourceUrl, areaName))
		if err != nil {
			return nil, err
		}

		resource := locationAreaResource{
			value: value,
		}
		
		return &resource, nil
	}
}

func (c *locationAreaResource) Data() pokedex.LocationArea {
	return c.value
}

func fetchResource(fetcher http.FetchFunc) func (url string) (pokedex.LocationArea, error) {
	return func (url string) (noop pokedex.LocationArea, err error) {
		data, err := fetcher(url)
		
		if err != nil {
			return noop, err
		}

		var resource pokedex.LocationArea
		if err := json.Unmarshal(data, &resource); err != nil {
			return noop, err
		}

		return resource, nil
	}
}

